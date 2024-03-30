# User Microservice Project

Welcome to the User Microservice project! This project is built using the Go programming language and is designed to manage user roles and access through APIs. It utilizes JWT token authentication for validating users and MongoDB for database management, along with Redis for caching authentication and session management.

## Roles

The User Microservice project supports the following user roles:

1. **Guest**: Users who are not logged in or authenticated.
2. **Buyer**: Users who have registered as buyers and have access to specific buyer functionalities.
3. **Supplier**: Users who have registered as suppliers and have access to supplier-specific functionalities.
4. **Admin**: Users with administrative privileges, granting them access to manage users, roles, and other administrative tasks.

## API Documentation

For detailed information on the APIs provided by the User Microservice, please refer to the [API documentation](https://github.com/MicroMasters/user-service).

## Getting Started

To get started with using the User Microservice project, follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/MicroMasters/user-service.git
   ```

2. Install dependencies:

   ```
   cd user-service
   go mod download
   ```

3. Set up MongoDB and Redis:

   Ensure you have MongoDB and Redis installed and running on your system. Update the configuration file (`config.yaml`) with your MongoDB and Redis connection details.

4. Build and run the project:

   ```
   go build
   ./user-service
   ```

5. Access the APIs:

   Once the service is up and running, you can access the APIs using the provided endpoints as documented in the [API documentation](https://github.com/MicroMasters/user-service).

## Contributing

We welcome contributions to the User Microservice project! Feel free to fork the repository, make changes, and submit pull requests.

## Issues

If you encounter any issues or have suggestions for improvements, please open an issue on the [issue tracker](https://github.com/MicroMasters/user-service/issues).

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it as per the terms of the license.

---
Developed by [MicroMasters](https://github.com/MicroMasters)

## 🌱 Contributors </br>

<a href="https://github.com/MicroMasters/user-service/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=MicroMasters/user-service" />
</a>
