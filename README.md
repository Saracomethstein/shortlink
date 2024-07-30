
# Shortlink

Shortlink is a web application for creating and managing short URLs. It is built with Go, using the Echo web framework and Gorilla for sessions and cookie handling.

## Features

- User authentication (login and registration)
- URL shortening
- Redirects to original URLs

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
│   │   │   └── handlers.go
│   │   ├── repositories/
│   │   │   ├── user_repository.go
│   │   │   ├── link_repository.go
│   │   │   └── repositories.go
│   │   ├── services/
│   │   │   ├── auth_service.go
│   │   │   ├── link_service.go
│   │   │   └── services.go
│   │   ├── config/
│   │   │   └── config.go
│   │   └── models/
│   │       ├── user.go
│   │       ├── link.go
│   │       └── models.go
├── website/
│   ├── static/
│   └── templates/
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

## Setup and Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Saracomethstein/shortlink.git
    cd shortlink
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

   Modify `.env` to match your database and other configuration settings.

3. Run the application:
    ```sh
    make run
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
