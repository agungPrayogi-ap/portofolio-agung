# Portfolio Fiber

A professional and modern portfolio website built using **Go** and the **Fiber** web framework.

## Features
- **High Performance**: Powered by Go Fiber.
- **Server-Side Rendering**: Uses HTML templates for SEO and speed.
- **Modular Architecture**: Separated handlers and routes for easy maintenance.
- **Modern & Responsive Design**: Clean CSS with a focus on UI/UX and mobile-first approach.
- **Micro-animations**: Subtle transitions to enhance user engagement.

## Structure
- `/handlers`: Contains application logic for each page.
- `/routes`: Centralized route definitions.
- `/views`: HTML templates using the Go HTML template engine.
- `/public`: Static assets (CSS, JS, Images).

## Prerequisites
- **Go** (1.20 or later recommended)

## Installation & Running Locally
1. Clone or download the project.
2. Open your terminal in the project directory.
3. Run `go mod tidy` to install dependencies.
4. Run the application:
   ```bash
   go run main.go
   ```
5. Visit `http://localhost:3000` in your browser.

## Deployment Readiness
- **Static Files**: Fiber is pre-configured to serve static assets efficiently.
- **Configuration**: Port can be set via the `PORT` environment variable.
- **Docker**: Easily containerizable for cloud deployment.
- **Performance**: High concurrency support due to the Fiber framework.

---
Designed and developed by **Agung Prayogi**.
