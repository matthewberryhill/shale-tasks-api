# Shale Tasks API

[Project Management](https://github.com/orgs/matthewberryhill/projects/1)

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
# point to local mongo instance
$ export MONGO=dev
# up local mongoDB instance
$ docker-compose -f mongo-compose.yaml up -d
# see mongo running
$ docker ps
# run app
$ cd $GOPATH/src/github.com/matthewberryhill/shale-tasks-api
$ go run main.go
```

### CI/CD Pipeline

[Pipeline](https://travis-ci.org/matthewberryhill/shale-tasks-api)

## API Spec

### GET/config

**Headers**

```text
Accept: application/json
```

**Response Payload**

```
{
  "name": {string},
  "version": {string},
  "environment": {string},
  "error": {string},
}
```

**Response Codes**

```text
200: Ok
500: Internal Server Error
```

### POST/tasks

**Headers**

```text
Accept: application/json
Content-Type: application/json
```

**Request Payload**

*Task cannot share the same string as the task field*

```
{
  "task": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Payload**

```
{
  "id": {string},
  "task": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Codes**

```text
201: Created
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
    "task": {string},
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
  "task": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Codes**

```text
200: Ok
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

*both parameters are optional*

```
{
  "task": {string},
  "completed": {bool}
}
```

**Response Payload**

*Task cannot share the same string as the task field*

```
{
  "id": {string},
  "task": {string},
  "date_created": {unix_timestamp},
  "completed": {bool}
}
```

**Response Codes**

```text
200: Ok
400: Bad Request
404: Not Found
409: Conflict
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
204: No Content
404: Not Found
500: Internal Server Error
```
