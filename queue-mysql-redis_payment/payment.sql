CREATE TABLE payments (
  payment_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  payment_date DATETIME,first_name VARCHAR(50),
  last_name VARCHAR(50),
  payment_mode VARCHAR(255),
  payment_ref_no VARCHAR (255),
  amount DECIMAL(17,4)
) ENGINE = InnoDB;