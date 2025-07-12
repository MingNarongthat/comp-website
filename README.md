# Personal Web Project

A full-stack web application built with Go (Gin) backend, NuxtJS frontend, and PostgreSQL database.

- **personal-web/**  
  - **backend/** _Go Gin backend_  
    - **cmd/** _Main applications of the project_  
      - **api/** _API server entry point_  
        - `main.go`  
    - **internal/** _Private application code_  
      - **config/** _Configuration_  
      - **handlers/** _HTTP handlers_  
      - **middleware/** _HTTP middleware_  
      - **models/** _Database models_  
      - **repository/** _Database operations_  
      - **services/** _Business logic_  
    - **pkg/** _Public libraries that can be used by external applications_  
    - **migrations/** _Database migrations_  
    - **tests/** _Test files_  
    - `.env` _Environment variables_  
    - `go.mod` _Go modules file_  
    - `go.sum` _Go modules checksum_  

  - **frontend/** _NuxtJS frontend_  
    - **assets/** _Uncompiled assets (fonts, images, etc.)_  
    - **components/** _Vue components_  
    - **composables/** _Composable functions_  
    - **layouts/** _Page layouts_  
    - **pages/** _Application views and routes_  
    - **plugins/** _Vue plugins_  
    - **public/** _Static files_  
    - **stores/** _State management (if using Pinia)_  
    - **types/** _TypeScript type definitions_  
    - **utils/** _Utility functions_  
    - `.env` _Environment variables_  
    - `nuxt.config.ts` _Nuxt configuration_  
    - `package.json` _Node dependencies_  

- **docker/** _Docker configuration files_  
  - **backend/** _Backend Dockerfile_  
  - **frontend/** _Frontend Dockerfile_  
  - **postgres/** _PostgreSQL initialization scripts_  

- `.gitignore` _Git ignore file_  
- `docker-compose.yml` _Docker compose configuration_  
- `README.md` _Project documentation_  
- `Makefile` _Build and deployment commands_


## Tech Stack

- **Backend**: Go with Gin framework
- **Frontend**: NuxtJS (Vue 3)
- **Database**: PostgreSQL
- **Development**: Docker & Docker Compose

## Project Structure

```
├── backend/         # Go Gin backend
├── frontend/        # NuxtJS frontend
└── docker/          # Docker configuration files
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- Docker and Docker Compose
- PostgreSQL

### Development Setup

1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd personal-web
   ```

2. Start the development environment:
   ```bash
   docker-compose up -d
   ```

3. Backend development:
   ```bash
   cd backend
   go mod download
   go run cmd/api/main.go
   ```

4. Frontend development:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## License

[Your chosen license]
