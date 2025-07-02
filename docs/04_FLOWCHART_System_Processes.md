# 🔄 SYSTEM FLOWCHARTS
## HRIS Process Flow Diagrams

---

## 🎯 OVERVIEW
Dokumentasi visual untuk semua proses bisnis utama dalam sistem HRIS, menggambarkan alur kerja dari input hingga output dengan decision points dan exception handling.

---

## 🔐 1. USER AUTHENTICATION FLOW

```mermaid
flowchart TD
    A[User Access System] --> B{Has Account?}
    B -->|No| C[Register New Account]
    B -->|Yes| D[Enter Credentials]
    
    C --> E[Fill Registration Form]
    E --> F[Submit Registration]
    F --> G{Validation OK?}
    G -->|No| H[Show Error Message]
    G -->|Yes| I[Create User Account]
    I --> J[Assign Default Role]
    J --> K[Send Welcome Email]
    K --> L[Login Page]
    
    D --> M{Credentials Valid?}
    M -->|No| N[Show Error Message]
    M -->|Yes| O[Generate JWT Token]
    O --> P[Set User Session]
    P --> Q{Check User Role}
    
    Q -->|Admin| R[Admin Dashboard]
    Q -->|HR| S[HR Dashboard]
    Q -->|Manager| T[Manager Dashboard]
    Q -->|Employee| U[Employee Dashboard]
    
    H --> L
    N --> L
    
    style A fill:#e1f5fe
    style R fill:#c8e6c9
    style S fill:#c8e6c9
    style T fill:#c8e6c9
    style U fill:#c8e6c9
```

### Process Description:
1. **Entry Point**: User accesses the system
2. **Authentication**: Credential validation
3. **Authorization**: Role-based dashboard routing
4. **Session Management**: JWT token generation and validation

---

## 👥 2. EMPLOYEE ONBOARDING FLOW

```mermaid
flowchart TD
    A[New Hire Decision] --> B[HR Receives Hiring Request]
    B --> C[Create Employee Profile]
    C --> D[Fill Basic Information]
    D --> E[Set Job Details]
    E --> F[Configure Compensation]
    F --> G[Upload Documents]
    G --> H{Data Complete?}
    
    H -->|No| I[Request Missing Info]
    I --> D
    H -->|Yes| J[Generate Employee ID]
    
    J --> K[Create User Account]
    K --> L[Assign Role & Permissions]
    L --> M[Set Manager Assignment]
    M --> N[Configure Work Schedule]
    N --> O[Setup Payroll Details]
    O --> P[Send Welcome Package]
    
    P --> Q[Employee First Login]
    Q --> R[Complete Profile Setup]
    R --> S[Digital Signature]
    S --> T[System Orientation]
    T --> U[Active Employee Status]
    
    style A fill:#fff3e0
    style U fill:#c8e6c9
    style I fill:#ffcdd2
```

### Key Decision Points:
- **Data Completeness**: Ensures all required information is captured
- **Role Assignment**: Determines system access levels
- **Manager Assignment**: Establishes reporting hierarchy

---

## ⏰ 3. DAILY ATTENDANCE FLOW

```mermaid
flowchart TD
    A[Employee Arrives] --> B[Open Attendance System]
    B --> C{Already Checked In?}
    
    C -->|No| D[Click Check-In]
    C -->|Yes| E[Show Current Status]
    
    D --> F[Capture Timestamp]
    F --> G[Verify Location]
    G --> H{Location Valid?}
    
    H -->|No| I[Location Error Message]
    H -->|Yes| J[Record Check-In]
    J --> K[Update Attendance Status]
    K --> L[Show Success Message]
    
    L --> M[Work Period]
    M --> N{Break Time?}
    N -->|Yes| O[Record Break Start]
    O --> P[Break Period]
    P --> Q[Record Break End]
    Q --> M
    N -->|No| R{End of Day?}
    
    R -->|No| M
    R -->|Yes| S[Click Check-Out]
    S --> T[Capture End Timestamp]
    T --> U[Calculate Work Hours]
    U --> V{Overtime?}
    
    V -->|Yes| W[Calculate Overtime]
    V -->|No| X[Standard Hours]
    W --> Y[Record Final Attendance]
    X --> Y
    Y --> Z[Generate Daily Report]
    
    I --> B
    E --> N
    
    style A fill:#e8f5e8
    style Z fill:#c8e6c9
    style I fill:#ffcdd2
```

