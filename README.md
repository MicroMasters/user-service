# User Microservice Project

Welcome to the **user-service** repository! This project is a microservice designed to manage user roles and access permissions within an application. It provides functionality for managing users with different roles including guest, buyer, supplier, and admin. Users are authenticated using JSON Web Tokens (JWT) for secure access to the application's APIs.

## Technologies Used

- **Programming Language**: Go (Golang)
- **Database Management**:
  - MongoDB for storing user data
  - Redis for caching, authentication, and session management

## Features

- **User Roles**: Users can have different roles such as guest, buyer, supplier, or admin.
- **Access Management**: Each user role has specific permissions and access to certain APIs.
- **JWT Authentication**: Secure authentication using JSON Web Tokens for accessing protected APIs.
- **MongoDB Integration**: Data persistence for user management using MongoDB.
- **Redis Caching**: Utilizes Redis for caching user sessions and authentication tokens.

## Usage

To use the **user-service** microservice, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/user-service.git
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up MongoDB and Redis:

   - Ensure MongoDB and Redis are installed and running.
   - Configure the connection details in the project's configuration files.

4. Build and run the application:

   ```bash
   go build
   ./user-service
   ```

5. Access the APIs based on user roles using JWT for authentication.

## API Documentation

The API documentation for the **user-service** microservice can be found [here](#) (https://github.com/MicroMasters/user-service).

## License

This project is licensed under the [MIT License](LICENSE).
