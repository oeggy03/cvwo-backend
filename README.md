# cvwo-backend

##Setting up MySQL

A cvwo.sql file has been added to include pre-made communities, posts and comments. You can import it.

1. Install XAMPP.
2. In the XAMPP control panel, start both Apache and MySQL.
3. Click on admin next to MySQL to access PHPMyAdmin
4. You may import the included cvwo.sql file. Or, you may start fresh and create a DB named "cvwoforum".
5. Starting fresh means that you have to sign up for an account yourself, then create communities, comments and posts yourself.

##Running the GO program
### `go mod tidy`
Run this to install any necessary dependencies

### `go run main.go`
Run this to initialize the project. It is also important to run it on port 3001, which has been declared in the .env file.
