## Go Server

## API Usage with cURL

#### Check Server Health

```bash
curl -X GET http://localhost:8080/health
````

---

### Create Users

#### Create User: John Doe

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com"
  }'
```

#### Create User: Jane Smith

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "last_name": "Smith",
    "email": "jane.smith@example.com"
  }'
```

---

### Retrieve Users

#### Get All Users

```bash
curl -X GET http://localhost:8080/api/users
```

#### Get User by ID (e.g., ID = 1)

```bash
curl -X GET http://localhost:8080/api/users/1
```

---

### Update a User

#### Update User with ID 1

```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Updated",
    "email": "john.updated@example.com"
  }'
```

#### Verify Updated User

```bash
curl -X GET http://localhost:8080/api/users/1
```

---

### Delete a User

#### Delete User with ID 2

```bash
curl -X DELETE http://localhost:8080/api/users/2
```

#### Verify After Deletion

```bash
curl -X GET http://localhost:8080/api/users
```

---

### Error Handling

#### Get a Non-Existent User (Expect 404)

```bash
curl -X GET http://localhost:8080/api/users/999
```

#### Create a User with Missing Required Fields (Expect 400)

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "",
    "email": ""
  }'
```



