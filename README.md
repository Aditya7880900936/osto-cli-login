# 🔐 OSTO CLI Authentication System

A secure, containerized Command Line Authentication System built with **Go**, featuring **user registration**, **secure authentication**, **TOTP-based Two-Factor Authentication (Google Authenticator compatible)**, **session management**, **account lockout**, and **persistent SQLite storage**.

The project follows a layered architecture (Controller → Service → Repository) with a strong focus on clean code, security best practices, and maintainability.

---

## 📌 Features

### Authentication

- User Registration
- User Login
- Secure Password Hashing using bcrypt
- Google Authenticator compatible TOTP Authentication
- Account Lockout after multiple failed login attempts
- Session Management
- Configurable Session Timeout
- Last Login Tracking

### Interactive CLI

- Interactive Command Prompt
- Command History
- Tab Completion
- Helpful Error Messages
- Success Feedback
- Auto Display User Information after Login

### User Commands

**Before Login**

- register
- login
- help
- exit

**After Login**

- whoami
- enable-2fa
- disable-2fa
- logout
- help

### Security Features

- bcrypt Password Hashing
- TOTP (RFC 6238)
- Google Authenticator Compatible
- Session Timeout
- Failed Login Lockout
- Secure Password Input (Hidden Terminal Input)
- Persistent Database

---

## 🏗️ Architecture

The application follows a layered architecture to separate responsibilities and improve maintainability.

```
CLI
 │
 ▼
Controller
 │
 ▼
Service
 │
 ▼
Repository
 │
 ▼
SQLite Database
```

### Layer Responsibilities

| Layer | Responsibility |
|-------|-----------------|
| CLI | Interactive command processing |
| Controller | User input/output handling |
| Service | Business logic |
| Repository | Database operations |
| Database | SQLite persistence |

### 📂 Project Structure

```
osto-cli-login/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── cli/
│   ├── config/
│   ├── controllers/
│   ├── database/
│   ├── models/
│   ├── repository/
│   ├── services/
│   ├── session/
│   └── utils/
│
├── storage/
│
├── Dockerfile
├── docker-compose.yml
├── README.md
├── go.mod
└── go.sum
```

### 🛠️ Tech Stack

| Category | Technology |
|----------|------------|
| Language | Go 1.25 |
| ORM | GORM |
| Database | SQLite |
| Authentication | bcrypt |
| MFA | TOTP (Google Authenticator) |
| CLI | chzyer/readline |
| Containerization | Docker |
| Orchestration | Docker Compose |

### ✨ Highlights

- Production-inspired project structure
- Layered architecture
- Clean separation of concerns
- Dockerized deployment
- Persistent storage
- Secure authentication workflow
- Google Authenticator support
- Interactive terminal experience
- Session management
- Security-first implementation

---

## 🚀 Getting Started

### Prerequisites

Make sure the following tools are installed on your system.

- Go 1.25+
- Docker
- Docker Compose
- Git

### 📥 Clone the Repository

```bash
git clone https://github.com/Aditya7880900936/osto-cli-login.git

cd osto-cli-login
```

### ⚙️ Local Installation

Install dependencies

```bash
go mod tidy
```

Run the application

```bash
go run ./cmd
```

### 🐳 Running with Docker

Build the Docker image

```bash
docker build -t osto-cli .
```

Run the container

```bash
docker run --rm -it osto-cli
```

### 🐳 Running with Docker Compose

Start the application

```bash
docker compose up --build
```

Stop the application

```bash
docker compose down
```

