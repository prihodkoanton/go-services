DROP INDEX IF EXISTS users_email_unique_not_deleted;

ALTER TABLE users
    ADD CONSTRAINT users_email_key UNIQUE (email);