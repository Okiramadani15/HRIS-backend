package controllers

import (
	"encoding/csv"
	"fmt"
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreateAllowance - Buat tunjangan
func CreateAllowance(c *fiber.Ctx) error {
	var allowance models.Allowance
	if err := c.BodyParser(&allowance); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	// Validation
	if allowance.EmployeeID == "" {
		return utils.ValidationErrorResponse(c, "Employee ID is required")
	}
	if allowance.Type == "" {
		return utils.ValidationErrorResponse(c, "Allowance type is required")
	}
	if allowance.Amount <= 0 {
		return utils.ValidationErrorResponse(c, "Amount must be greater than 0")
	}

	// Check if employee exists
	var employee models.Employee
	if err := config.DB.Where("employee_id = ?", allowance.EmployeeID).First(&employee).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}

	if err := config.DB.Create(&allowance).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to create allowance")
	}

	return utils.SuccessResponse(c, "Allowance created successfully", allowance)
}

// CreateDeduction - Buat potongan
func CreateDeduction(c *fiber.Ctx) error {
	var deduction models.Deduction
	if err := c.BodyParser(&deduction); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	// Validation
	if deduction.EmployeeID == "" {
		return utils.ValidationErrorResponse(c, "Employee ID is required")
	}
	if deduction.Type == "" {
		return utils.ValidationErrorResponse(c, "Deduction type is required")
	}
	if deduction.Amount <= 0 {
		return utils.ValidationErrorResponse(c, "Amount must be greater than 0")
	}

	// Check if employee exists
	var employee models.Employee
	if err := config.DB.Where("employee_id = ?", deduction.EmployeeID).First(&employee).Error; err != nil {
		return utils.NotFoundResponse(c, "Employee not found")
	}

	if err := config.DB.Create(&deduction).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to create deduction")
	}

	return utils.SuccessResponse(c, "Deduction created successfully", deduction)
}

// GeneratePayroll - Generate payroll untuk semua karyawan
func GeneratePayroll(c *fiber.Ctx) error {
	type Request struct {
		Month string `json:"month"`
		Year  string `json:"year"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Validation
	if req.Month == "" || req.Year == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Month and year are required"})
	}

	// Check if payroll already exists
	var existingCount int64
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", req.Month, req.Year).Count(&existingCount)
	if existingCount > 0 {
		return c.Status(409).JSON(fiber.Map{"error": "Payroll for this month/year already exists"})
	}

	// Ambil semua employee
	var employees []models.Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch employees"})
	}

	var payrolls []models.Payroll

	for _, emp := range employees {
		// Hitung allowances
		var totalAllowances float64
		config.DB.Model(&models.Allowance{}).Where("employee_id = ?", emp.EmployeeID).Select("COALESCE(sum(amount), 0)").Scan(&totalAllowances)

		// Hitung deductions
		var totalDeductions float64
		config.DB.Model(&models.Deduction{}).Where("employee_id = ?", emp.EmployeeID).Select("COALESCE(sum(amount), 0)").Scan(&totalDeductions)

		// Buat payroll
		payroll := models.Payroll{
			EmployeeID:      emp.EmployeeID,
			Month:           req.Month,
			Year:            req.Year,
			BaseSalary:      emp.Salary,
			TotalAllowances: totalAllowances,
			TotalDeductions: totalDeductions,
			NetSalary:       emp.Salary + totalAllowances - totalDeductions,
			Status:          "generated",
		}

		if err := config.DB.Create(&payroll).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create payroll"})
		}

		payrolls = append(payrolls, payroll)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Payroll generated",
		"data":    payrolls,
	})
}

// GetPayrolls - Ambil semua payroll
func GetPayrolls(c *fiber.Ctx) error {
	var payrolls []models.Payroll
	if err := config.DB.Find(&payrolls).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch payrolls"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    payrolls,
	})
}

// UpdatePayrollStatus - Update status payroll
func UpdatePayrollStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	
	type Request struct {
		Status string `json:"status"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var payroll models.Payroll
	if err := config.DB.First(&payroll, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payroll not found"})
	}

	payroll.Status = req.Status
	if req.Status == "paid" {
		now := time.Now()
		payroll.PaidAt = &now
	}

	if err := config.DB.Save(&payroll).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update status"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Status updated",
		"data":    payroll,
	})
}

