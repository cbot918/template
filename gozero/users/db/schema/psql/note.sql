create table users (id int NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(255) not null, email varchar(255) not null, password varchar(255) not null)

insert into users (name, email, password) values('yale', 'yale918@gmail.com', '12345');
insert into users (name, email, password) values('node', 'nodev918@gmail.com', '12345');


select * from users;


select * from users where `id` = 1 limit 1

select * from users where name = 'yale';