### Automation Features:
- **Location Verification**: GPS/IP-based validation
- **Overtime Calculation**: Automatic computation
- **Break Time Tracking**: Optional break period recording

---

## 🏖️ 4. LEAVE REQUEST WORKFLOW

```mermaid
flowchart TD
    A[Employee Needs Leave] --> B[Open Leave Request Form]
    B --> C[Select Leave Type]
    C --> D[Choose Dates]
    D --> E[Enter Reason]
    E --> F[Check Leave Balance]
    F --> G{Sufficient Balance?}
    
    G -->|No| H[Insufficient Balance Error]
    G -->|Yes| I[Submit Request]
    
    I --> J[Notify Direct Manager]
    J --> K[Manager Reviews Request]
    K --> L{Manager Decision}
    
    L -->|Reject| M[Send Rejection Notice]
    L -->|Approve| N{Requires HR Approval?}
    
    N -->|No| O[Auto-Approve Request]
    N -->|Yes| P[Forward to HR]
    
    P --> Q[HR Reviews Request]
    Q --> R{HR Decision}
    
    R -->|Reject| S[Send HR Rejection]
    R -->|Approve| T[Final Approval]
    
    O --> U[Update Leave Balance]
    T --> U
    U --> V[Update Calendar]
    V --> W[Notify Employee]
    W --> X[Generate Leave Certificate]
    
    M --> Y[Update Request Status]
    S --> Y
    Y --> Z[Notify Employee]
    
    H --> B
    
    style A fill:#fff3e0
    style X fill:#c8e6c9
    style H fill:#ffcdd2
    style M fill:#ffcdd2
    style S fill:#ffcdd2
```

### Approval Matrix:
- **1-3 days**: Manager approval only
- **4-7 days**: Manager + HR approval
- **8+ days**: Manager + HR + Director approval

---

## 💰 5. MONTHLY PAYROLL PROCESSING

```mermaid
flowchart TD
    A[Month End Trigger] --> B[HR Initiates Payroll]
    B --> C[Collect Attendance Data]
    C --> D[Calculate Work Hours]
    D --> E[Process Overtime]
    E --> F[Apply Leave Deductions]
    F --> G[Add Allowances]
    G --> H[Calculate Deductions]
    H --> I[Compute Gross Salary]
    I --> J[Apply Tax Calculations]
    J --> K[Calculate Net Salary]
    K --> L[Generate Payroll Summary]
    
    L --> M{Data Validation OK?}
    M -->|No| N[Review & Correct Errors]
    N --> C
    M -->|Yes| O[Manager Review]
    
    O --> P{Manager Approval?}
    P -->|No| Q[Request Corrections]
    Q --> N
    P -->|Yes| R[Generate Payslips]
    
    R --> S[Export Bank Transfer File]
    S --> T[Process Payments]
    T --> U{Payment Successful?}
    
    U -->|No| V[Handle Payment Errors]
    V --> T
    U -->|Yes| W[Update Payment Status]
    W --> X[Send Payslips to Employees]
    X --> Y[Generate Reports]
    Y --> Z[Archive Payroll Data]
    
    style A fill:#fff3e0
    style Z fill:#c8e6c9
    style N fill:#ffcdd2
    style V fill:#ffcdd2
```

### Validation Checkpoints:
- **Data Integrity**: Attendance vs payroll data consistency
- **Calculation Accuracy**: Mathematical validation
- **Approval Gates**: Multi-level approval process

---

## 📊 6. KPI EVALUATION PROCESS

```mermaid
flowchart TD
    A[Performance Period Start] --> B[HR Sets KPI Targets]
    B --> C[Manager Reviews Targets]
    C --> D{Targets Approved?}
    
    D -->|No| E[Revise Targets]
    E --> B
    D -->|Yes| F[Communicate to Employee]
    
    F --> G[Employee Acknowledges]
    G --> H[Performance Period]
    H --> I[Regular Check-ins]
    I --> J{Mid-Period Review?}
    
    J -->|Yes| K[Conduct Mid Review]
    K --> L[Adjust if Needed]
    L --> H
    J -->|No| M{Period End?}
    
    M -->|No| H
    M -->|Yes| N[Collect Performance Data]
    N --> O[Manager Input Actuals]
    O --> P[Calculate Achievement %]
    P --> Q[Generate Performance Score]
    Q --> R[Manager Review]
    R --> S{Score Approved?}
    
    S -->|No| T[Discuss & Adjust]
    T --> O
    S -->|Yes| U[Employee Self-Assessment]
    U --> V[Final Performance Meeting]
    V --> W[Document Results]
    W --> X[Development Planning]
    X --> Y[Archive Performance Data]
    
    style A fill:#e8f5e8
    style Y fill:#c8e6c9
    style E fill:#fff3e0
    style T fill:#fff3e0
```

