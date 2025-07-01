# 📊 KPI Management System - Business Flow & Features

## 🎯 Overview

KPI (Key Performance Indicator) Management System adalah modul untuk mengelola dan melacak kinerja karyawan berdasarkan target yang telah ditetapkan. Sistem ini mendukung berbagai jenis metrik dengan perhitungan otomatis dan pelaporan komprehensif.

---

## 🔄 Business Flow

### **1. Setup Phase (HR/Admin)**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Create KPI     │───►│  Define KPI     │───►│  Set Employee   │
│  Categories     │    │  Metrics        │    │  Targets        │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

**Steps:**
1. **HR Admin** membuat kategori KPI (Sales, Quality, Productivity, dll.)
2. **HR Admin** mendefinisikan metrik KPI dengan target type dan bobot
3. **HR Manager** menetapkan target untuk setiap karyawan per periode

### **2. Performance Tracking Phase**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Record Actual  │───►│  Auto Calculate │───►│  Generate       │
│  Achievement    │    │  Score & Grade  │    │  Reports        │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

**Steps:**
1. **Employee/Manager** mencatat pencapaian aktual
2. **System** menghitung achievement percentage dan weighted score
3. **System** menghasilkan grade dan laporan kinerja

### **3. Evaluation Phase**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Review         │───►│  Performance    │───►│  Action Plan    │
│  Dashboard      │    │  Discussion     │    │  & Follow-up    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

**Steps:**
1. **Manager** review dashboard kinerja karyawan
2. **Manager & Employee** diskusi hasil evaluasi
3. **Manager** membuat action plan untuk improvement

---

## 🏗️ System Architecture

### **Data Model Structure**
```
KPICategory (1) ──── (N) KPIMetric
                           │
                           │ (1)
                           │
Employee (1) ──── (N) KPITarget ──── (1) KPIMetric
    │                      │
    │ (1)                  │ (1)
    │                      │
    └──── (N) KPIActual ───┘
              │
              │ (1)
              │
         KPIReview (N) ──── (1) Employee
```

### **Core Components**

#### **1. KPI Categories**
- **Purpose**: Mengelompokkan KPI berdasarkan area bisnis
- **Default Categories**: 
  - Sales & Revenue
  - Quality
  - Productivity
  - Customer Service
  - Innovation
  - Leadership
  - Learning & Development

#### **2. KPI Metrics**
- **Purpose**: Template KPI yang bisa digunakan berulang
- **Target Types**:
  - `higher_better`: Semakin tinggi semakin baik (Sales, Revenue)
  - `lower_better`: Semakin rendah semakin baik (Defects, Complaints)
  - `exact`: Target tepat (Budget adherence, Timeline)
- **Weighted Scoring**: Setiap metrik memiliki bobot untuk perhitungan total

#### **3. KPI Targets**
- **Purpose**: Target spesifik untuk karyawan di periode tertentu
- **Period Types**: Monthly, Quarterly, Yearly
- **Range Support**: Min/Max values untuk target yang fleksibel

#### **4. KPI Actuals**
- **Purpose**: Pencapaian aktual yang dicatat
- **Auto Calculation**: Achievement percentage dan weighted score
- **Audit Trail**: Siapa dan kapan data diinput

---

## 📈 Calculation Logic

### **Achievement Percentage**
```go
switch targetType {
case "higher_better":
    achievement = (actual / target) * 100
    // Example: Sales 55M vs target 50M = 110%

case "lower_better":
    achievement = (target / actual) * 100
    // Example: Defects 2 vs target 5 = 250%

case "exact":
    achievement = (1 - |actual - target| / target) * 100
    // Example: Budget 98M vs target 100M = 98%
}
```

### **Weighted Score**
```go
weightedScore = (achievement / 100) * weight
totalScore = sum(allWeightedScores)
finalPercentage = (totalScore / totalWeight) * 100
```

### **Grading System**
```go
func calculateGrade(score float64) string {
    if score >= 90 { return "A" }      // Excellent
    if score >= 80 { return "B" }      // Good
    if score >= 70 { return "C" }      // Satisfactory
    if score >= 60 { return "D" }      // Needs Improvement
    return "E"                         // Poor
}
```

---

## 🚀 Key Features

### **1. Multi-Category Support**
- Flexible categorization untuk berbagai area bisnis
- Easy expansion untuk kategori baru
- Category-based filtering dan reporting

### **2. Flexible Target Types**
- **Higher Better**: Sales, Revenue, Customer Acquisition
- **Lower Better**: Defects, Complaints, Response Time
- **Exact Target**: Budget Adherence, Project Timeline

