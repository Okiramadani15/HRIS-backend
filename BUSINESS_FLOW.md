# 🏢 HRIS System - Business Flow Documentation

## 📋 System Overview
Sistem HRIS (Human Resource Information System) yang komprehensif dengan fitur:
- User Management & RBAC
- Employee Management
- Master Data Management
- Attendance Tracking
- Leave Management
- Payroll Processing
- KPI Performance Management

---

## 🔐 Role-Based Access Control (RBAC)

### **Admin** - System Administrator
- ✅ Full system access
- ✅ User & role management
- ✅ System configuration
- ✅ All HR functions

### **HR** - Human Resources Manager
- ✅ Employee management
- ✅ Master data management
- ✅ Payroll processing
- ✅ KPI management
- ✅ Leave approvals
- ✅ Attendance oversight

### **Manager** - Department Manager
- ✅ View team KPI data
- ✅ Team performance reports
- ✅ Read-only access to employee data
- ✅ Leave request approvals (team)

### **Employee** - Regular Employee
- ✅ Self-service portal
- ✅ View own profile & payroll
- ✅ Submit attendance
- ✅ Request leave
- ✅ View own KPI dashboard

---

## 🔄 Complete Business Flow

### **Phase 1: System Setup (Admin)**
```
1. Register Admin User
2. Login as Admin
3. Create Roles (admin, hr, manager, employee)
4. Create HR Manager User
5. Create Department Managers
6. Setup Master Data (departments, positions, locations, etc.)
```

### **Phase 2: HR Operations (HR Manager)**
```
1. Login as HR
2. Setup Master Data:
   - Departments
   - Positions
   - Work Shifts
   - Leave Types
   - Pay Grades
3. Create Employee Records
4. Setup KPI Framework:
   - KPI Categories
   - KPI Metrics
   - Set Employee Targets
```

### **Phase 3: Daily Operations (All Users)**
```
Employees:
- Clock in/out (Attendance)
- Submit leave requests
- View payroll slips
- Check KPI dashboard

Managers:
- Review team performance
- Approve leave requests
- Monitor KPI achievements

HR:
- Process payroll
- Record KPI actuals
- Generate reports
- Manage employee lifecycle
```

### **Phase 4: Monthly/Periodic Tasks (HR)**
```
1. Generate monthly payroll
2. Process allowances & deductions
3. Record KPI achievements
4. Generate performance reports
5. Conduct performance reviews
```

---

## 📊 Data Flow Relationships

### **Employee Lifecycle**
```
User Registration → Role Assignment → Employee Profile Creation → 
Department Assignment → KPI Target Setting → Daily Operations
```

### **Attendance Flow**
```
Employee Clock In → Attendance Record → HR Review → 
Payroll Calculation → Monthly Reports
```

### **Leave Management Flow**
```
Employee Request → Manager/HR Approval → Leave Balance Update → 
Attendance Impact → Payroll Adjustment
```

### **Payroll Flow**
```
Base Salary + Allowances - Deductions - Leave Deductions = 
Net Salary → Payroll Generation → Payment Processing
```

### **KPI Flow**
```
Category Creation → Metric Definition → Target Setting → 
Performance Recording → Achievement Calculation → 
Grade Assignment → Reports Generation
```

---

## 🎯 Key Business Rules

### **Authentication & Authorization**
- All endpoints require valid JWT token
- Business operations require appropriate role
- Users can only access their own data (employees)
- Managers can view team data
- HR can access all employee data

### **Employee Management**
- Unique employee ID required
- Email must be unique
- Active status for current employees
- Department and position assignment mandatory

### **Attendance Rules**
- Daily attendance recording
- Clock in/out timestamps
- Overtime calculation
- Leave impact on attendance

### **Leave Management**
- Leave balance tracking
- Approval workflow
- Leave type restrictions
- Annual leave entitlements

### **Payroll Rules**
- Monthly payroll processing
- Automatic allowance/deduction calculation
- Tax calculations
- Leave impact on salary

### **KPI Management**
- Target vs Actual tracking
- Weighted scoring system
- Grade calculation (A-E)
- Performance periods (monthly/quarterly/annual)

---

## 📈 Reporting & Analytics

### **HR Dashboard**
- Employee count by department
- Attendance statistics
- Leave utilization
- Payroll summaries
- KPI achievements

### **Manager Dashboard**
- Team performance metrics
- Attendance overview
- Leave requests pending
- KPI progress tracking

### **Employee Dashboard**
- Personal attendance record
- Leave balance
- Payroll history
- KPI achievements
- Performance grade

---

## 🔧 Technical Implementation

### **Database Design**
- Normalized relational structure
- Foreign key relationships
- Audit trails (created_at, updated_at)
- Soft deletes where applicable

### **API Design**
- RESTful endpoints
- Consistent response format
- Proper HTTP status codes
- Input validation
- Error handling

### **Security**
- JWT-based authentication
- Role-based authorization
- Password hashing
- Input sanitization
- SQL injection prevention

---

## 🚀 Deployment Considerations

### **Environment Setup**
- Development, Staging, Production
- Database migrations
- Environment variables
- Logging configuration

### **Monitoring**
- API performance metrics
- Database query optimization
- Error tracking
- User activity logs

### **Backup & Recovery**
- Regular database backups
- Data retention policies
- Disaster recovery plan
- System health checks