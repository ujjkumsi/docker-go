ALTER KEYSPACE system_auth WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
DROP KEYSPACE IF EXISTS movieapi;
CREATE KEYSPACE movieapi
WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
USE movieapi;
DROP TABLE IF EXISTS movie;
CREATE TABLE movie (
  id uuid PRIMARY KEY,
  name text,
  cover_image text,
  description text
);