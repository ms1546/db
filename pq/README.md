## Usage
1. compile
```
go build -o app ./cmd/server
```
2. start the app server
```
./app

```
3. post(example)
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"john.doe@example.com"}' http://localhost:8080/user

```
