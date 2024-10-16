# Weather App Backend

## Overview

This Go-based GraphQL API serves as a wrapper around the OpenMeteo service to fetch weather data. It was developed using the repository pattern and integrates GraphQL using `gqlgen`.

## Project Setup

### Prerequisites
- Go 1.19 or later
- Docker (for containerized deployment)

### Running the Project Locally

1. Clone the repository and navigate to the project directory.

   ```bash
   git clone <your-repo-url>
   cd <project-folder>
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Run the server for development:

   ```bash
   go run main.go
   ```

4. The server will be available on [http://localhost:8080/](http://localhost:8080/), and you can access the GraphQL playground at this URL.

## Project Structure

- **`graph/`**: Contains all GraphQL-related files, such as schema and resolvers.
- **`repo/`**: Houses the repository that interacts with external APIs (OpenMeteo in this case).
- **`response/`**: Contains types for the responses from the OpenMeteo service.
- **`service/`**: Contains the service logic that connects the repository and GraphQL resolvers.
- **`gqlgen.yml`**: Configuration for gqlgen, which generates the GraphQL schema and resolver stubs.
- **`Dockerfile`**: Docker configuration for containerizing the application.

## Configuration

In a typical production application, configuration should be global and the application should be configurable via environment variables. This would allow for different settings across environments (e.g., development, staging, production). Ideally, this would be done through a separate package dedicated to handling configurations.

However, for simplicity, this project does not use a dedicated configuration package, and it runs directly on port 8080 without additional configuration.

## Frontend Integration

The frontend for this project is located in the following repository:

- **Frontend Repository**: [https://github.com/Am1rArsalan/fe-kg](https://github.com/Am1rArsalan/fe-kg)

This repository contains the React application that interacts with the backend's GraphQL API to display weather data.

## Deployment

To deploy the project, use the infrastructure setup provided in the following repository:

- **Infra Repository**: [https://github.com/Am1rArsalan/kg-infra](https://github.com/Am1rArsalan/kg-infra)

This repository contains Docker Compose configurations to manage both the backend and frontend services, allowing for easy deployment and orchestration.

## Development Approach

This backend follows the repository pattern, which separates the data access logic from the service layer, making the codebase modular and easier to maintain. Due to time constraints, tests have not been added to this project, but the code is structured to facilitate future testing.

## GraphQL Integration

The API uses GraphQL, with `gqlgen` as the tool for generating the schema and resolvers. The GraphQL schema is defined in `schema.graphqls`, and the resolvers can be found in the `graph/resolvers.go` file. For querying weather data, the OpenMeteo service is accessed through the repository layer.
