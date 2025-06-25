package controllers

import (
	"hris-backend/config"
	"hris-backend/models"
	"hris-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateLeaveRequest(c *fiber.Ctx) error {
	var input models.LeaveRequest
	if err := c.BodyParser(&input); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}
	if input.EmployeeID == "" || input.LeaveTypeID == 0 || input.StartDate.IsZero() || input.EndDate.IsZero() {
		return utils.ValidationErrorResponse(c, "Employee ID, Leave Type ID, Start Date, and End Date are required")
	}
	input.Status = "pending"
	if err := config.DB.Create(&input).Error; err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	return utils.SuccessResponse(c, "Leave request created successfully", input)
}

func GetLeaveRequests(c *fiber.Ctx) error {
	var leaves []models.LeaveRequest
	if err := config.DB.Preload("Employee").Preload("LeaveType").Find(&leaves).Error; err != nil {
		return utils.InternalErrorResponse(c, err.Error())
	}
	return utils.SuccessResponse(c, "Leave requests retrieved successfully", leaves)
}

func GetLeaveRequestByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var leave models.LeaveRequest
	if err := config.DB.Preload("Employee").Preload("LeaveType").First(&leave, id).Error; err != nil {
		return utils.NotFoundResponse(c, "Leave request not found")
	}
	return utils.SuccessResponse(c, "Leave request retrieved successfully", leave)
}

func UpdateLeaveStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var leave models.LeaveRequest
	if err := config.DB.First(&leave, id).Error; err != nil {
		return utils.NotFoundResponse(c, "Leave request not found")
	}

	var data struct {
		Status string `json:"status"`
		Note   string `json:"note"`
	}

	if err := c.BodyParser(&data); err != nil {
		return utils.ValidationErrorResponse(c, "Invalid input format")
	}

	leave.Status = data.Status
	leave.Note = data.Note
	config.DB.Save(&leave)
	return utils.SuccessResponse(c, "Leave request status updated successfully", leave)
}
