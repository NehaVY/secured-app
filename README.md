# Secure Echo API with Dockerized Deployment and CI/CD Pipeline

This project demonstrates a secure and scalable Go-based REST API with features like input sanitization, token-based authentication, and environment variable management. The application is Dockerized using a minimal base image for lightweight and secure deployment. Additionally, a GitHub Actions workflow is provided for CI/CD integration, focusing on code quality, security, and container image scanning.

---

## Features

### REST API
- **Endpoint**: `/echo`
- **Functionality**: Accepts user input via query parameters and responds with an echo message.
- **Security Measures**:
  - Input validation to prevent injection attacks.
  - Token-based authentication with bearer tokens.
  - Proper error handling with appropriate HTTP status codes.
  - JSON-formatted responses.

### Authentication
- **Middleware**: Authenticates requests using the `Authorization` header.
- **Token**: Hardcoded for demonstration (`Bearer securetoken123`).

### Environment Variable Management
- Securely fetches sensitive data using environment variables (e.g., `SECRET_KEY`).

---

## Project Structure
```plaintext
.
├── main.go                # Application source code (Go code for echo API with security middleware)
├── Dockerfile             # Docker configuration for building and running the app in a container
├── .github                # Directory containing CI/CD configurations
│   └── workflows          # Contains the workflow files for GitHub Actions
│       └── secure-build.yml   # GitHub Actions workflow file for secure build, testing, and deployment
├── README.md              # Documentation about the project and its setup


`````
2. Install dependencies and run the application.
    ```
    go mod tidy
    
    go run main.go
    
    ```

3. Test the API
   ````
   curl -H "Authorization: Bearer securetoken123" "http://localhost:8080/echo?input=HelloWorld"
     
4  Dockerized Deployment
   
   Build and run  the docker image.
   ```
    docker build -t secure-app:latest .
    docker run -p 8080:8080 secure-echo-api:latest

```    

5. Test the API
   ```
   curl -H "Authorization: Bearer securetoken123" "http://localhost:8080/echo?input=HelloWorld"
   
  
  # CI/CD Pipeline

This project includes a GitHub Actions workflow (`.github/workflows/secure-build.yml`) with the following steps:

1. **Static Code Analysis**:  
   Uses `gosec` to scan for security vulnerabilities in Go code.

2. **Dependency Scanning**:  
   Uses `syft` to identify vulnerabilities in dependencies.

3. **Docker Image Build**:  
   Builds the Docker image using BuildKit.

4. **Image Scanning**:  
   Uses `trivy` to detect vulnerabilities in the built Docker image.

5. **Image Signing**:  
   Signs the Docker image using `cosign` for authenticity.

6. **Secret Scanning**:  
   Scans the repository for accidental exposure of secrets using `trufflehog`.

## Environment Variables

Ensure the following environment variables are set in the runtime environment:

- `SECRET_KEY`: A secure secret key required by the application.

## GitHub Actions Workflow

The CI/CD pipeline triggers on pushes to the `main` branch and pull requests. It performs automated checks, scans, and builds to ensure security and quality.

## Security Best Practices

- Use a minimal base image (`golang:1.20-alpine`) to reduce the attack surface.
- Employ static and dynamic code analysis tools for vulnerability detection.
- Sign Docker images to ensure integrity.
- Regularly scan dependencies and container images for known vulnerabilities.
- Avoid hardcoding sensitive values; use environment variables instead.

   
