# TICTAC

## Installation

run `go get drew138/tictac` on your local machine.

## Environment Variables

This project implements environment variables to handle sensitive information and various encryption and hashing algorithms.

### Recommended .env file

Users of this repository must be wary that a .env file is enforced, and its structure must contain at least the following variables:

```
DB_USERNAME=yourDatabaseUsername
DB_PASSWORD=yourDatabasePassword
DB_NAME=yourDatabaseName
DB_HOST=yourDatabaseHost
DB_PORT=yourDatabasePort
JWT_SECRET_KEY=randomJWTSecretKey
JWT_REFRESH_SECRET_KEY=anotherRandomJWTSecretKey
```
