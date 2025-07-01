package controllers

import (
	"fmt"
	"hris-backend/config"
	"hris-backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ===== KPI CATEGORIES =====

func CreateKPICategory(c *fiber.Ctx) error {
	var category models.KPICategory
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	if category.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Category name is required"})
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create KPI category"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "KPI category created successfully",
		"data":    category,
	})
}

func GetKPICategories(c *fiber.Ctx) error {
	var categories []models.KPICategory
	query := config.DB.Model(&models.KPICategory{})

	if c.Query("active") == "true" {
		query = query.Where("is_active = ?", true)
	}

	if err := query.Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch KPI categories"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    categories,
	})
}

// ===== KPI METRICS =====

func CreateKPIMetric(c *fiber.Ctx) error {
	var metric models.KPIMetric
	if err := c.BodyParser(&metric); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Validation
	if metric.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Metric name is required"})
	}
	if metric.CategoryID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Category ID is required"})
	}
	if metric.TargetType == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Target type is required"})
	}

	// Validate target type
	validTargetTypes := []string{"higher_better", "lower_better", "exact"}
	isValidTargetType := false
	for _, validType := range validTargetTypes {
		if metric.TargetType == validType {
			isValidTargetType = true
			break
		}
	}
	if !isValidTargetType {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid target type"})
	}

	if err := config.DB.Create(&metric).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create KPI metric"})
	}

	// Load category
	config.DB.Preload("Category").First(&metric, metric.ID)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "KPI metric created successfully",
		"data":    metric,
	})
}

func GetKPIMetrics(c *fiber.Ctx) error {
	var metrics []models.KPIMetric
	query := config.DB.Preload("Category")

	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if c.Query("active") == "true" {
		query = query.Where("is_active = ?", true)
	}

	if err := query.Find(&metrics).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch KPI metrics"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    metrics,
	})
}

// ===== KPI TARGETS =====

func SetKPITarget(c *fiber.Ctx) error {
	var target models.KPITarget
	if err := c.BodyParser(&target); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Validation
	if target.EmployeeID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Employee ID is required"})
	}
	if target.MetricID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Metric ID is required"})
	}
	if target.Period == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Period is required"})
	}

	// Check if employee exists
	var employee models.Employee
	if err := config.DB.Where("employee_id = ?", target.EmployeeID).First(&employee).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	// Set metadata
	userID := c.Locals("user_id")
	if userID != nil {
		target.SetBy = fmt.Sprintf("%v", userID)
	}
	target.SetAt = time.Now()

	if err := config.DB.Create(&target).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to set KPI target"})
	}

	// Load relationships
	config.DB.Preload("Employee").Preload("Metric").First(&target, target.ID)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "KPI target set successfully",
		"data":    target,
	})
}

func GetKPITargets(c *fiber.Ctx) error {
	var targets []models.KPITarget
	query := config.DB.Preload("Employee").Preload("Metric").Preload("Metric.Category")

	if employeeID := c.Query("employee_id"); employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}
	if period := c.Query("period"); period != "" {
		query = query.Where("period = ?", period)
	}
	if periodType := c.Query("period_type"); periodType != "" {
		query = query.Where("period_type = ?", periodType)
	}

	if err := query.Find(&targets).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch KPI targets"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    targets,
	})
}

// ===== KPI ACTUALS =====

func RecordKPIActual(c *fiber.Ctx) error {
	var actual models.KPIActual
	if err := c.BodyParser(&actual); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// Validation
	if actual.EmployeeID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Employee ID is required"})
	}
	if actual.MetricID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Metric ID is required"})
	}
	if actual.Period == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Period is required"})
	}

	// Set metadata
	userID := c.Locals("user_id")
	if userID != nil {
		actual.RecordedBy = fmt.Sprintf("%v", userID)
	}
	actual.RecordedAt = time.Now()

	if err := config.DB.Create(&actual).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to record KPI actual"})
	}

	// Calculate achievement
	calculateKPIAchievement(&actual)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "KPI actual recorded successfully",
		"data":    actual,
	})
}

func GetKPIActuals(c *fiber.Ctx) error {
	var actuals []models.KPIActual
	query := config.DB.Preload("Employee").Preload("Metric").Preload("Metric.Category")

	if employeeID := c.Query("employee_id"); employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}
	if period := c.Query("period"); period != "" {
		query = query.Where("period = ?", period)
	}
	if periodType := c.Query("period_type"); periodType != "" {
		query = query.Where("period_type = ?", periodType)
	}

	if err := query.Find(&actuals).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch KPI actuals"})
	}

	// Calculate achievements for each record
	for i := range actuals {
		calculateKPIAchievement(&actuals[i])
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    actuals,
	})
}

// ===== KPI DASHBOARD & ANALYTICS =====

