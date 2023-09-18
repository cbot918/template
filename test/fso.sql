CREATE TABLE users(
  id        int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email     varchar(255) NOT NULL,
  name      varchar(255) NOT NULL,
  password  varchar(255) NOT NULL
);