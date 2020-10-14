# TICTAC Website API

## Installation

run `go get drew138/tictac` on your local machine.

## Environment Variables

This project implements environment variables to handle sensitive information and various encryption and hashing algorithms.

### Recommended .env file

Although the use of a .env file is not enforced, the following environment variables structure is required:

```
DB_USERNAME=yourDatabaseUsername
DB_PASSWORD=yourDatabasePassword
DB_NAME=yourDatabaseName
DB_HOST=yourDatabaseHost
DB_PORT=yourDatabasePort
JWT_SECRET_KEY=randomJWTSecretKey
JWT_REFRESH_SECRET_KEY=anotherRandomJWTSecretKey
```
