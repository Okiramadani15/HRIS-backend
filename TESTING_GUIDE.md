# 🧪 HRIS System Testing Guide

## 📋 Overview
Panduan lengkap untuk testing sistem HRIS mengikuti business flow yang sebenarnya.

---

## 🚀 Quick Start (5 Menit)

### **1. Persiapan**
```bash
# Start server
cd hris-backend
go run cmd/main.go

# Import collection ke Postman
# File: COMPLETE_HRIS_COLLECTION.json
```

### **2. Run Complete Test**
1. Import collection ke Postman
2. Buka Collection Runner
3. Select "HRIS System - Complete Business Flow"
4. Run entire collection
5. Monitor console untuk progress

---

## 📊 Testing Phases

### **🔐 Phase 1: System Setup (Admin)**
**Tujuan**: Setup sistem dan user management
**Duration**: ~2 menit

**Steps**:
1. Register Admin User
2. Login Admin
3. Assign Admin Role
4. Create HR Manager
5. Create Department Manager
6. Create Employee User

**Expected Results**:
- ✅ All users created successfully
- ✅ Roles assigned properly
- ✅ Authentication working

---

### **🏢 Phase 2: HR Operations Setup**
**Tujuan**: Setup master data dan struktur organisasi
**Duration**: ~3 menit

**Steps**:
1. Login HR Manager
2. Create Departments (Sales, HR, IT)
3. Create Positions (Manager, Senior Staff, Staff)
4. Create Work Shifts (Regular, Flexible)
5. Create Leave Types (Annual, Sick)
6. Create Employee Records

**Expected Results**:
- ✅ Master data created
- ✅ Employee records established
- ✅ Organizational structure ready

---

### **📊 Phase 3: KPI Framework Setup**
**Tujuan**: Setup performance management system
**Duration**: ~2 menit

**Steps**:
1. Create KPI Categories
2. Create KPI Metrics
3. Set Employee KPI Targets

**Expected Results**:
- ✅ KPI framework established
- ✅ Performance targets set
- ✅ Metrics properly configured

---

### **👥 Phase 4: Daily Operations Testing**
**Tujuan**: Test daily user operations
**Duration**: ~3 menit

**Employee Operations**:
- Login Employee
- Clock In/Out
- Submit Leave Request
- View KPI Dashboard

**Manager Operations**:
- Login Manager
- View KPI Data
- Monitor Team Performance

**HR Operations**:
- Approve Leave Requests
- Record KPI Actuals
- Monitor System

**Expected Results**:
- ✅ All user roles working
- ✅ Daily operations functional
- ✅ RBAC properly enforced

---

### **💰 Phase 5: Payroll Processing**
**Tujuan**: Test payroll system
**Duration**: ~2 menit

**Steps**:
1. Setup Allowances (Transport, Bonus)
2. Setup Deductions (BPJS, Tax)
3. Generate Monthly Payroll
4. View Payroll Summary

**Expected Results**:
- ✅ Payroll components configured
- ✅ Payroll generation working
- ✅ Calculations accurate

---

### **📈 Phase 6: Reports & Analytics**
**Tujuan**: Test reporting capabilities
**Duration**: ~2 menit

**Reports Tested**:
- KPI Reports
- Payroll Reports
- Attendance Reports
- Employee Dashboards

**Expected Results**:
- ✅ All reports generating
- ✅ Data accuracy verified
- ✅ Role-based access working

---

### **🔍 Phase 7: System Validation**
**Tujuan**: Validate system integrity and security
**Duration**: ~1 menit

**Validation Tests**:
- Data integrity checks
- RBAC validation
- Security testing

**Expected Results**:
- ✅ All data consistent
- ✅ RBAC blocking unauthorized access
- ✅ System security validated

---

## 🎯 Success Criteria

### **✅ Complete Success Indicators**
- All 7 phases pass without errors
- Console shows ✅ for each step
- Final message: "🎉 COMPLETE SYSTEM TEST SUCCESSFUL!"
- No 500 errors in any request
- RBAC properly blocks unauthorized access

### **📊 Key Metrics to Verify**
- **Users**: 4 users with different roles
- **Employees**: 2+ employee records
- **Departments**: 3+ departments
- **KPI Categories**: 2+ categories
- **KPI Metrics**: 2+ metrics
- **Payroll Records**: Generated successfully
- **Reports**: All accessible by appropriate roles

---

## 🔧 Troubleshooting

### **Common Issues**

**❌ "Invalid user context"**
```
Solution:
1. Ensure user has role assigned
2. Check token validity
3. Verify RBAC middleware
```

**❌ "Employee not found"**
```
Solution:
1. Run Phase 2 completely
2. Verify employee creation
3. Check employee_id consistency
```

**❌ "Forbidden - insufficient role permissions"**
```
Solution:
1. Verify user role assignment
2. Check endpoint permissions
3. Use correct token for role
```

**❌ "Failed to generate payroll"**
```
Solution:
1. Ensure employees exist
2. Verify allowances/deductions setup
3. Check period format (YYYY-MM)
```

---

## 📋 Manual Testing Checklist

### **Pre-Testing**
- [ ] Server running on localhost:3000
- [ ] Database connected
- [ ] Postman collection imported
- [ ] Environment variables set

### **During Testing**
- [ ] Monitor console output
- [ ] Check response status codes
- [ ] Verify data consistency
- [ ] Test RBAC restrictions

### **Post-Testing**
- [ ] All phases completed
- [ ] No error responses
- [ ] Data relationships intact
- [ ] Security validations passed

---

## 🎮 Testing Modes

### **🏃‍♂️ Quick Test (5 min)**
Run entire collection with Collection Runner

### **🔍 Detailed Test (15 min)**
Run each phase individually, verify responses

### **🧪 Development Test (30 min)**
Manual testing with custom data and edge cases

### **🚀 Production Readiness (60 min)**
Complete testing with performance and security validation

---

## 📞 Support

### **If Tests Fail**
1. Check server logs for errors
2. Verify database connection
3. Ensure all dependencies installed
4. Check environment configuration

### **For Custom Testing**
1. Modify collection variables
2. Adjust test data as needed
3. Add custom validation scripts
4. Create additional test scenarios

**Happy Testing! 🎉**