# Requirements
* Go: go version go1.18.3 darwin/amd64
* Docker-compose: Docker Compose version v2.6.1
* Docker: Docker version 20.10.17, build 100c701

# Installation
## Sign up
* Method: POST
* URL: /api/auth/register

  ```json
  # request
  {
    "email":"madara_so_cool@email.com",
    "name":"Uchiha Madara",
    "password":"qaswedfr",
    "password_confirm":"qaswedfr",
    "phone":"0902902209",
    "address":"Ho Chi Minh City, Vietnam"
  }

  # expected response
  {
    "data": {
        "user": {
            "id": "62d516e67bdf8e9642d0a23a",
            "name": "Uchiha Madara",
            "email": "madara_so_cool@email.com",
            "phone": "0902902209",
            "address": "Ho Chi Minh City, Vietnam",
            "created_at": "2022-07-18T08:16:38.813Z",
            "updated_at": "2022-07-18T08:16:38.813Z"
        }
    },
    "status": "success"
  }
  ```

## Log In
* Method: POST
* URL: /api/auth/login
  ```json
  # request
  {
    "email":"madara_so_cool@email.com",
    "password":"qaswedfr"
  }

  # expected response
  {
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzQwMjUsImlhdCI6MTY1ODEzMzEyNSwibmJmIjoxNjU4MTMzMTI1LCJzdWIiOiI2MmQ1MTZlNjdiZGY4ZTk2NDJkMGEyM2EifQ.Ret0Wv7n0srxYu56KIUBz9ZWLKWbmEFVsGmA-rUmH_GBe0mEwVC80e-66GUGplbsiZ5fG0G59fnQZWyHlEisxA",
    "status": "success"
  }
  ```

  **NOTE**: switch to `Cookies` tab in PostMan to have a look at **Cookie**

## Get User detail
* Method: GET
* URL: /api/users/me
  ```json
  # expected response
  {
    "data": {
        "user": {
            "id": "62d516e67bdf8e9642d0a23a",
            "name": "Uchiha Madara",
            "email": "madara_so_cool@email.com",
            "phone": "0902902209",
            "address": "Ho Chi Minh City, Vietnam",
            "created_at": "2022-07-18T08:16:38.813Z",
            "updated_at": "2022-07-18T08:16:38.813Z"
        }
    },
    "status": "success"
  }
  ```

## Logout
* Method: GET
* URL: /api/auth/logout
  ```json
  # expected response
  {
    "status": "success"
  }
  ```

## Refresh token
* Method: GET
* URL: /api/auth/refresh
* **Note**: switch to **Cookies** tab in PostMan to have a look at few **new access_token**.