The SQLite database is stored inside the mounted **storage/** directory, allowing user data to persist across container restarts.

---

## 💾 Database

The project uses **SQLite** with **GORM**.

Database file

```
storage/auth.db
```

Features

- Automatic schema migration
- Persistent storage
- Lightweight deployment
- No external database configuration required

---

## 💻 Available Commands

### Before Login

| Command | Description |
|----------|-------------|
| register | Register a new user |
| login | Login with username and password |
| help | Display available commands |
| exit | Exit the application |

### After Login

| Command | Description |
|----------|-------------|
| whoami | Display current logged in user |
| enable-2fa | Enable Google Authenticator based MFA |
| disable-2fa | Disable MFA |
| logout | Logout current session |
| help | Display available commands |

---

## 🔄 Authentication Flow

```
User
   │
   ▼
Register
   │
   ▼
Validate Input
   │
   ▼
Hash Password (bcrypt)
   │
   ▼
Store User
```

---

## 🔐 Login Flow

```
User
   │
   ▼
Username
Password
   │
   ▼
Verify Password
   │
   ▼
Account Locked?
   │
   ├── Yes → Reject Login
   │
   └── No
         │
         ▼
MFA Enabled?
         │
   ├── No
   │      │
   │      ▼
   │  Login Success
   │
   └── Yes
          │
          ▼
     Verify OTP
          │
          ▼
     Login Success
```

---

## 🔑 Two Factor Authentication

The application supports **Time-Based One-Time Password (TOTP)** authentication compatible with:

- Google Authenticator
- Microsoft Authenticator
- Authy
- Any RFC 6238 compatible authenticator application

### Enabling MFA

1. Login
2. Execute `enable-2fa`
3. Copy the generated Secret or OTPAuth URL
4. Add it to Google Authenticator
5. Enter the generated OTP
6. MFA Enabled

Once enabled, every login requires:

- Username
- Password
- Current TOTP Code

---

## ⏳ Session Management

The application maintains an in-memory authenticated session.

Features include

- Session Creation
- Session Timeout
- Session Refresh
- Session Expiration
- Logout

The authenticated user can inspect the active session using

```
whoami
```

which displays

- Username
- Registration Date
- Last Login
- MFA Status
- Session Expiration Time

---

## 🔒 Security Features

Security has been one of the primary focuses while building this project.

### Password Security

- Passwords are never stored in plain text.
- Passwords are securely hashed using **bcrypt**.
- Password verification uses bcrypt hash comparison.

### Account Lockout

To mitigate brute-force attacks, the application temporarily locks an account after multiple consecutive failed login attempts.

Features

- Failed login tracking
- Automatic account lock
- Configurable lock duration
- Automatic reset after successful login

### Multi-Factor Authentication

The project supports optional **Time-Based One-Time Password (TOTP)** authentication.

Features

- Google Authenticator compatible
- OTP verification during login
- Enable/Disable MFA
- Secret key generation
- OTPAuth URL generation

### Session Security

Authenticated users are protected using an in-memory session manager.

Features include

- Session timeout
- Session expiration
- Automatic session refresh
- Logout support

### Secure CLI Input

Passwords are entered using hidden terminal input.

This prevents passwords from being displayed on screen while typing.

---

## 📁 Folder Description

| Folder | Purpose |
|----------|---------|
| cmd | Application entry point |
| internal/cli | Interactive CLI implementation |
| internal/controllers | User interaction layer |
| internal/services | Business logic |
| internal/repository | Database access layer |
| internal/database | Database initialization |
| internal/models | GORM models |
| internal/session | Session management |
| internal/utils | Helper utilities |
| storage | SQLite database |

---

## 📈 Future Improvements

Potential enhancements include

- JWT Authentication
- Role Based Access Control (RBAC)
- Password Reset
- Email Verification
- OAuth Login
- Audit Logging
- User Activity Logs
- Password Strength Meter
- QR Code Rendering in CLI
- Unit & Integration Tests
- CI/CD Pipeline
- GitHub Actions
- PostgreSQL/MySQL Support
- Configuration using Environment Variables

---

## 📸 Example Workflow

```text
Register
      │
      ▼
Login
      │
      ▼
Enable 2FA
      │
      ▼
Scan Secret in Google Authenticator
      │
      ▼
Logout
      │
      ▼
Login
      │
      ▼
Username
Password
OTP
      │
      ▼
Authenticated
```

---

## 🧪 Testing

The application has been tested for the following scenarios.

- User Registration
- Duplicate Registration
- Login Success
- Invalid Password
- Invalid Username
- Password Hashing
- Account Lockout
- Session Creation
- Session Expiration
- Logout
- Google Authenticator TOTP
- Docker Build
- Docker Compose
- SQLite Persistence

---

## 🤝 Contributing

Contributions, suggestions and improvements are welcome.

Feel free to fork the repository and submit a pull request.

---

## 👨‍💻 Author

**Aditya Sanskar Srivastav**

- GitHub: https://github.com/Aditya7880900936
- LinkedIn: https://www.linkedin.com/in/aditya-sanskar-srivastav-a13b08327/

---

## 📄 License

This project was developed as part of the **OSTO Backend Internship Assignment**.

It is intended for educational and evaluation purposes.

---

## 🙏 Acknowledgements

This project uses the following open-source libraries:

- Go
- GORM
- SQLite
- bcrypt
- chzyer/readline
- pquerna/otp
- Docker

---

### ⭐ If you found this project useful, consider giving the repository a star.