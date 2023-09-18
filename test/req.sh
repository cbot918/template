URL=http://localhost:8083/signup

curl -X POST -H "Content-Type: application/json" -d '{
  "name": "test123",
  "email": "test123@example.com",
  "password": "12345"
}' $URL