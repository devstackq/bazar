-- authorization
CREATE TABLE bazar_user
(
    id BIGSERIAL PRIMARY KEY,
    email varchar(255) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    role_id INTEGER NOT NULL REFERENCES bazar_roles(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- is_superuser BOOLEAN NOT NULL,

-- admin, saler, buyer
CREATE TABLE bazar_roles
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE bazar_car
(
    id BIGSERIAL PRIMARY KEY,
    vin VARCHAR(200) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    year INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    country_id INTEGER NOT NULL REFERENCES bazar_country(id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id INTEGER NOT NULL REFERENCES bazar_category(id) ON DELETE CASCADE ON UPDATE CASCADE,
    state_id INTEGER NOT NULL REFERENCES bazar_state(id) ON DELETE CASCADE ON UPDATE CASCADE, 
    brand_id INTEGER NOT NULL REFERENCES bazar_brand(id) ON DELETE CASCADE ON UPDATE CASCADE,
    model_id INTEGER NOT NULL REFERENCES bazar_model(id) ON DELETE CASCADE ON UPDATE CASCADE,
    saler_id INTEGER NOT NULL REFERENCES bazar_user(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- country, city
CREATE TABLE bazar_country
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    city_id INTEGER NOT NULL REFERENCES bazar_city(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE bazar_city
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- yacht, car, rocket, etc
CREATE TABLE bazar_category
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50)
);
-- new, b/u, razbit
CREATE TABLE bazar_state
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50)
);
-- brands car
CREATE TABLE bazar_brand
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50)
);
-- models cars
CREATE TABLE bazar_model
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE bazar_session
(
    session_key VARCHAR(40) NOT NULL PRIMARY KEY,
    "session_data" text NOT NULL,
    "expire_date" TIMESTAMP NOT NULL
);





