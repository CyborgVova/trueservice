CREATE SCHEMA service
	CREATE TABLE IF NOT EXISTS service.users (
    	id bigint PRIMARY KEY,
    	name varchar(30) NOT NULL,
    	role varchar(100) NOT NULL
);