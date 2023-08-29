URL=http://localhost:8081/signup

curl -X POST -H "Content-Type: application/json" -d '{
  "name": "yale1",
  "email": "yale9181@example.com",
  "password": "12345"
}' $URL