CREATE TABLE users
(
    id serial PRIMARY KEY not null,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE categories(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) not null
);

CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) not null,
    description TEXT
);

CREATE TABLE products(
    id serial PRIMARY KEY,
    title VARCHAR(50) not null,
    description TEXT,
    price DECIMAL(10, 2),
    brand_id INTEGER REFERENCES brands (id) on delete cascade not null,
    categories_id INTEGER REFERENCES categories (id) on delete cascade not null
);