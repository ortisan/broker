CREATE SCHEMA IF NOT EXISTS application;

CREATE TABLE application."user" (
	id VARCHAR(36) PRIMARY KEY,
	name VARCHAR (255) NOT NULL,
	username VARCHAR (50) UNIQUE NOT NULL,
	password VARCHAR (255) NOT NULL,
	email VARCHAR (255) UNIQUE NOT NULL,
	federation_id VARCHAR (255) UNIQUE,
	profile_avatar_url VARCHAR (255),
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP 
);