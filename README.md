# GoLang User Management System

This project consists of basic CRUD operations for managing users, roles & user-role mappings

## Technologies

```bash
Language: GoLang
DB: MySQL
ORM framework: GORM
HTTP Request handler framework: Gorilla Mux
Package Manager: Glide
```

## Operations
| User Services | Endpoint |   
| :---:   | :-: |
| Create User | POST: /v1/users/ |
| Update User | PUT: /v1/users/{userId} |
| Get All Users | GET: /v1/users/ |
| Get User By Id | GET: /v1/users/{userId} |
| Delete User | DELETE: /v1/users/{userId} |


| Role Service | Endpoint |
| :---:   | :-: |
| Create Role | POST: /v1/roles/ |
| Update Role | PUT: /v1/roles/{roleId} |
| Get All Roles | GET: /v1/roles/ |
| Get Role By Id | GET: /v1/roles/{roleId} |
| Delete Role | DELETE: /v1/roles/{roleId} |

| User-Role Mapping | Endpoint |
| :---:   | :-: |
| Create User Role Mapping | POST: /v1/users/{userId}/roles |
| Get User Roles | PUT: /v1/users/{userId}/roles |