func GetKPIDashboard(c *fiber.Ctx) error {
	employeeID := c.Query("employee_id")
	period := c.Query("period")
	
	if employeeID == "" || period == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Employee ID and period are required"})
	}

	type DashboardData struct {
		Employee      models.Employee    `json:"employee"`
		TotalMetrics  int64             `json:"total_metrics"`
		CompletedKPIs int64             `json:"completed_kpis"`
		AverageScore  float64           `json:"average_score"`
		Grade         string            `json:"grade"`
		KPIDetails    []models.KPIActual `json:"kpi_details"`
	}

	var dashboard DashboardData

	// Get employee
	if err := config.DB.Where("employee_id = ?", employeeID).First(&dashboard.Employee).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	// Get KPI statistics
	config.DB.Model(&models.KPITarget{}).Where("employee_id = ? AND period = ?", employeeID, period).Count(&dashboard.TotalMetrics)
	config.DB.Model(&models.KPIActual{}).Where("employee_id = ? AND period = ?", employeeID, period).Count(&dashboard.CompletedKPIs)

	// Get KPI details with achievements
	var actuals []models.KPIActual
	config.DB.Preload("Metric").Preload("Metric.Category").Where("employee_id = ? AND period = ?", employeeID, period).Find(&actuals)

	totalScore := 0.0
	totalWeight := 0.0
	for i := range actuals {
		calculateKPIAchievement(&actuals[i])
		if actuals[i].Score != nil {
			totalScore += *actuals[i].Score
			totalWeight += actuals[i].Metric.Weight
		}
	}

	if totalWeight > 0 {
		dashboard.AverageScore = (totalScore / totalWeight) * 100
	}

	// Calculate grade
	dashboard.Grade = calculateGrade(dashboard.AverageScore)
	dashboard.KPIDetails = actuals

	return c.JSON(fiber.Map{
		"success": true,
		"data":    dashboard,
	})
}

func GetKPIReport(c *fiber.Ctx) error {
	period := c.Query("period")
	periodType := c.Query("period_type")
	
	if period == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Period is required"})
	}

	type EmployeeKPIReport struct {
		EmployeeID   string  `json:"employee_id"`
		EmployeeName string  `json:"employee_name"`
		Department   string  `json:"department"`
		TotalKPIs    int64   `json:"total_kpis"`
		CompletedKPIs int64  `json:"completed_kpis"`
		AverageScore float64 `json:"average_score"`
		Grade        string  `json:"grade"`
	}

	var reports []EmployeeKPIReport

	// Get all employees with KPI data
	rows, err := config.DB.Raw(`
		SELECT 
			e.employee_id,
			e.full_name as employee_name,
			e.department,
			COUNT(DISTINCT t.id) as total_kpis,
			COUNT(DISTINCT a.id) as completed_kpis,
			COALESCE(AVG(
				CASE 
					WHEN m.target_type = 'higher_better' THEN (a.actual_value / t.target_value) * m.weight
					WHEN m.target_type = 'lower_better' THEN (t.target_value / a.actual_value) * m.weight
					WHEN m.target_type = 'exact' THEN (1 - ABS(a.actual_value - t.target_value) / t.target_value) * m.weight
				END
			) * 100, 0) as average_score
		FROM employees e
		LEFT JOIN kpi_targets t ON e.employee_id = t.employee_id AND t.period = ?
		LEFT JOIN kpi_actuals a ON e.employee_id = a.employee_id AND a.period = ? AND a.metric_id = t.metric_id
		LEFT JOIN kpi_metrics m ON t.metric_id = m.id
		WHERE t.id IS NOT NULL
		GROUP BY e.employee_id, e.full_name, e.department
		ORDER BY average_score DESC
	`, period, period).Rows()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate KPI report"})
	}
	defer rows.Close()

	for rows.Next() {
		var report EmployeeKPIReport
		rows.Scan(&report.EmployeeID, &report.EmployeeName, &report.Department, 
				 &report.TotalKPIs, &report.CompletedKPIs, &report.AverageScore)
		report.Grade = calculateGrade(report.AverageScore)
		reports = append(reports, report)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    reports,
	})
}

// ===== HELPER FUNCTIONS =====

func calculateKPIAchievement(actual *models.KPIActual) {
	// Get target value
	var target models.KPITarget
	if err := config.DB.Where("employee_id = ? AND metric_id = ? AND period = ?", 
		actual.EmployeeID, actual.MetricID, actual.Period).First(&target).Error; err != nil {
		return
	}

	// Get metric info
	var metric models.KPIMetric
	if err := config.DB.First(&metric, actual.MetricID).Error; err != nil {
		return
	}

	actual.TargetValue = &target.TargetValue

	// Calculate achievement percentage
	var achievement float64
	switch metric.TargetType {
	case "higher_better":
		achievement = (actual.ActualValue / target.TargetValue) * 100
	case "lower_better":
		achievement = (target.TargetValue / actual.ActualValue) * 100
	case "exact":
		if target.TargetValue != 0 {
			achievement = (1 - (abs(actual.ActualValue-target.TargetValue) / target.TargetValue)) * 100
		}
	}

	actual.Achievement = &achievement

	// Calculate weighted score
	score := achievement * metric.Weight / 100
	actual.Score = &score
}

func calculateGrade(score float64) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}
	return "E"
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// ===== BULK OPERATIONS =====

func SetBulkKPITargets(c *fiber.Ctx) error {
	type BulkTargetRequest struct {
		Targets []models.KPITarget `json:"targets"`
	}

	var req BulkTargetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	if len(req.Targets) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No targets provided"})
	}

	// Set metadata for all targets
	userID := c.Locals("user_id")
	now := time.Now()
	for i := range req.Targets {
		if userID != nil {
			req.Targets[i].SetBy = fmt.Sprintf("%v", userID)
		}
		req.Targets[i].SetAt = now
	}

	if err := config.DB.Create(&req.Targets).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to set bulk KPI targets"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("%d KPI targets set successfully", len(req.Targets)),
		"data":    req.Targets,
	})
}