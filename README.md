
# Shortlink

Shortlink is a web application for creating and managing short URLs. It is built with Go, using the Echo web framework and Gorilla for sessions and cookie handling.

## Features

- User authentication (login and registration)
- URL shortening
- Redirects to original URLs

## Requirements
 - Docker
 - Web browser

## Project Structure 
```
├── cmd/
│   └── shortlink/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   ├── link.go
│   │   │   ├── profile.go
│   │   │   └── handlers.go
│   │   ├── repositories/
│   │   │   ├── user_repository.go
│   │   │   ├── link_repository.go
│   │   │   ├── profile_repository.go
│   │   │   └── repositories.go
│   │   ├── services/
│   │   │   ├── auth_service.go
│   │   │   ├── link_service.go
│   │   │   ├── profile_service.go
│   │   │   └── services.go
│   │   └── models/
│   │       └── models.go
├── scripts/
│   └── create_table.sql
├── website/
│   └── static/
│       ├── auth/
│       │   ├── index.html
│       │   └── auth.js
│       ├── registration/
│       │   ├── index.html
│       │   └── regist.js
│       ├── profile/
│       │   ├── index.html
│       │   └── profile.js
│       ├── shorten/
│       │   ├── index.html
│       │   └── shorten.js
│       ├── output/
│       │   ├── index.html
│       │   └── output.js
│       └── styles/
│           └── styles.css
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── Dockerfile.nginx
├── nginx.conf
├── README.md
├── go.mod
└── go.sum
```

## Setup and Installation
1. Clone the repository:
    ```sh
    git clone git@github.com:Saracomethstein/shortlink.git
    cd shortlink
    ```

2. Pull docker images:
    ```sh
    make docker-pull
    ``` 

3. Build and run the application:
    ```sh
    make docker-up
    ``` 
   
4. Open in browser:
    ```sh
    http://localhost/auth/
    ```

## Usage

- **Authentication**
  - Register a new user.
  - Login with existing credentials.
  
- **URL Shortening**
  - Create a short URL by providing the original URL.
  - Redirect to the original URL by accessing the short URL.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For any questions or inquiries, please contact [Saracomethstein](https://github.com/Saracomethstein).
