package services

import (
	"hris-backend/config"
	"hris-backend/models"

	"gorm.io/gorm"
)

// DatabaseService provides optimized database operations using pointers
type DatabaseService struct {
	db *gorm.DB
}

// NewDatabaseService creates a new database service instance
func NewDatabaseService() *DatabaseService {
	if config.DB == nil {
		panic("Database connection is not initialized")
	}
	return &DatabaseService{db: config.DB}
}

// Employee operations with pointer optimization
func (s *DatabaseService) CreateEmployee(employee *models.Employee) error {
	if employee == nil {
		panic("Employee pointer cannot be nil")
	}
	return s.db.Create(employee).Error
}

func (s *DatabaseService) GetEmployeeByID(id uint, employee *models.Employee) error {
	if employee == nil {
		panic("Employee pointer cannot be nil")
	}
	return s.db.Preload("User").First(employee, id).Error
}

func (s *DatabaseService) GetEmployees(employees *[]models.Employee) error {
	if employees == nil {
		panic("Employees slice pointer cannot be nil")
	}
	return s.db.Preload("User").Find(employees).Error
}

func (s *DatabaseService) UpdateEmployee(employee *models.Employee) error {
	if employee == nil {
		panic("Employee pointer cannot be nil")
	}
	return s.db.Save(employee).Error
}

func (s *DatabaseService) DeleteEmployee(id uint) error {
	result := s.db.Delete(&models.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		panic("Employee not found for deletion")
	}
	return nil
}

// Attendance operations with pointer optimization
func (s *DatabaseService) CreateAttendance(attendance *models.Attendance) error {
	if attendance == nil {
		panic("Attendance pointer cannot be nil")
	}
	return s.db.Create(attendance).Error
}

func (s *DatabaseService) GetAttendances(attendances *[]models.Attendance) error {
	if attendances == nil {
		panic("Attendances slice pointer cannot be nil")
	}
	return s.db.Preload("Employee").Find(attendances).Error
}

func (s *DatabaseService) GetAttendancesWithFilters(attendances *[]models.Attendance, date, employeeID, month string) error {
	if attendances == nil {
		panic("Attendances slice pointer cannot be nil")
	}
	
	query := s.db.Preload("Employee")
	
	// Filter by specific date
	if date != "" {
		query = query.Where("DATE(date) = ?", date)
	}
	
	// Filter by employee ID
	if employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}
	
	// Filter by month (format: YYYY-MM)
	if month != "" {
		query = query.Where("DATE_FORMAT(date, '%Y-%m') = ?", month)
	}
	
	return query.Find(attendances).Error
}

func (s *DatabaseService) GetAttendanceByID(id uint, attendance *models.Attendance) error {
	if attendance == nil {
		panic("Attendance pointer cannot be nil")
	}
	return s.db.Preload("Employee").First(attendance, id).Error
}

func (s *DatabaseService) UpdateAttendance(attendance *models.Attendance) error {
	if attendance == nil {
		panic("Attendance pointer cannot be nil")
	}
	return s.db.Save(attendance).Error
}

func (s *DatabaseService) DeleteAttendance(id uint) error {
	result := s.db.Delete(&models.Attendance{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		panic("Attendance not found for deletion")
	}
	return nil
}

// LeaveRequest operations with pointer optimization
func (s *DatabaseService) CreateLeaveRequest(leaveRequest *models.LeaveRequest) error {
	if leaveRequest == nil {
		panic("LeaveRequest pointer cannot be nil")
	}
	return s.db.Create(leaveRequest).Error
}

func (s *DatabaseService) GetLeaveRequests(leaveRequests *[]models.LeaveRequest) error {
	if leaveRequests == nil {
		panic("LeaveRequests slice pointer cannot be nil")
	}
	return s.db.Preload("Employee").Preload("LeaveType").Find(leaveRequests).Error
}

func (s *DatabaseService) GetLeaveRequestByID(id uint, leaveRequest *models.LeaveRequest) error {
	if leaveRequest == nil {
		panic("LeaveRequest pointer cannot be nil")
	}
	return s.db.Preload("Employee").Preload("LeaveType").First(leaveRequest, id).Error
}

func (s *DatabaseService) UpdateLeaveRequest(leaveRequest *models.LeaveRequest) error {
	if leaveRequest == nil {
		panic("LeaveRequest pointer cannot be nil")
	}
	return s.db.Save(leaveRequest).Error
}