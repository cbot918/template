CREATE TABLE `users` (
  `id` BINARY(16) PRIMARY KEY,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `pic` varchar(255) NOT NULL
);

CREATE TABLE `follow` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `from_user` BINARY(16) NOT NULL,
  `to_user` BINARY(16) NOT NULL
);

CREATE TABLE `posts` (
  `id` BINARY(16) PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `body` varchar(255) NOT NULL,
  `posted_by` BINARY(16) NOT NULL,
  `photo` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp
);

CREATE TABLE `post_like` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `liked_user` BINARY(16) NOT NULL,
  `target_post` BINARY(16) NOT NULL
);

CREATE TABLE `comments` (
  `id` BINARY(16) PRIMARY KEY,
  `texts` varchar(255) NOT NULL,
  `posted_by` BINARY(16) NOT NULL,
  `target_post` BINARY(16) NOT NULL
);

CREATE TABLE `comment_like` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `liked_user` BINARY(16) NOT NULL,
  `target_comment` BINARY(16) NOT NULL
);

ALTER TABLE `follow` ADD FOREIGN KEY (`from_user`) REFERENCES `users` (`id`);

ALTER TABLE `follow` ADD FOREIGN KEY (`to_user`) REFERENCES `users` (`id`);

ALTER TABLE `posts` ADD FOREIGN KEY (`posted_by`) REFERENCES `users` (`id`);

ALTER TABLE `post_like` ADD FOREIGN KEY (`liked_user`) REFERENCES `users` (`id`);

ALTER TABLE `post_like` ADD FOREIGN KEY (`target_post`) REFERENCES `posts` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`posted_by`) REFERENCES `users` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`target_post`) REFERENCES `posts` (`id`);

ALTER TABLE `comment_like` ADD FOREIGN KEY (`liked_user`) REFERENCES `users` (`id`);

ALTER TABLE `comment_like` ADD FOREIGN KEY (`target_comment`) REFERENCES `comments` (`id`);