CREATE SCHEMA IF NOT EXISTS sts_application;

CREATE TABLE sts_application."user" (
	client_name VARCHAR(255) PRIMARY KEY,
	client_id VARCHAR (200) UNIQUE NOT NULL,
	client_secret VARCHAR (255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP 
);