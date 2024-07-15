# go-folder-structure

# In this setup:

# Model: Represents the user data and database schema using GORM.

# Service: Handles the business logic and database interactions with GORM.

# Controller: Handles HTTP requests and responses.

# Routes: Defines the API endpoints and injects the database connection into controllers.

# This approach makes the system more modular and easily extensible for future microservices by centralizing the database initialization and using dependency injection for controllers.

## Clone the Repository

# To clone this repository, run:

```sh
git clone https://github.com/pavanmuttange/go-folder-structure.git
cd repo


# Running the Project

Navigate to the project directory: 

 cd go-folder-structure

Download the necessary dependencies:

 go mod tidy

Run the application:

 go run main.go