### Performance Metrics:
- **Quantitative KPIs**: Measurable targets (sales, productivity)
- **Qualitative KPIs**: Behavioral assessments
- **Development Goals**: Skill improvement targets

---

## 🔄 7. SYSTEM INTEGRATION FLOW

```mermaid
flowchart TD
    A[External System Request] --> B{Authentication Valid?}
    B -->|No| C[Return Auth Error]
    B -->|Yes| D[Validate Request Format]
    D --> E{Format Valid?}
    
    E -->|No| F[Return Format Error]
    E -->|Yes| G[Process Request]
    G --> H{Data Operation}
    
    H -->|Read| I[Query Database]
    H -->|Write| J[Validate Business Rules]
    H -->|Update| K[Check Permissions]
    H -->|Delete| L[Verify Constraints]
    
    I --> M[Return Data]
    J --> N{Rules Valid?}
    K --> O{Permission OK?}
    L --> P{Safe to Delete?}
    
    N -->|No| Q[Return Business Error]
    N -->|Yes| R[Execute Write]
    O -->|No| S[Return Permission Error]
    O -->|Yes| T[Execute Update]
    P -->|No| U[Return Constraint Error]
    P -->|Yes| V[Execute Delete]
    
    R --> W[Log Transaction]
    T --> W
    V --> W
    W --> X[Return Success Response]
    
    M --> Y[Format Response]
    Y --> Z[Return to Client]
    X --> Z
    
    C --> Z
    F --> Z
    Q --> Z
    S --> Z
    U --> Z
    
    style A fill:#e1f5fe
    style Z fill:#c8e6c9
    style C fill:#ffcdd2
    style F fill:#ffcdd2
    style Q fill:#ffcdd2
    style S fill:#ffcdd2
    style U fill:#ffcdd2
```

### Integration Points:
- **Banking APIs**: Salary transfer processing
- **Email Services**: Notification delivery
- **Time Clock Systems**: Attendance data sync
- **Accounting Systems**: Financial data exchange

---

## 🚨 8. ERROR HANDLING & RECOVERY FLOW

```mermaid
flowchart TD
    A[System Error Occurs] --> B[Capture Error Details]
    B --> C[Log Error Information]
    C --> D{Error Type}
    
    D -->|Validation| E[Return User-Friendly Message]
    D -->|Business Logic| F[Return Business Error]
    D -->|System| G[Return Generic Error]
    D -->|Database| H[Check Connection]
    D -->|Network| I[Retry Operation]
    
    H --> J{Connection OK?}
    J -->|No| K[Attempt Reconnection]
    J -->|Yes| L[Retry Database Operation]
    
    K --> M{Reconnection Success?}
    M -->|No| N[Escalate to Admin]
    M -->|Yes| L
    
    I --> O{Retry Successful?}
    O -->|No| P[Check Retry Count]
    O -->|Yes| Q[Continue Processing]
    
    P --> R{Max Retries?}
    R -->|No| S[Wait & Retry]
    R -->|Yes| T[Fail Operation]
    
    S --> I
    
    E --> U[User Notification]
    F --> U
    G --> U
    T --> U
    N --> V[Admin Notification]
    Q --> W[Success Response]
    L --> X{Operation Success?}
    X -->|Yes| W
    X -->|No| T
    
    style A fill:#ffebee
    style W fill:#c8e6c9
    style N fill:#ff5722
    style V fill:#ff5722
```

### Error Categories:
- **User Errors**: Invalid input, validation failures
- **Business Errors**: Rule violations, insufficient permissions
- **System Errors**: Server issues, database problems
- **Network Errors**: Connectivity issues, timeouts

---

## 📱 9. MOBILE APP SYNC FLOW

