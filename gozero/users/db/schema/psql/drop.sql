ALTER TABLE IF EXISTS "follow" DROP CONSTRAINT IF EXISTS "follow_from_user_fkey";
ALTER TABLE IF EXISTS "follow" DROP CONSTRAINT IF EXISTS "follow_to_user_fkey";
ALTER TABLE IF EXISTS "posts" DROP CONSTRAINT IF EXISTS "posts_posted_by_fkey";
ALTER TABLE IF EXISTS "post_like" DROP CONSTRAINT IF EXISTS "post_like_target_post_fkey";
ALTER TABLE IF EXISTS "post_like" DROP CONSTRAINT IF EXISTS "post_like_liked_user_fkey";
ALTER TABLE IF EXISTS "comments" DROP CONSTRAINT IF EXISTS "comments_posted_by_fkey";
ALTER TABLE IF EXISTS "comments" DROP CONSTRAINT IF EXISTS "comments_target_post_fkey";
ALTER TABLE IF EXISTS "comment_like" DROP CONSTRAINT IF EXISTS "comment_like_target_comment_fkey";
ALTER TABLE IF EXISTS "comment_like" DROP CONSTRAINT IF EXISTS "comment_like_liked_user_fkey";


DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS follow;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS post_like;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS comment_like;