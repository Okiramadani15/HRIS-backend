# 🧪 HRIS API Testing dengan Postman

## 📁 Files yang Tersedia

- `HRIS_API_Collection.json` - Postman Collection dengan semua endpoint
- `Environment.json` - Environment variables untuk testing
- `README.md` - Panduan penggunaan

## 🚀 Cara Import ke Postman

### 1. Import Collection
1. Buka Postman
2. Click **Import** button
3. Pilih file `HRIS_API_Collection.json`
4. Collection akan muncul di sidebar

### 2. Import Environment
1. Click **Environments** tab
2. Click **Import**
3. Pilih file `Environment.json`
4. Set environment sebagai **active**

## 🔧 Setup Testing

### 1. Start Server
```bash
cd hris-backend
go run cmd/main.go
```

### 2. Testing Flow
1. **Register User** - Buat user baru
2. **Login User** - Dapatkan JWT token (otomatis tersimpan)
3. **Create Master Data** - Buat data referensi
4. **Test CRUD Operations** - Test semua endpoint

## 📋 Testing Scenarios

### ✅ Authentication Flow
- [x] Register new user
- [x] Login and get JWT token
- [x] Access protected endpoints

### ✅ Employee Management
- [x] Create employee (HR/Admin only)
- [x] Get all employees
- [x] Get employee by ID
- [x] Update employee
- [x] Delete employee (Admin only)

### ✅ Attendance Management
- [x] Create attendance record
- [x] Get all attendances
- [x] Update attendance
- [x] Delete attendance

### ✅ Leave Management
- [x] Create leave types
- [x] Create leave requests
- [x] Update leave status

## 🎯 Expected Responses

### Success Response
```json
{
    "success": true,
    "message": "Operation successful",
    "data": {...}
}
```

### Error Response
```json
{
    "success": false,
    "error": "Error message"
}
```

## 🔐 Authentication

Semua protected endpoints memerlukan header:
```
Authorization: Bearer <JWT_TOKEN>
```

Token akan otomatis tersimpan setelah login berhasil.

## 🚨 Error Testing

Collection juga include testing untuk:
- Invalid JWT token
- Missing required fields
- Panic recovery
- RBAC permissions

## 📊 Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `BASE_URL` | Server URL | `http://localhost:3000` |
| `JWT_TOKEN` | Authentication token | Auto-filled after login |
| `ADMIN_EMAIL` | Admin email | `admin@hris.com` |
| `ADMIN_PASSWORD` | Admin password | `password123` |

## 🎓 Tips

1. **Run tests in order** - Authentication first, then CRUD operations
2. **Check environment** - Pastikan environment aktif
3. **Monitor console** - Lihat logs untuk debugging
4. **Use variables** - Gunakan `{{variable}}` untuk dynamic values

Happy Testing! 🚀