// ExportPayrollCSV - Export payroll ke CSV
func ExportPayrollCSV(c *fiber.Ctx) error {
	month := c.Query("month")
	year := c.Query("year")
	
	// Query payroll dengan filter
	query := config.DB.Model(&models.Payroll{})
	if month != "" {
		query = query.Where("month = ?", month)
	}
	if year != "" {
		query = query.Where("year = ?", year)
	}
	
	var payrolls []models.Payroll
	if err := query.Find(&payrolls).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch payroll data"})
	}
	
	// Set headers untuk download CSV
	filename := fmt.Sprintf("payroll_%s_%s.csv", month, year)
	if month == "" || year == "" {
		filename = "payroll_all.csv"
	}
	
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	
	// Buat CSV writer
	writer := csv.NewWriter(c.Response().BodyWriter())
	defer writer.Flush()
	
	// Header CSV
	headers := []string{
		"ID",
		"Employee ID",
		"Employee Name",
		"Month",
		"Year",
		"Base Salary",
		"Total Allowances",
		"Total Deductions",
		"Net Salary",
		"Status",
		"Paid At",
		"Created At",
	}
	
	if err := writer.Write(headers); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to write CSV headers"})
	}
	
	// Data rows
	for _, payroll := range payrolls {
		// Ambil nama employee
		var employee models.Employee
		employeeName := "Unknown"
		if err := config.DB.Where("employee_id = ?", payroll.EmployeeID).First(&employee).Error; err == nil {
			employeeName = employee.FullName
		}
		
		// Format paid_at
		paidAt := ""
		if payroll.PaidAt != nil {
			paidAt = payroll.PaidAt.Format("2006-01-02 15:04:05")
		}
		
		row := []string{
			strconv.Itoa(int(payroll.ID)),
			payroll.EmployeeID,
			employeeName,
			payroll.Month,
			payroll.Year,
			strconv.FormatFloat(payroll.BaseSalary, 'f', 0, 64),
			strconv.FormatFloat(payroll.TotalAllowances, 'f', 0, 64),
			strconv.FormatFloat(payroll.TotalDeductions, 'f', 0, 64),
			strconv.FormatFloat(payroll.NetSalary, 'f', 0, 64),
			payroll.Status,
			paidAt,
			payroll.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		
		if err := writer.Write(row); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to write CSV row"})
		}
	}
	
	return nil
}

// GetPayrollDetail - Get detail payroll by employee
func GetPayrollDetail(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	month := c.Query("month")
	year := c.Query("year")
	
	query := config.DB.Where("employee_id = ?", employeeID)
	if month != "" {
		query = query.Where("month = ?", month)
	}
	if year != "" {
		query = query.Where("year = ?", year)
	}
	
	var payrolls []models.Payroll
	if err := query.Find(&payrolls).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch payroll"})
	}
	
	// Ambil data employee
	var employee models.Employee
	if err := config.DB.Where("employee_id = ?", employeeID).First(&employee).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"employee": employee,
			"payrolls": payrolls,
		},
	})
}

// UpdatePayroll - Update payroll data
func UpdatePayroll(c *fiber.Ctx) error {
	id := c.Params("id")
	
	type UpdateRequest struct {
		BaseSalary      *float64 `json:"base_salary"`
		TotalAllowances *float64 `json:"total_allowances"`
		TotalDeductions *float64 `json:"total_deductions"`
	}
	
	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	
	var payroll models.Payroll
	if err := config.DB.First(&payroll, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payroll not found"})
	}
	
	// Update fields if provided
	if req.BaseSalary != nil {
		payroll.BaseSalary = *req.BaseSalary
	}
	if req.TotalAllowances != nil {
		payroll.TotalAllowances = *req.TotalAllowances
	}
	if req.TotalDeductions != nil {
		payroll.TotalDeductions = *req.TotalDeductions
	}
	
	// Recalculate net salary
	payroll.NetSalary = payroll.BaseSalary + payroll.TotalAllowances - payroll.TotalDeductions
	
	if err := config.DB.Save(&payroll).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update payroll"})
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Payroll updated",
		"data":    payroll,
	})
}

