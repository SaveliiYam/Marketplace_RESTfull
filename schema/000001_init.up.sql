CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY not null,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    status boolean default false
);

CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) not null
);

CREATE TABLE IF NOT EXISTS brands (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) not null,
    description TEXT
);

CREATE TABLE IF NOT EXISTS products(
    id serial PRIMARY KEY,
    title VARCHAR(50) not null,
    description TEXT,
    price DECIMAL(10, 2),
    brand_id INTEGER REFERENCES brands (id) on delete cascade not null,
    categories_id INTEGER REFERENCES categories (id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS baskets(
    id serial PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) on delete cascade not null,
    product_id INTEGER REFERENCES products (id) on delete cascade not null
);

CREATE INDEX user_all ON users(id, name, username, password_hash, status);
CREATE INDEX user_id_status ON users(id, status);
CREATE INDEX user_status ON users(status);
CREATE INDEX category_id_title ON categories (id, title);
CREATE INDEX category_id ON categories(id);
CREATE INDEX brands_id_title_description ON brands (id, title, description);
CREATE INDEX brands_id ON brands(id);


