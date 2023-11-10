CREATE TABLE users
(
    id serial PRIMARY KEY not null,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    status boolean default false
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

CREATE TABLE baskets(
    id serial PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) on delete cascade not null,
    product_id INTEGER REFERENCES products (id) on delete cascade not null
);


-- INSERT INTO brands (title) values ('Nike');
-- INSERT INTO brands (title) values ('Adidas');
-- INSERT INTO brands (title) values ('Gucci');

-- INSERT INTO categories (title) values ('Спортивная обувь');
-- INSERT INTO categories (title) values ('Летняя обувь');
-- INSERT INTO categories (title) values ('Зимняя обувь');

