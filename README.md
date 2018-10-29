# Shale Tasks API

## Details

This project is deployed on GKE with a MongoDB data layer as a 3 node stateful set. There is a TLS reverse proxy
in place, however, the let's encrypt token isn't quite working correctly. That may or may not be an error in the
proxy's deployment. A postman library has been added to hit the endpoints, but one must turn off the `SSL certificate verification`
setting to make it work in its current state. All repos regarding the project can be found [here](https://github.com/matthewberryhill).
The tasks service is scaled to 3 nodes on the cluster.

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
$ cd $GOPATH/src/github.com/matthewberryhill/shale-tasks-api
# point to local mongo instance
$ export MONGO=dev
# up local mongoDB instance
$ docker-compose -f mongo-compose.yaml up -d
# see mongo running
$ docker ps
# run app
$ go run main.go
```

### CI/CD Pipeline

[Pipeline](https://travis-ci.org/matthewberryhill/shale-tasks-api)

## API Spec

**Use the following URL to hit the endpoints**

`https://matthewberryhill.com`

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
  "date_completed": {unix_timestamp},
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
    "date_completed": {unix_timestamp},
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
  "date_completed": {unix_timestamp},
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

*Once a task in complete, that task cannot be incomplete*

```
{
  "id": {string},
  "task": {string},
  "date_created": {unix_timestamp},
  "date_completed": {unix_timestamp},
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
