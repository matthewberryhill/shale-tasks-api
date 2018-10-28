# Shale Tasks API

## Development

### Clone Project

```bash
$ go get github.com/matthewberryhill/shale-tasks-api
$ cd $GOPATH/src/github.com/matthewberryhill/shale-tasks-api
```

### Get Dependencies

```bash
$ cd $GOPATH/src/github.com/matthewberryhill/shale-tasks-api
# must have the dep dependency manager
$ dep ensure
```

### Run Locally
 
```bash
$ cd $GOPATH/src/github.com/matthewberryhill/shale-tasks-api
$ go run main.go
```

### Build/Run on Docker
 
```bash
# must have docker installed
$ docker build -t shale-tasks-api .
$ docker run -itp 1323:1323 shale-tasks-api
```

### CI/CD Pipeline

[Pipeline](https://travis-ci.org/matthewberryhill/shale-tasks-api)

## API Spec

### POST/tasks

**Headers**

```text
Accept: application/json
Content-Type: application/json
```

**Request Payload**

```
{
  "title": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Payload**

```
{
  "id": {string},
  "title": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Codes**

```text
204: Created
400: Bad Request
409: Conflict
500: Internal Server Error
```

### GET/tasks

**Headers**

```text
Accept: application/json
```

**Response Payload**

```
[
  {
    "id": {string},
    "title": {string},
    "date_created": {unix_timestamp},
    "completed": {bool}
  }
]
```

**Response Codes**

```text
200: Ok
500: Internal Server Error
```

### GET/tasks/:id

**Headers**

```text
Accept: application/json
```

**URI Paramenters**

```text
id: {string}
```

**Response Payload**

```
{
  "id": {string},
  "title": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Codes**

```text
200: Ok
400: Bad Request
404: Not Found
500: Internal Server Error
```

### PUT/tasks/:id

**Headers**

```text
Accept: application/json
Content-Type: application/json
```

**URI Paramenters**

```text
id: {string}
```

**Request Payload**

```
{
  "title": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Payload**

```
{
  "id": {string},
  "title": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Codes**

```text
200: Ok
400: Bad Request
404: Not Found
500: Internal Server Error
```

### DELETE/tasks/:id

**Headers**

```text
Accept: application/json
```

**URI Paramenters**

```text
id: {string}
```

**Response Payload**

```
{}
```

**Response Codes**

```text
200: Ok
400: Bad Request
404: Not Found
500: Internal Server Error
```