// DeleteAllowance - Delete allowance
func DeleteAllowance(c *fiber.Ctx) error {
	id := c.Params("id")
	
	if err := config.DB.Delete(&models.Allowance{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete allowance"})
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Allowance deleted",
	})
}

// DeleteDeduction - Delete deduction
func DeleteDeduction(c *fiber.Ctx) error {
	id := c.Params("id")
	
	if err := config.DB.Delete(&models.Deduction{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete deduction"})
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Deduction deleted",
	})
}

// GetAllowances - Get allowances
func GetAllowances(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	employeeID := c.Query("employee_id")
	
	query := config.DB.Model(&models.Allowance{})
	
	// Check if user has role loaded
	if user.Role.Name == "admin" || user.Role.Name == "hr" {
		// HR/Admin can see all or filter by employee_id
		if employeeID != "" {
			query = query.Where("employee_id = ?", employeeID)
		}
	} else {
		// For regular employees or users without specific roles
		var employee models.Employee
		err := config.DB.Where("user_id = ?", user.ID).First(&employee).Error
		if err != nil {
			// Return empty result if employee record not found
			return utils.SuccessResponse(c, "No allowances found", []models.Allowance{})
		}
		query = query.Where("employee_id = ?", employee.EmployeeID)
	}
	
	var allowances []models.Allowance
	if err := query.Find(&allowances).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to fetch allowances")
	}
	
	return utils.SuccessResponse(c, "Allowances retrieved successfully", allowances)
}

// GetDeductions - Get deductions
func GetDeductions(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	employeeID := c.Query("employee_id")
	
	query := config.DB.Model(&models.Deduction{})
	
	// If user is HR/Admin, they can see all or filter by employee_id
	if user.Role.Name == "admin" || user.Role.Name == "hr" {
		if employeeID != "" {
			query = query.Where("employee_id = ?", employeeID)
		}
	} else {
		// For regular employees, find their employee record
		var employee models.Employee
		err := config.DB.Where("user_id = ?", user.ID).First(&employee).Error
		if err != nil {
			// Return empty result if employee record not found
			return utils.SuccessResponse(c, "No deductions found", []models.Deduction{})
		}
		query = query.Where("employee_id = ?", employee.EmployeeID)
	}
	
	var deductions []models.Deduction
	if err := query.Find(&deductions).Error; err != nil {
		return utils.InternalErrorResponse(c, "Failed to fetch deductions")
	}
	
	return utils.SuccessResponse(c, "Deductions retrieved successfully", deductions)
}

// GetPayrollSummary - Get payroll summary statistics
func GetPayrollSummary(c *fiber.Ctx) error {
	month := c.Query("month")
	year := c.Query("year")
	
	if month == "" || year == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Month and year are required"})
	}
	
	type Summary struct {
		TotalEmployees   int64   `json:"total_employees"`
		TotalPaid        int64   `json:"total_paid"`
		TotalPending     int64   `json:"total_pending"`
		TotalBaseSalary  float64 `json:"total_base_salary"`
		TotalAllowances  float64 `json:"total_allowances"`
		TotalDeductions  float64 `json:"total_deductions"`
		TotalNetSalary   float64 `json:"total_net_salary"`
		AverageNetSalary float64 `json:"average_net_salary"`
	}
	
	var summary Summary
	
	// Count total employees with payroll
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", month, year).Count(&summary.TotalEmployees)
	
	// Count paid vs pending
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ? AND status = ?", month, year, "paid").Count(&summary.TotalPaid)
	summary.TotalPending = summary.TotalEmployees - summary.TotalPaid
	
	// Sum totals
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", month, year).Select("COALESCE(sum(base_salary), 0)").Scan(&summary.TotalBaseSalary)
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", month, year).Select("COALESCE(sum(total_allowances), 0)").Scan(&summary.TotalAllowances)
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", month, year).Select("COALESCE(sum(total_deductions), 0)").Scan(&summary.TotalDeductions)
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", month, year).Select("COALESCE(sum(net_salary), 0)").Scan(&summary.TotalNetSalary)
	
	// Calculate average
	if summary.TotalEmployees > 0 {
		summary.AverageNetSalary = summary.TotalNetSalary / float64(summary.TotalEmployees)
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"data":    summary,
	})
}