```mermaid
flowchart TD
    A[Mobile App Launch] --> B[Check Network Connection]
    B --> C{Online?}
    
    C -->|No| D[Offline Mode]
    C -->|Yes| E[Authenticate User]
    
    D --> F[Load Cached Data]
    F --> G[Enable Offline Features]
    G --> H[Queue Sync Operations]
    H --> I{Network Available?}
    I -->|No| G
    I -->|Yes| J[Sync Queued Data]
    
    E --> K{Auth Success?}
    K -->|No| L[Show Login Screen]
    K -->|Yes| M[Check Data Freshness]
    
    M --> N{Data Stale?}
    N -->|No| O[Load Cached Data]
    N -->|Yes| P[Fetch Fresh Data]
    
    P --> Q[Update Local Cache]
    Q --> R[Merge with Local Changes]
    R --> S{Conflicts?}
    
    S -->|Yes| T[Resolve Conflicts]
    S -->|No| U[Update UI]
    
    T --> V[User Conflict Resolution]
    V --> U
    
    J --> W{Sync Success?}
    W -->|No| X[Retry Later]
    W -->|Yes| Y[Update Sync Status]
    
    O --> U
    U --> Z[App Ready]
    Y --> Z
    X --> H
    
    style A fill:#e8f5e8
    style Z fill:#c8e6c9
    style D fill:#fff3e0
    style L fill:#ffcdd2
```

### Sync Strategy:
- **Real-time**: Critical data (attendance, approvals)
- **Periodic**: Regular data (employee info, schedules)
- **On-demand**: Large data sets (reports, analytics)

---

## 🔒 10. SECURITY AUDIT FLOW

```mermaid
flowchart TD
    A[Security Event Triggered] --> B[Capture Event Details]
    B --> C[Classify Security Level]
    C --> D{Risk Level}
    
    D -->|Low| E[Log Event]
    D -->|Medium| F[Alert Security Team]
    D -->|High| G[Immediate Response]
    D -->|Critical| H[Emergency Protocol]
    
    E --> I[Regular Monitoring]
    F --> J[Investigate Event]
    G --> K[Block Suspicious Activity]
    H --> L[System Lockdown]
    
    J --> M{Threat Confirmed?}
    M -->|No| N[False Positive]
    M -->|Yes| O[Escalate Response]
    
    K --> P[Analyze Attack Pattern]
    L --> Q[Notify Stakeholders]
    
    P --> R[Update Security Rules]
    Q --> S[Incident Response Team]
    
    N --> T[Update Detection Rules]
    O --> U[Implement Countermeasures]
    R --> V[Monitor Effectiveness]
    S --> W[Recovery Planning]
    
    T --> X[Continue Monitoring]
    U --> X
    V --> X
    W --> Y[System Recovery]
    
    I --> Z[Periodic Review]
    X --> Z
    Y --> Z
    
    style A fill:#ffebee
    style H fill:#d32f2f
    style L fill:#d32f2f
    style Z fill:#c8e6c9
```

### Security Monitoring:
- **Login Attempts**: Failed authentication tracking
- **Data Access**: Unusual data access patterns
- **System Changes**: Configuration modifications
- **Network Activity**: Suspicious network traffic

---

## 📋 FLOWCHART LEGEND

### Symbols Used:
- **🟢 Start/End**: Process initiation and completion points
- **🔷 Process**: Standard processing steps
- **🔶 Decision**: Decision points with Yes/No branches
- **🔴 Error**: Error conditions and handling
- **🟡 Warning**: Caution points requiring attention
- **📊 Data**: Data storage or retrieval operations

### Color Coding:
- **Green**: Success states and completion
- **Red**: Error states and failures
- **Yellow**: Warning states and attention points
- **Blue**: Information and processing states
- **Orange**: Decision points and branching

---

## ✅ PROCESS VALIDATION CHECKLIST

### Flow Completeness:
- [ ] All entry points identified
- [ ] All exit points defined
- [ ] Decision branches covered
- [ ] Error handling included

### Business Logic:
- [ ] Business rules implemented
- [ ] Approval workflows defined
- [ ] Validation checkpoints included
- [ ] Exception handling covered

### User Experience:
- [ ] User-friendly error messages
- [ ] Clear process feedback
- [ ] Intuitive navigation flow
- [ ] Mobile-responsive design

### Technical Implementation:
- [ ] Database transactions defined
- [ ] API endpoints mapped
- [ ] Security checkpoints included
- [ ] Performance considerations addressed