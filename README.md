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

## Websocket

The implemented websocket supports the following messages

### Sending Private Message

Sending message:
```
{
  "action": "privateMessage",
  "recipientID": "oihdfq89ewdfhsdfu80qu8943",
  "body": "Hi, how are you?"
}
```

Received message:
```
{
  "action": "privateMessage",
  "recipientID": "a0sg8enfdls0d72hfla0adimg",
  "body": "Hi, how are you?"
}
```

### Sending Game Actions

Sending message:
```
{
  "action": "gameAction",
  "recipientID": "oihdfq89ewdfhsdfu80qu8943",
  "body": "Hi, how are you?"
}
```

Received message:
```
{
  "action": "gameAction",
  "recipientID": "a0sg8enfdls0d72hfla0adimg",
  "body": "Hi, how are you?"
}
```
