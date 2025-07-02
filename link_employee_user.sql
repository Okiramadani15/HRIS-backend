-- Script untuk menghubungkan Employee dengan User berdasarkan email
-- Asumsi: Employee dengan employee_id "EMP001" terhubung dengan user email "employee@company.com"

-- Update Employee EMP001 untuk terhubung dengan user employee@company.com
UPDATE employees 
SET user_id = (SELECT id FROM users WHERE email = 'employee@company.com' LIMIT 1)
WHERE employee_id = 'EMP001';

-- Update Employee EMP002 jika ada
UPDATE employees 
SET user_id = (SELECT id FROM users WHERE email = 'employee2@company.com' LIMIT 1)
WHERE employee_id = 'EMP002';

-- Update Employee EMP003 jika ada  
UPDATE employees 
SET user_id = (SELECT id FROM users WHERE email = 'employee3@company.com' LIMIT 1)
WHERE employee_id = 'EMP003';

-- Verifikasi hasil
SELECT e.employee_id, e.full_name, u.email, u.name as user_name
FROM employees e 
LEFT JOIN users u ON e.user_id = u.id
WHERE e.user_id IS NOT NULL;