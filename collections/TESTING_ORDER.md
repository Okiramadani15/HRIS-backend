# 🎯 HRIS Testing Order & Flow Guide

## 📋 Testing Sequence

### **WAJIB DIIKUTI URUTAN INI:**

---

## **1️⃣ AUTHENTICATION (PERTAMA KALI)**
**File**: `01_Authentication.json`
**Tujuan**: Setup sistem dan user management
**Durasi**: ~3 menit

### **Urutan Eksekusi:**
1. **Register Admin User** ✅
2. **Login Admin** ✅ (Token tersimpan otomatis)
3. **Assign Admin Role** ✅
4. **Get All Roles** ✅ (Verifikasi roles tersedia)
5. **Create HR Manager User** ✅
6. **Create Department Manager** ✅
7. **Create Employee User** ✅

### **✅ Success Criteria:**
- Semua users berhasil dibuat
- Admin token tersimpan
- Roles ter-assign dengan benar

---

## **2️⃣ HR MASTER DATA**
**File**: `02_HR_MasterData.json`
**Tujuan**: Setup master data organisasi
**Durasi**: ~4 menit

### **Urutan Eksekusi:**
1. **Login HR Manager** ✅ (Token tersimpan)
2. **Departments**: Create Sales, HR, IT
3. **Positions**: Create Manager, Senior Staff, Staff
4. **Locations**: Create Jakarta Office, Surabaya Branch
5. **Work Shifts**: Create Regular, Flexible
6. **Leave Types**: Create Annual, Sick Leave

### **✅ Success Criteria:**
- Semua master data terbuat
- HR token tersimpan
- Data siap untuk employee creation

---

## **3️⃣ EMPLOYEE MANAGEMENT**
**File**: `03_Employee_Management.json`
**Tujuan**: Buat data karyawan
**Durasi**: ~2 menit

### **Urutan Eksekusi:**
1. **Login HR Manager** ✅
2. **Create Sales Employee** (EMP001)
3. **Create HR Employee** (EMP002)
4. **Create IT Employee** (EMP003)
5. **Get All Employees** ✅
6. **Update Employee** (Optional)

### **✅ Success Criteria:**
- 3 employees terbuat dengan ID: EMP001, EMP002, EMP003
- Data employee lengkap dan konsisten

---

## **4️⃣ KPI MANAGEMENT**
**File**: `07_KPI_Management.json`
**Tujuan**: Setup performance management
**Durasi**: ~5 menit

### **Urutan Eksekusi:**
1. **Setup Tokens**: Login HR, Manager, Employee
2. **KPI Categories**: Create 3 categories
3. **KPI Metrics**: Create 4 metrics
4. **KPI Targets**: Set individual & bulk targets
5. **KPI Actuals**: Record achievements
6. **KPI Views**: Test manager access
7. **Reports & Dashboard**: Generate reports

### **✅ Success Criteria:**
- KPI framework lengkap
- Targets dan actuals terecord
- Dashboard accessible

---

## **5️⃣ ATTENDANCE MANAGEMENT**
**File**: `04_Attendance_Management.json`
**Tujuan**: Test attendance system
**Durasi**: ~3 menit

### **Urutan Eksekusi:**
1. **Setup Tokens**: Login Employee & HR
2. **Employee Self-Service**: Clock in/out, full day attendance
3. **HR Management**: View, update, reports

### **✅ Success Criteria:**
- Attendance records created
- HR dapat manage attendance
- Reports accessible

---

## **6️⃣ LEAVE MANAGEMENT**
**File**: `05_Leave_Management.json`
**Tujuan**: Test leave system
**Durasi**: ~3 menit

### **Urutan Eksekusi:**
1. **Setup Tokens**: Login Employee & HR
2. **Employee Leave Requests**: Submit various leave types
3. **HR Leave Management**: Approve/reject, view history

### **✅ Success Criteria:**
- Leave requests submitted
- HR approval workflow working
- Leave history accessible

---

## **7️⃣ PAYROLL MANAGEMENT**
**File**: `06_Payroll_Management.json`
**Tujuan**: Test payroll system
**Durasi**: ~4 menit

### **Urutan Eksekusi:**
1. **Setup Tokens**: Login HR & Employee
2. **Allowances Management**: Create individual & bulk
3. **Deductions Management**: Create BPJS, tax, bulk
4. **Payroll Processing**: Generate, summary, dashboard

### **✅ Success Criteria:**
- Allowances & deductions setup
- Payroll generated successfully
- Employee dapat view payslip

---

## 🚀 **QUICK TESTING (15 Menit)**

### **Minimal Flow:**
1. **01_Authentication** → Run semua (3 min)
2. **02_HR_MasterData** → Run "Login HR" + "Departments" saja (1 min)
3. **03_Employee_Management** → Run "Login HR" + Create 1 employee (1 min)
4. **07_KPI_Management** → Run "Setup Tokens" + "KPI Categories" (2 min)
5. **04_Attendance** → Run "Setup Tokens" + 1 attendance (1 min)
6. **06_Payroll** → Run "Setup Tokens" + Generate payroll (2 min)

---

## 🔍 **COMPLETE TESTING (25 Menit)**

### **Full Flow:**
Run semua collections sesuai urutan 1-7 secara lengkap

---

## ⚠️ **IMPORTANT NOTES:**

### **❌ JANGAN SKIP:**
- **Authentication WAJIB pertama** - tanpa ini semua akan gagal
- **HR Master Data** - diperlukan untuk employee creation
- **Employee Management** - diperlukan untuk KPI, attendance, payroll

### **✅ TIPS SUKSES:**
- **Ikuti urutan** - jangan loncat-loncat
- **Check console** - lihat ✅ atau ❌ di setiap step
- **Token auto-save** - tidak perlu copy-paste manual
- **Consistent data** - gunakan EMP001, EMP002, EMP003

### **🔧 TROUBLESHOOTING:**
- **403 Forbidden** → User belum punya role, jalankan Authentication dulu
- **404 Not Found** → Data belum ada, jalankan prerequisite collections
- **500 Error** → Check server logs, mungkin ada database issue

---

## 📊 **SUCCESS INDICATORS:**

### **✅ System Working If:**
- All tokens auto-saved
- No 403/404 errors
- Data flows between modules
- Reports generate correctly
- Console shows ✅ for each step

### **🎯 Final Validation:**
- **Users**: 4 users with different roles
- **Employees**: 3 employee records
- **KPI**: Categories, metrics, targets, actuals
- **Attendance**: Clock in/out records
- **Payroll**: Generated with allowances/deductions

**Happy Testing! 🎉**