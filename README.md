# Student API

This repository contains a **Postman collection** for performing CRUD operations on a Student API. The project was created as part of learning API development and API testing using **Postman**. The backend is built using **Go (Golang)**.

---

## 📝 Description

The **Student API Postman Collection** demonstrates the following HTTP operations on the Student resource:

- **Create** (POST)
- **Read (Single User)** (GET)
- **Read (All Users)** (GET)

These API requests help to practice and showcase how to work with REST APIs using **Postman**, which is widely used for API testing and development.

---

## 🎓 Source / Tutorial Reference

This project is inspired by the tutorial from **Coder's Gyan** on YouTube.

🔗 **Tutorial Link:**  
[REST API Tutorial by Coder's Gyan](https://www.youtube.com/watch?v=OGhQhFKvMiM&list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6&index=31)

---

## ⚙️ Requirements

- **Go (Golang)** installed on your system.
- **Postman** — for API testing ([Download here](https://www.postman.com/downloads/)).
- **TablePlus** — to view and verify database tables ([Download here](https://tableplus.com/)).

---

## 🚀 How to Run & Use

### 1. Clone this repository:

```bash
git clone https://github.com/meet4041/Student-api.git
cd Student-api

➤ To start the server:
go run cmd/students-api/main.go -config config/local.yaml
The server will start at:
http://localhost:8082

➤ Create User:
Method: POST
URL: http://localhost:8082/api/students
Body (JSON):
{
    "name": "John Doe",
    "age": 21,
    "grade": "A"
}

➤ Get User by ID:
Method: GET
URL: http://localhost:8082/api/students/1

➤ Get All Users:
Method: GET
URL: http://localhost:8082/api/students
