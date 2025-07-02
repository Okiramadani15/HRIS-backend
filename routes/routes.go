package routes

import (
	"hris-backend/config"
	"hris-backend/controllers"
	"hris-backend/middleware"
	"hris-backend/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// Protected routes
	protected := api.Group("/", middleware.JWTMiddlewareWithDB(config.DB))
	protected.Get("/profile", getProfile)

	// Employee routes (JWT protected)
	protected.Post("/employees", controllers.CreateEmployee)
	protected.Get("/employees", controllers.GetEmployees)
	protected.Get("/employees/:id", controllers.GetEmployee)
	protected.Put("/employees/:id", controllers.UpdateEmployee)
	protected.Get("/my-profile", controllers.GetMyProfile)
	
	// Employee Photo routes (JWT protected)
	protected.Post("/employees/:id/photo", controllers.UploadEmployeePhoto)
	protected.Delete("/employees/:id/photo", controllers.DeleteEmployeePhoto)
	
	// Serve static files for uploads
	app.Static("/uploads", "./uploads")

	// Position routes
	protected.Post("/positions", controllers.CreatePosition)
	protected.Get("/positions", controllers.GetAllPositions)
	protected.Get("/positions/:id", controllers.GetPositionByID)
	protected.Put("/positions/:id", controllers.UpdatePosition)
	protected.Delete("/positions/:id", controllers.DeletePosition)

	// Department routes
	protected.Post("/departments", controllers.CreateDepartment)
	protected.Get("/departments", controllers.GetDepartments)
	protected.Get("/departments/:id", controllers.GetDepartment)
	protected.Put("/departments/:id", controllers.UpdateDepartment)
	protected.Delete("/departments/:id", controllers.DeleteDepartment)

	// Location routes
	protected.Post("/locations", controllers.CreateLocation)
	protected.Get("/locations", controllers.GetLocations)
	protected.Get("/locations/:id", controllers.GetLocation)
	protected.Put("/locations/:id", controllers.UpdateLocation)
	protected.Delete("/locations/:id", controllers.DeleteLocation)

	// Job Level routes
	protected.Post("/job-levels", controllers.CreateJobLevel)
	protected.Get("/job-levels", controllers.GetJobLevels)
	protected.Get("/job-levels/:id", controllers.GetJobLevel)
	protected.Put("/job-levels/:id", controllers.UpdateJobLevel)
	protected.Delete("/job-levels/:id", controllers.DeleteJobLevel)

	// Education Routes
	protected.Post("/educations", controllers.CreateEducation)
	protected.Get("/educations", controllers.GetEducations)
	protected.Get("/educations/:id", controllers.GetEducation)
	protected.Put("/educations/:id", controllers.UpdateEducation)
	protected.Delete("/educations/:id", controllers.DeleteEducation)

	// Marital Status Routes
	protected.Post("/marital-statuses", controllers.CreateMaritalStatus)
	protected.Get("/marital-statuses", controllers.GetMaritalStatuses)
	protected.Get("/marital-statuses/:id", controllers.GetMaritalStatus)
	protected.Put("/marital-statuses/:id", controllers.UpdateMaritalStatus)
	protected.Delete("/marital-statuses/:id", controllers.DeleteMaritalStatus)

	//Religion Routes
	protected.Post("/religions", controllers.CreateReligion)
	protected.Get("/religions", controllers.GetReligions)
	protected.Get("/religions/:id", controllers.GetReligion)
	protected.Put("/religions/:id", controllers.UpdateReligion)
	protected.Delete("/religions/:id", controllers.DeleteReligion)

	//Employee Type Routes
	protected.Post("/employment-types", controllers.CreateEmploymentType)
	protected.Get("/employment-types", controllers.GetEmploymentTypes)
	protected.Get("/employment-types/:id", controllers.GetEmploymentType)
	protected.Put("/employment-types/:id", controllers.UpdateEmploymentType)
	protected.Delete("/employment-types/:id", controllers.DeleteEmploymentType)

	// Work Shift routes
	protected.Post("/work-shifts", controllers.CreateWorkShift)
	protected.Get("/work-shifts", controllers.GetWorkShifts)
	protected.Get("/work-shifts/:id", controllers.GetWorkShift)
	protected.Put("/work-shifts/:id", controllers.UpdateWorkShift)
	protected.Delete("/work-shifts/:id", controllers.DeleteWorkShift)

	// Leave Type routes
	protected.Post("/leave-types", controllers.CreateLeaveType)
	protected.Get("/leave-types", controllers.GetLeaveTypes)
	protected.Get("/leave-types/:id", controllers.GetLeaveType)
	protected.Put("/leave-types/:id", controllers.UpdateLeaveType)
	protected.Delete("/leave-types/:id", controllers.DeleteLeaveType)

	// Bank List routes
	protected.Post("/banks", controllers.CreateBank)
	protected.Get("/banks", controllers.GetBanks)
	protected.Get("/banks/:id", controllers.GetBank)
	protected.Put("/banks/:id", controllers.UpdateBank)
	protected.Delete("/banks/:id", controllers.DeleteBank)

	// Pay Grade routes
	protected.Post("/pay-grades", controllers.CreatePayGrade)
	protected.Get("/pay-grades", controllers.GetPayGrades)
	protected.Get("/pay-grades/:id", controllers.GetPayGrade)
	protected.Put("/pay-grades/:id", controllers.UpdatePayGrade)
	protected.Delete("/pay-grades/:id", controllers.DeletePayGrade)

	// Attendance routes
	protected.Post("/attendances", controllers.CreateAttendance)
	protected.Get("/attendances", controllers.GetAttendances)
	protected.Get("/attendances/:id", controllers.GetAttendanceByID)
	protected.Put("/attendances/:id", controllers.UpdateAttendance)
	protected.Delete("/attendances/:id", controllers.DeleteAttendance)

	// Leave Request routes
	protected.Post("/leave-requests", controllers.CreateLeaveRequest)
	protected.Get("/leave-requests", controllers.GetLeaveRequests)
	protected.Get("/leave-requests/:id", controllers.GetLeaveRequestByID)
	protected.Put("/leave-requests/:id/status", controllers.UpdateLeaveStatus)

	// Role Management routes (Admin only)
	protected.Post("/roles", controllers.CreateRole)
	protected.Get("/roles", controllers.GetRoles)
	protected.Get("/roles/:id", controllers.GetRoleByID)
	protected.Put("/roles/:id", controllers.UpdateRole)
	protected.Delete("/roles/:id", controllers.DeleteRole)
	protected.Post("/users/create-with-role", controllers.CreateUserWithRole)
	protected.Post("/assign-role", controllers.AssignUserRole)

	// Payroll View (All authenticated users can view) - MUST BE BEFORE RBAC ROUTES
	protected.Get("/allowances", controllers.GetAllowances)
	protected.Get("/deductions", controllers.GetDeductions)
	protected.Get("/payrolls", controllers.GetPayrolls)

	// Payroll routes (Enhanced) - HR/Admin Only - SPECIFIC ROUTES FIRST
	payrollRoutes := protected.Group("/", middleware.RequireRoles(config.DB, "admin", "hr"))
	
	// Allowances (HR/Admin Only)
	payrollRoutes.Post("/allowances", controllers.CreateAllowance)
	payrollRoutes.Post("/allowances/bulk", controllers.CreateBulkAllowances)
	payrollRoutes.Delete("/allowances/:id", controllers.DeleteAllowance)
	
	// Deductions (HR/Admin Only)
	payrollRoutes.Post("/deductions", controllers.CreateDeduction)
	payrollRoutes.Post("/deductions/bulk", controllers.CreateBulkDeductions)
	payrollRoutes.Delete("/deductions/:id", controllers.DeleteDeduction)
	
	// Payroll Management (HR/Admin Only) - SPECIFIC ROUTES FIRST
	payrollRoutes.Post("/payroll/generate", controllers.GeneratePayroll)
	payrollRoutes.Get("/payroll/summary", controllers.GetPayrollSummary)
	payrollRoutes.Get("/payroll/dashboard", controllers.GetPayrollDashboard)
	payrollRoutes.Get("/payroll/export/csv", controllers.ExportPayrollCSV)
	payrollRoutes.Put("/payroll/:id", controllers.UpdatePayroll)
	payrollRoutes.Patch("/payroll/:id/status", controllers.UpdatePayrollStatus)
	
	// Payroll Detail (All authenticated users) - PARAMETER ROUTE LAST
	protected.Get("/payroll/:employee_id", controllers.GetPayrollDetail)

	// KPI Management (HR/Admin Only for CRUD)
	kpiRoutes := protected.Group("/", middleware.RequireRoles(config.DB, "admin", "hr"))
	
	// KPI Categories
	kpiRoutes.Post("/kpi/categories", controllers.CreateKPICategory)
	
	// KPI Metrics
	kpiRoutes.Post("/kpi/metrics", controllers.CreateKPIMetric)
	
	// KPI Targets
	kpiRoutes.Post("/kpi/targets", controllers.SetKPITarget)
	kpiRoutes.Post("/kpi/targets/bulk", controllers.SetBulkKPITargets)
	
	// KPI Actuals
	kpiRoutes.Post("/kpi/actuals", controllers.RecordKPIActual)
	
	// KPI Views (Manager & Above can view)
	kpiViewRoutes := protected.Group("/", middleware.RequireRoles(config.DB, "admin", "hr", "manager"))
	kpiViewRoutes.Get("/kpi/categories", controllers.GetKPICategories)
	kpiViewRoutes.Get("/kpi/metrics", controllers.GetKPIMetrics)
	kpiViewRoutes.Get("/kpi/targets", controllers.GetKPITargets)
	kpiViewRoutes.Get("/kpi/actuals", controllers.GetKPIActuals)
	kpiViewRoutes.Get("/kpi/report", controllers.GetKPIReport)
	
	// KPI Dashboard (All authenticated users can view their own)
	protected.Get("/kpi/dashboard", controllers.GetKPIDashboard)


	
	// Admin only routes
	adminRoutes := protected.Group("/", middleware.RequireRoles(config.DB, "admin"))
	adminRoutes.Delete("/employees/:id", controllers.DeleteEmployee)
	
	// HR can also link employees to users
	hrRoutes := protected.Group("/", middleware.RequireRoles(config.DB, "admin", "hr"))
	hrRoutes.Post("/employees/:id/link-user", controllers.LinkEmployeeToUser)

}

func getProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
