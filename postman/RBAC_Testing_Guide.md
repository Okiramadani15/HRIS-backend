# 🛡️ RBAC Testing Guide untuk HRIS API

## 📋 **Testing Scenarios Overview**

### **1. Setup Roles & Users**
- Create 3 roles: admin, hr, employee
- Create users dengan masing-masing role
- Auto-save user IDs untuk testing selanjutnya

### **2. Login Different Roles**
- Login sebagai Admin → Save ADMIN_TOKEN
- Login sebagai HR → Save HR_TOKEN  
- Login sebagai Employee → Save EMPLOYEE_TOKEN

### **3. Test Admin Access (Full Access)**
- ✅ Create Employee (Should Work)
- ✅ Delete Employee (Should Work)
- ✅ Manage Roles (Should Work)

### **4. Test HR Access (Limited Access)**
- ✅ Create Employee (Should Work)
- ❌ Delete Employee (Should Fail - Admin Only)
- ✅ View Employees (Should Work)

### **5. Test Employee Access (Read Only)**
- ❌ Create Employee (Should Fail)
- ❌ Delete Employee (Should Fail)
- ✅ View Employees (Should Work)
- ✅ View Own Profile (Should Work)

### **6. Test Role Assignment**
- ✅ Admin can assign roles (Should Work)
- ❌ Employee cannot manage roles (Should Fail)

### **7. Test Invalid Access**
- ❌ No Token (Should Fail)
- ❌ Invalid Token (Should Fail)

## 🚀 **How to Run Tests**

### **Step 1: Import Collection**
1. Open Postman
2. Import `RBAC_Testing_Collection.json`
3. Set environment variables:
   - BASE_URL: `http://localhost:3000`

### **Step 2: Run Tests in Order**
1. **Setup Roles & Users** - Jalankan semua requests
2. **Login Different Roles** - Dapatkan tokens
3. **Test Admin Access** - Verify admin permissions
4. **Test HR Access** - Verify HR permissions  
5. **Test Employee Access** - Verify employee permissions
6. **Test Role Assignment** - Test role management
7. **Test Invalid Access** - Test security

### **Step 3: Expected Results**

#### **Admin Role (Full Access)**
```json
// Create Employee - 201 Created
{
    "message": "Employee created successfully",
    "employee": {...}
}

// Delete Employee - 200 OK
{
    "message": "Employee deleted successfully"
}
```

#### **HR Role (Limited Access)**
```json
// Create Employee - 201 Created
{
    "message": "Employee created successfully",
    "employee": {...}
}

// Delete Employee - 403 Forbidden
{
    "success": false,
    "error": "Forbidden - insufficient role permissions"
}
```

#### **Employee Role (Read Only)**
```json
// Create Employee - 403 Forbidden
{
    "success": false,
    "error": "Forbidden - insufficient role permissions"
}

// View Employees - 200 OK
{
    "employees": [...],
    "total": 5
}
```

## 🎯 **Key Testing Points**

### **1. Role Hierarchy**
- **Admin**: Full access (CRUD + Role Management)
- **HR**: Employee management (CR + Read)
- **Employee**: Read-only access

### **2. Protected Endpoints**
- `POST /api/employees` - HR + Admin only
- `DELETE /api/employees/:id` - Admin only
- `POST /api/roles` - Admin only
- `POST /api/assign-role` - Admin only

### **3. Public Endpoints**
- `POST /api/register` - No auth required
- `POST /api/login` - No auth required

### **4. JWT Protected Endpoints**
- All other endpoints require valid JWT token
- Token contains user role information

## 🔍 **Troubleshooting**

### **Common Issues:**

1. **403 Forbidden**
   - User doesn't have required role
   - Check if user has correct role assigned

2. **401 Unauthorized**
   - Invalid or expired JWT token
   - Re-login to get fresh token

3. **Role Not Found**
   - Roles not created yet
   - Run "Setup Roles & Users" first

### **Debug Steps:**
1. Check server logs for RBAC middleware
2. Verify JWT token contains role information
3. Confirm user has correct role in database

## 📊 **Test Results Matrix**

| Endpoint | Admin | HR | Employee |
|----------|-------|----|---------| 
| GET /employees | ✅ | ✅ | ✅ |
| POST /employees | ✅ | ✅ | ❌ |
| DELETE /employees/:id | ✅ | ❌ | ❌ |
| POST /roles | ✅ | ❌ | ❌ |
| POST /assign-role | ✅ | ❌ | ❌ |
| GET /my-profile | ✅ | ✅ | ✅ |

## 🎓 **Learning Objectives**

Setelah menjalankan tests ini, Anda akan memahami:
1. **Role-based Access Control** implementation
2. **JWT token** dengan role information
3. **Middleware** untuk permission checking
4. **Security layers** dalam API design
5. **Error handling** untuk unauthorized access

Happy Testing! 🚀