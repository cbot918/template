ALTER TABLE `follow` DROP FOREIGN KEY `follow_ibfk_1`;
ALTER TABLE `follow` DROP FOREIGN KEY `follow_ibfk_2`;
ALTER TABLE `posts` DROP FOREIGN KEY `posts_ibfk_1`;
ALTER TABLE `post_like` DROP FOREIGN KEY `post_like_ibfk_1`;
ALTER TABLE `post_like` DROP FOREIGN KEY `post_like_ibfk_2`;
ALTER TABLE `comments` DROP FOREIGN KEY `comments_ibfk_1`;
ALTER TABLE `comments` DROP FOREIGN KEY `comments_ibfk_2`;
ALTER TABLE `comment_like` DROP FOREIGN KEY `comment_like_ibfk_1`;
ALTER TABLE `comment_like` DROP FOREIGN KEY `comment_like_ibfk_2`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `follow`;
DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `post_like`;
DROP TABLE IF EXISTS `comments`;
DROP TABLE IF EXISTS `comment_like`;