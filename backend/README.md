# Backend
* Database  : MongoDB
* Framework : Gin

```
├── database
│   └── mongo.go
├── domain
│   ├── jwt.go
│   ├── login.go
│   ├── sign.go
│   └── user.go
├── go.mod
├── go.sum
├── login
│   ├── delivery
│   │   └── http
│   │       └── login.go
│   └── usecase
│       └── login_usecase.go
├── main.go
├── responses
│   ├── response.go
│   └── user.go
├── signup
│   ├── delivery
│   │   └── http
│   │       └── user.go
│   └── usecase
│       └── user.go
└── user
    ├── delivery
    │   └── http
    │       ├── middleware
    │       └── user.go
    ├── repository
    │   └── user.go
    └── usecase
        └── user_usecase.go
```