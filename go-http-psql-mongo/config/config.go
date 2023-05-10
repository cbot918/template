package config

const (
  // host     = "localhost"
  host = "db"
  port     = 5432
  user     = "postgres"
  password = "12345"
  dbname   = "testdb"

  web_port = ":5455"
)

type Config struct {
  Host string
  Port int32
  User string
  Password string
  Dbname string
  WebPort string
}

func NewConfig() *Config{
  return &Config{
    Host: host,
    Port: port,
    User: user,
    Password: password,
    Dbname: dbname,
    WebPort: web_port,
  }
}
