# Golang Gin-Gonic App with OAuth2 Authentication

This repository contains a sample Golang web application built using the [Gin-Gonic](https://github.com/gin-gonic/gin) framework with OAuth2 authentication support.
This project is using go modules for package management.

## Prerequisites
- Go 1.20 or higher
- OAuth2 application registered and configuration settings in hand

## Getting Started
1. Clone this repository:
```bash
git clone https://github.com/kjpopov/gingonic-oauth-example/
cd gingonic-oauth-example
```
2. Install dependencies:
```bash
go mod download
```
3. Configure OAuth2:
Update the OAuth2 provider credentials (client ID, client secret, redirect URL) in config.toml. You can also pass client ID and client secret as env variables. See .env.dist for more info.
4. Run the application:
```bash
go run main.go
```
5. Access the app:
Open your web browser and navigate to http://localhost:8080. You should see the landing page of the app.

## Features
User authentication and authorization using OAuth2
Secure session management

## Project Desired Structure
main.go: Entry point of the application.
config/: Configuration files OAuth2 and other settings.
handlers/: HTTP request handlers for different routes.
models/: Data models and database interaction.
templates/: HTML templates for rendering views.
static/: Static files (CSS, JavaScript, etc.).
routes/: Definition of different routes and their handlers.

## Contributing
Feel free to contribute by opening issues or submitting pull requests. Please follow the Contributing Guidelines when submitting code changes.

# License
This project is licensed under the GNU GENERAL PUBLIC LICENSE Version 3