### **3. Weighted Scoring System**
- Setiap KPI memiliki bobot berbeda
- Perhitungan total score yang akurat
- Prioritas KPI berdasarkan importance

### **4. Automated Calculations**
- Real-time achievement calculation
- Automatic grade assignment
- Weighted score computation

### **5. Comprehensive Reporting**
- Individual employee dashboard
- Company-wide performance reports
- Period-based comparisons
- Performance ranking

### **6. Bulk Operations**
- Mass target setting untuk multiple employees
- Bulk data import/export
- Efficient data management

### **7. Audit Trail**
- Track siapa yang set target
- Record siapa yang input actual
- Timestamp untuk semua perubahan

---

## 📊 Dashboard Features

### **Employee Dashboard**
```json
{
  "employee": {
    "employee_id": "EMP001",
    "name": "John Doe",
    "department": "Sales"
  },
  "period": "2025-07",
  "summary": {
    "total_metrics": 5,
    "completed_kpis": 4,
    "average_score": 92.5,
    "grade": "A"
  },
  "kpi_details": [
    {
      "metric": "Monthly Sales Target",
      "target": 50000000,
      "actual": 55000000,
      "achievement": 110.0,
      "score": 2.2,
      "weight": 2.0
    }
  ]
}
```

### **Company Report**
```json
{
  "period": "2025-07",
  "summary": {
    "total_employees": 25,
    "average_company_score": 78.5,
    "grade_distribution": {
      "A": 5, "B": 8, "C": 7, "D": 3, "E": 2
    }
  },
  "top_performers": [
    {
      "employee_id": "EMP001",
      "name": "John Doe",
      "score": 92.5,
      "grade": "A"
    }
  ]
}
```

---

## 🔒 Access Control

### **HR/Admin (Full Access)**
- Create/Edit KPI categories dan metrics
- Set targets untuk semua karyawan
- View semua reports dan analytics
- Bulk operations

### **Manager (Team Access)**
- Set targets untuk team members
- Record actuals untuk team
- View team performance reports
- Individual employee dashboard

### **Employee (Self Access)**
- View own KPI dashboard
- View own targets dan achievements
- Input self-assessment (if enabled)

---

## 📋 API Usage Examples

### **1. Setup KPI System**
```bash
# Create category
POST /api/kpi/categories
{
  "name": "Sales Performance",
  "description": "Sales and revenue KPIs"
}

# Create metric
POST /api/kpi/metrics
{
  "category_id": 1,
  "name": "Monthly Sales Target",
  "unit": "IDR",
  "target_type": "higher_better",
  "weight": 2.0
}
```

### **2. Set Targets (Bulk)**
```bash
POST /api/kpi/targets/bulk
{
  "targets": [
    {
      "employee_id": "EMP001",
      "metric_id": 1,
      "period": "2025-07",
      "target_value": 50000000
    }
  ]
}
```

### **3. Record Achievement**
```bash
POST /api/kpi/actuals
{
  "employee_id": "EMP001",
  "metric_id": 1,
  "period": "2025-07",
  "actual_value": 55000000,
  "notes": "Exceeded target"
}
```

### **4. Get Dashboard**
```bash
GET /api/kpi/dashboard?employee_id=EMP001&period=2025-07
```

---

## 💡 Business Benefits

### **1. Performance Transparency**
- Clear visibility ke target dan achievement
- Objective performance measurement
- Fair evaluation process

### **2. Goal Alignment**
- Company objectives cascade ke individual targets
- Department-level performance tracking
- Strategic goal achievement

### **3. Continuous Improvement**
- Regular performance monitoring
- Trend analysis dan insights
- Action plan untuk improvement

### **4. Data-Driven Decisions**
- Performance analytics untuk HR decisions
- Promotion dan compensation basis
- Training needs identification

### **5. Employee Engagement**
- Clear expectations dan feedback
- Recognition untuk high performers
- Career development planning

---

## 🎯 Implementation Roadmap

### **Phase 1: Core KPI System ✅**
- KPI categories dan metrics management
- Target setting dan tracking
- Basic reporting dan dashboard

### **Phase 2: Advanced Features 🔄**
- KPI templates untuk quick setup
- Performance trends dan analytics
- Email notifications untuk milestones

### **Phase 3: Integration 🔮**
- Integration dengan payroll untuk bonus calculation
- Performance review workflow
- 360-degree feedback integration

---

**© 2025 OKI. All Rights Reserved.**