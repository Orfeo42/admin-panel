# Admin Panel for Invoice and Order Management

Welcome to the **Admin Panel for Invoice and Order Management** project! This Go-based application is designed to
provide administrators with a comprehensive tool to manage invoices and orders efficiently. With this admin panel, you
can track orders, generate invoices, and manage customer information all from one centralized interface.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)

## Features

- **Order Management**: View, edit, and track customer orders.
- **Invoice Management**: Generate and manage invoices with ease.
- **Customer Management**: Store and update customer information.
- **Dash Board**: home page with charts and totals
- **Responsive Design**: Access from both desktop and mobile devices.

## Requirements

- Go 1.23 or higher
- PostgreSQL database
- [Optional] Docker for containerized deployment
- [Optional] Make to run scripts

## Installation

To get started with the Admin Panel, follow these steps:

1. **Clone the repository:**

    ```bash
    git clone https://github.com/Orfeo42/admin-panel.git
    cd admin-panel
    ```

2. **Set up your Go environment:**

   Make sure you have Go installed and properly set up. You can check this by running:

    ```bash
    go version
    ```

3. **Install dependencies:**

    ```bash
    go mod tidy
    ```

4. **Set up your database:**

   you can define the configuration parameter for your database in the .env file

    ```dotenv
    DB_HOST=localhost
    DB_PORT=5432
    DB_DATABASE=admin_panel_db
    DB_USERNAME=admin_panel_user
    DB_PASSWORD=admin_panel_password
   ```
   then to start the db container just run:

    ```bash
    docker-compose -f docker-compose.yaml up -d postgres
    ```

   or run the make command:

    ```bash
    make db-up
    ```

5. **Run the database creation and data preload (you need a xlsx to fill the data in the db):**

    ```bash
	@go run ./preload/main.go
    ```

   or run the make command:

    ```bash
    make db-init
    ```

6. **Start the application:**

    ```bash
	@go run ./cmd/api/main.go
    ```

7. **Access the admin panel:**

   you can define the port for your application in the .env file

    ```dotenv
    PORT=8080
    ```
   Open your web browser and navigate to `http://localhost:8080` to start using the admin panel.
