DROP KEYSPACE IF EXISTS movieapi;
CREATE KEYSPACE movieapi
WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'dc1' : 1 };
USE movieapi;
DROP TABLE IF EXISTS movie;
CREATE TABLE movie (
  id uuid PRIMARY KEY,
  name text,
  cover_image text,
  description text
);