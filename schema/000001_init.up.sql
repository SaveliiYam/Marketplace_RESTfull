CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE categories(
    id INT PRIMARY KEY
    title VARCHAR(50) not null
);

CREATE TABLE brands (
    brand_id INT PRIMARY KEY,
    title VARCHAR(50) not null,
    description TEXT
);

CREATE TABLE shoes (
    id INT PRIMARY KEY,
    brand_id INT not null,
    title VARCHAR(50),
    description TEXT,
    price DECIMAL(10,2)
);

CREATE TABLE cloth (
    id INT PRIMARY KEY,
    brand_id INT not null,
    name VARCHAR(50),
    description TEXT,
    price DECIMAL(10,2)
);

CREATE TABLE accessories (
    id INT PRIMARY KEY,
    brand_id INT not null,
    name VARCHAR(50),
    description TEXT,
    price DECIMAL(10,2)
);