// GetPayrollDashboard - Get dashboard data
func GetPayrollDashboard(c *fiber.Ctx) error {
	type DashboardData struct {
		CurrentMonth struct {
			Month           string  `json:"month"`
			Year            string  `json:"year"`
			TotalEmployees  int64   `json:"total_employees"`
			TotalPaid       int64   `json:"total_paid"`
			TotalNetSalary  float64 `json:"total_net_salary"`
		} `json:"current_month"`
		RecentPayrolls []models.Payroll `json:"recent_payrolls"`
		TopAllowances  []struct {
			Name   string  `json:"name"`
			Total  float64 `json:"total"`
			Count  int64   `json:"count"`
		} `json:"top_allowances"`
	}
	
	var dashboard DashboardData
	
	// Current month data (assume current month)
	currentTime := time.Now()
	dashboard.CurrentMonth.Month = fmt.Sprintf("%02d", currentTime.Month())
	dashboard.CurrentMonth.Year = fmt.Sprintf("%d", currentTime.Year())
	
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", dashboard.CurrentMonth.Month, dashboard.CurrentMonth.Year).Count(&dashboard.CurrentMonth.TotalEmployees)
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ? AND status = ?", dashboard.CurrentMonth.Month, dashboard.CurrentMonth.Year, "paid").Count(&dashboard.CurrentMonth.TotalPaid)
	config.DB.Model(&models.Payroll{}).Where("month = ? AND year = ?", dashboard.CurrentMonth.Month, dashboard.CurrentMonth.Year).Select("COALESCE(sum(net_salary), 0)").Scan(&dashboard.CurrentMonth.TotalNetSalary)
	
	// Recent payrolls (last 10)
	config.DB.Order("created_at DESC").Limit(10).Find(&dashboard.RecentPayrolls)
	
	// Top allowances
	config.DB.Model(&models.Allowance{}).Select("name, sum(amount) as total, count(*) as count").Group("name").Order("total DESC").Limit(5).Scan(&dashboard.TopAllowances)
	
	return c.JSON(fiber.Map{
		"success": true,
		"data":    dashboard,
	})
}

// CreateBulkAllowances - Create multiple allowances
func CreateBulkAllowances(c *fiber.Ctx) error {
	type BulkRequest struct {
		Allowances []models.Allowance `json:"allowances"`
	}
	
	var req BulkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}
	
	if len(req.Allowances) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No allowances provided"})
	}
	
	// Validate each allowance
	for i, allowance := range req.Allowances {
		if allowance.EmployeeID == "" {
			return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Employee ID is required for allowance %d", i+1)})
		}
		if allowance.Amount <= 0 {
			return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Amount must be greater than 0 for allowance %d", i+1)})
		}
	}
	
	// Create all allowances
	if err := config.DB.Create(&req.Allowances).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create allowances"})
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("%d allowances created successfully", len(req.Allowances)),
		"data":    req.Allowances,
	})
}

// CreateBulkDeductions - Create multiple deductions
func CreateBulkDeductions(c *fiber.Ctx) error {
	type BulkRequest struct {
		Deductions []models.Deduction `json:"deductions"`
	}
	
	var req BulkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}
	
	if len(req.Deductions) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No deductions provided"})
	}
	
	// Validate each deduction
	for i, deduction := range req.Deductions {
		if deduction.EmployeeID == "" {
			return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Employee ID is required for deduction %d", i+1)})
		}
		if deduction.Amount <= 0 {
			return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Amount must be greater than 0 for deduction %d", i+1)})
		}
	}
	
	// Create all deductions
	if err := config.DB.Create(&req.Deductions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create deductions"})
	}
	
	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("%d deductions created successfully", len(req.Deductions)),
		"data":    req.Deductions,
	})
}