# Go-Ticketing System


## Overview
This is a RESTful API for a ticketing system built using **Golang (Gin Framework)**, **GORM ORM**, and **MySQL**. The system allows users to register, purchase tickets, and manage bookings, while admins can manage events, users, and reports.

## Features
### **User Management**
- Register and login using JWT authentication.
- View and manage purchased tickets.

### **Event Management**
- CRUD operations for events (Admin only).
- Event filtering and pagination.

### **Ticket Management**
- Users can purchase tickets.
- Cancel booked tickets with validation.
- Ticket status updates (Admin & User roles).

### **Reporting (Admin Only)**
- Summary of ticket sales and revenue.
- Event-based ticket reports.

## Tech Stack
- **Language**: Golang
- **Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL
- **Authentication**: JWT

## API Endpoints
### **User Routes**
- `POST /users/register` - Register a new user.
- `POST /users/login` - Login and get JWT token.
- `GET /admin/users/` - Get all users (Admin).
- `GET /admin/users/:id` - Get a specific user (Admin).
- `PUT /admin/users/:id/role` - Update user role (Admin).

### **Event Routes**
- `GET /events/` - Get all events.
- `GET /events/:id` - Get a specific event.
- `POST /events/` - Create a new event (Admin).
- `PUT /events/:id` - Update an event (Admin).
- `DELETE /events/:id` - Delete an event (Admin).

### **Ticket Routes**
- `GET /tickets/` - Get all tickets for a user.
- `GET /tickets/:id` - Get a specific ticket.
- `POST /tickets/` - Purchase a ticket.
- `PATCH /tickets/:id` - Update ticket status.

### **Report Routes (Admin Only)**
- `GET /reports/summary` - Get ticket sales summary.
- `GET /reports/event/:id` - Get event-specific report.

## API Testing Available on Postman
Postman Documentation: https://universal-desert-823258-1.postman.co/workspace/Ticketing~06ba4170-7213-404a-acdb-1587a8246542/collection/26349837-619c5ba7-8ca2-42b5-9795-675b0ff5815c?
