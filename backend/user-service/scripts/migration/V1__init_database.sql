CREATE SCHEMA IF NOT EXISTS application;

CREATE TABLE application."user" (
   id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
    profile_avatar_url VARCHAR ( 255 ),
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP 
);