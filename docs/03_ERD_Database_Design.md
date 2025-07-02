# 🗄️ ENTITY RELATIONSHIP DIAGRAM (ERD)
## HRIS Database Design

---

## 📊 DATABASE OVERVIEW

### Database Type: PostgreSQL
### Total Tables: 20+
### Relationships: 25+ Foreign Keys
### Indexes: Optimized for performance

---

## 🏗️ CORE ENTITIES

### 1. USERS TABLE
```sql
users {
  id SERIAL PRIMARY KEY
  name VARCHAR(255) NOT NULL
  email VARCHAR(255) UNIQUE NOT NULL
  password VARCHAR(255) NOT NULL
  role_id INTEGER REFERENCES roles(id)
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
  deleted_at TIMESTAMP NULL
}
```

### 2. ROLES TABLE
```sql
roles {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) UNIQUE NOT NULL -- admin, hr, manager, employee
  description TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 3. EMPLOYEES TABLE
```sql
employees {
  id SERIAL PRIMARY KEY
  user_id INTEGER UNIQUE REFERENCES users(id)
  employee_id VARCHAR(50) UNIQUE NOT NULL -- EMP001, EMP002
  full_name VARCHAR(255) NOT NULL
  phone VARCHAR(20)
  address TEXT
  date_of_birth DATE
  gender VARCHAR(10)
  position VARCHAR(100)
  department VARCHAR(100)
  hire_date DATE
  salary DECIMAL(15,2)
  status VARCHAR(20) DEFAULT 'active'
  photo_url VARCHAR(500)
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
  deleted_at TIMESTAMP NULL
}
```

---

## 🏢 ORGANIZATIONAL STRUCTURE

### 4. DEPARTMENTS TABLE
```sql
departments {
  id SERIAL PRIMARY KEY
  name VARCHAR(100) UNIQUE NOT NULL
  description TEXT
  manager_id INTEGER REFERENCES employees(id)
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
  deleted_at TIMESTAMP NULL
}
```

### 5. POSITIONS TABLE
```sql
positions {
  id SERIAL PRIMARY KEY
  name VARCHAR(100) NOT NULL
  description TEXT
  department_id INTEGER REFERENCES departments(id)
  level INTEGER DEFAULT 1
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
  deleted_at TIMESTAMP NULL
}
```

### 6. LOCATIONS TABLE
```sql
locations {
  id SERIAL PRIMARY KEY
  name VARCHAR(100) NOT NULL
  address TEXT
  city VARCHAR(50)
  country VARCHAR(50)
  timezone VARCHAR(50)
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
  deleted_at TIMESTAMP NULL
}
```

---

## ⏰ ATTENDANCE MANAGEMENT

### 7. ATTENDANCES TABLE
```sql
attendances {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  date DATE NOT NULL
  check_in TIMESTAMP
  check_out TIMESTAMP
  break_start TIMESTAMP
  break_end TIMESTAMP
  work_hours DECIMAL(4,2)
  overtime_hours DECIMAL(4,2)
  status VARCHAR(20) DEFAULT 'present'
  notes TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 8. WORK_SHIFTS TABLE
```sql
work_shifts {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL
  start_time TIME NOT NULL
  end_time TIME NOT NULL
  break_duration INTEGER DEFAULT 60 -- minutes
  description TEXT
  is_active BOOLEAN DEFAULT TRUE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

---

## 🏖️ LEAVE MANAGEMENT

### 9. LEAVE_TYPES TABLE
```sql
leave_types {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL -- annual, sick, emergency
  description TEXT
  max_days INTEGER
  is_paid BOOLEAN DEFAULT TRUE
  requires_approval BOOLEAN DEFAULT TRUE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 10. LEAVE_REQUESTS TABLE
```sql
leave_requests {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  leave_type_id INTEGER REFERENCES leave_types(id)
  start_date DATE NOT NULL
  end_date DATE NOT NULL
  days_requested INTEGER NOT NULL
  reason TEXT
  status VARCHAR(20) DEFAULT 'pending' -- pending, approved, rejected
  approved_by INTEGER REFERENCES users(id)
  approved_at TIMESTAMP
  notes TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

---

## 💰 PAYROLL SYSTEM

### 11. ALLOWANCES TABLE
```sql
allowances {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  type VARCHAR(50) NOT NULL -- transport, meal, bonus
  amount DECIMAL(15,2) NOT NULL
  description TEXT
  period VARCHAR(7) -- 2025-01
  is_taxable BOOLEAN DEFAULT TRUE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 12. DEDUCTIONS TABLE
```sql
deductions {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  type VARCHAR(50) NOT NULL -- tax, insurance, loan
  amount DECIMAL(15,2) NOT NULL
  description TEXT
  period VARCHAR(7) -- 2025-01
  is_mandatory BOOLEAN DEFAULT FALSE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 13. PAYROLLS TABLE
```sql
payrolls {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  month VARCHAR(2) NOT NULL -- 01, 02, 03
  year VARCHAR(4) NOT NULL -- 2025
  base_salary DECIMAL(15,2) NOT NULL
  total_allowances DECIMAL(15,2) DEFAULT 0
  total_deductions DECIMAL(15,2) DEFAULT 0
  net_salary DECIMAL(15,2) NOT NULL
  status VARCHAR(20) DEFAULT 'generated' -- generated, approved, paid
  paid_at TIMESTAMP
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

---

## 📈 KPI MANAGEMENT

### 14. KPI_CATEGORIES TABLE
```sql
kpi_categories {
  id SERIAL PRIMARY KEY
  name VARCHAR(100) NOT NULL
  description TEXT
  is_active BOOLEAN DEFAULT TRUE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 15. KPI_METRICS TABLE
```sql
kpi_metrics {
  id SERIAL PRIMARY KEY
  category_id INTEGER REFERENCES kpi_categories(id)
  name VARCHAR(200) NOT NULL
  description TEXT
  unit VARCHAR(20) -- %, number, currency
  target_type VARCHAR(20) -- higher_better, lower_better, exact
  weight DECIMAL(3,1) DEFAULT 1.0
  is_active BOOLEAN DEFAULT TRUE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 16. KPI_TARGETS TABLE
```sql
kpi_targets {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  metric_id INTEGER REFERENCES kpi_metrics(id)
  period VARCHAR(7) -- 2025-01
  period_type VARCHAR(20) -- monthly, quarterly, yearly
  target_value DECIMAL(15,2) NOT NULL
  min_value DECIMAL(15,2)
  max_value DECIMAL(15,2)
  set_by VARCHAR(50)
  set_at TIMESTAMP DEFAULT NOW()
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 17. KPI_ACTUALS TABLE
```sql
kpi_actuals {
  id SERIAL PRIMARY KEY
  employee_id VARCHAR(50) REFERENCES employees(employee_id)
  metric_id INTEGER REFERENCES kpi_metrics(id)
  period VARCHAR(7) -- 2025-01
  period_type VARCHAR(20) -- monthly, quarterly, yearly
  actual_value DECIMAL(15,2) NOT NULL
  notes TEXT
  recorded_by VARCHAR(50)
  recorded_at TIMESTAMP DEFAULT NOW()
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

---

## 🏷️ MASTER DATA TABLES

### 18. EDUCATIONS TABLE
```sql
educations {
  id SERIAL PRIMARY KEY
  name VARCHAR(100) NOT NULL -- SD, SMP, SMA, S1, S2, S3
  level INTEGER
  description TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 19. MARITAL_STATUSES TABLE
```sql
marital_statuses {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL -- Single, Married, Divorced, Widowed
  description TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 20. RELIGIONS TABLE
```sql
religions {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL -- Islam, Christian, Hindu, Buddhist, etc
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 21. EMPLOYMENT_TYPES TABLE
```sql
employment_types {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL -- Full-time, Part-time, Contract, Intern
  description TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 22. BANKS TABLE
```sql
banks {
  id SERIAL PRIMARY KEY
  name VARCHAR(100) NOT NULL -- BCA, Mandiri, BRI, BNI
  code VARCHAR(10) UNIQUE
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 23. PAY_GRADES TABLE
```sql
pay_grades {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL -- Grade A, Grade B, Grade C
  min_salary DECIMAL(15,2)
  max_salary DECIMAL(15,2)
  description TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

### 24. JOB_LEVELS TABLE
```sql
job_levels {
  id SERIAL PRIMARY KEY
  name VARCHAR(50) NOT NULL -- Junior, Senior, Lead, Manager
  level INTEGER UNIQUE
  description TEXT
  created_at TIMESTAMP DEFAULT NOW()
  updated_at TIMESTAMP DEFAULT NOW()
}
```

---

## 🔗 RELATIONSHIP MAPPING

### PRIMARY RELATIONSHIPS

#### User-Employee Relationship (1:1)
```
users.id ←→ employees.user_id
```

#### Employee-Department Relationship (N:1)
```
employees.department ←→ departments.name
```

#### Employee-Attendance Relationship (1:N)
```
employees.employee_id ←→ attendances.employee_id
```

#### Employee-Leave Relationship (1:N)
```
employees.employee_id ←→ leave_requests.employee_id
```

#### Employee-Payroll Relationship (1:N)
```
employees.employee_id ←→ payrolls.employee_id
```

#### Employee-KPI Relationship (1:N)
```
employees.employee_id ←→ kpi_targets.employee_id
employees.employee_id ←→ kpi_actuals.employee_id
```

### SECONDARY RELATIONSHIPS

#### Role-Based Access
```
users.role_id ←→ roles.id
```

#### KPI Hierarchy
```
kpi_categories.id ←→ kpi_metrics.category_id
kpi_metrics.id ←→ kpi_targets.metric_id
kpi_metrics.id ←→ kpi_actuals.metric_id
```

#### Leave Management
```
leave_types.id ←→ leave_requests.leave_type_id
users.id ←→ leave_requests.approved_by
```

---

## 📊 DATABASE INDEXES

### Performance Indexes
```sql
-- User authentication
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role_id ON users(role_id);

-- Employee lookups
CREATE INDEX idx_employees_employee_id ON employees(employee_id);
CREATE INDEX idx_employees_user_id ON employees(user_id);
CREATE INDEX idx_employees_department ON employees(department);

-- Attendance queries
CREATE INDEX idx_attendances_employee_date ON attendances(employee_id, date);
CREATE INDEX idx_attendances_date ON attendances(date);

-- Leave requests
CREATE INDEX idx_leave_requests_employee ON leave_requests(employee_id);
CREATE INDEX idx_leave_requests_status ON leave_requests(status);

-- Payroll processing
CREATE INDEX idx_payrolls_employee_period ON payrolls(employee_id, month, year);
CREATE INDEX idx_allowances_employee_period ON allowances(employee_id, period);
CREATE INDEX idx_deductions_employee_period ON deductions(employee_id, period);

-- KPI tracking
CREATE INDEX idx_kpi_targets_employee_period ON kpi_targets(employee_id, period);
CREATE INDEX idx_kpi_actuals_employee_period ON kpi_actuals(employee_id, period);
```

---

## 🔒 DATA CONSTRAINTS

### Business Rules
```sql
-- Employee ID format validation
ALTER TABLE employees ADD CONSTRAINT chk_employee_id 
CHECK (employee_id ~ '^EMP[0-9]{3,}$');

-- Salary must be positive
ALTER TABLE employees ADD CONSTRAINT chk_salary_positive 
CHECK (salary > 0);

-- Work hours validation
ALTER TABLE attendances ADD CONSTRAINT chk_work_hours 
CHECK (work_hours >= 0 AND work_hours <= 24);

-- Leave dates validation
ALTER TABLE leave_requests ADD CONSTRAINT chk_leave_dates 
CHECK (end_date >= start_date);

-- Payroll amounts validation
ALTER TABLE payrolls ADD CONSTRAINT chk_base_salary_positive 
CHECK (base_salary > 0);
```

---

## 📈 SCALABILITY CONSIDERATIONS

### Partitioning Strategy
```sql
-- Partition attendances by month
CREATE TABLE attendances_2025_01 PARTITION OF attendances
FOR VALUES FROM ('2025-01-01') TO ('2025-02-01');

-- Partition payrolls by year
CREATE TABLE payrolls_2025 PARTITION OF payrolls
FOR VALUES FROM ('2025') TO ('2026');
```

### Archive Strategy
```sql
-- Archive old attendance data (> 2 years)
-- Archive old payroll data (> 7 years for compliance)
-- Soft delete with deleted_at timestamp
```

---

## 🔄 DATA MIGRATION PLAN

### Phase 1: Master Data
1. Roles and permissions
2. Departments and positions
3. Master data tables

### Phase 2: Employee Data
1. User accounts
2. Employee profiles
3. Historical data

### Phase 3: Transactional Data
1. Attendance records
2. Leave history
3. Payroll history

### Phase 4: Validation
1. Data integrity checks
2. Business rule validation
3. Performance testing

---

## 📊 ERD VISUAL REPRESENTATION

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│    USERS    │────│    ROLES    │    │ EMPLOYEES   │
│             │    │             │    │             │
│ id (PK)     │    │ id (PK)     │    │ id (PK)     │
│ name        │    │ name        │    │ user_id (FK)│
│ email       │    │ description │    │ employee_id │
│ password    │    │             │    │ full_name   │
│ role_id (FK)│    │             │    │ department  │
└─────────────┘    └─────────────┘    │ salary      │
                                      └─────────────┘
                                             │
                   ┌─────────────────────────┼─────────────────────────┐
                   │                         │                         │
            ┌─────────────┐         ┌─────────────┐         ┌─────────────┐
            │ ATTENDANCES │         │LEAVE_REQUESTS│        │  PAYROLLS   │
            │             │         │             │         │             │
            │ id (PK)     │         │ id (PK)     │         │ id (PK)     │
            │employee_id  │         │employee_id  │         │employee_id  │
            │ date        │         │ start_date  │         │ month       │
            │ check_in    │         │ end_date    │         │ year        │
            │ check_out   │         │ status      │         │ net_salary  │
            └─────────────┘         └─────────────┘         └─────────────┘
```

---

## ✅ DATABASE VALIDATION CHECKLIST

### Data Integrity
- [ ] All foreign key constraints defined
- [ ] Primary keys on all tables
- [ ] Unique constraints where needed
- [ ] Check constraints for business rules

### Performance
- [ ] Indexes on frequently queried columns
- [ ] Composite indexes for complex queries
- [ ] Partitioning for large tables
- [ ] Query optimization

### Security
- [ ] Sensitive data encryption
- [ ] Access control at database level
- [ ] Audit trail implementation
- [ ] Backup and recovery procedures

### Compliance
- [ ] Data retention policies
- [ ] GDPR compliance for personal data
- [ ] Audit requirements met
- [ ] Regulatory compliance checks