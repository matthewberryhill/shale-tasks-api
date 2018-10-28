# Shale Tasks API

## API spec

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
