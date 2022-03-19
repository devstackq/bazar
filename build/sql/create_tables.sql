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
    name VARCHAR(50) UNIQUE
);
-- models cars
CREATE TABLE bazar_model
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    brand_id INTEGER NOT NULL REFERENCES bazar_brand(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE bazar_currency
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE bazar_fuel
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE bazar_drive_unit
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE bazar_trans
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE bazar_body_type
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE bazar_color
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- country, city
CREATE TABLE bazar_country
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
CREATE TABLE bazar_city
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    country_id INTEGER NOT NULL REFERENCES bazar_country(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- admin, saler, buyer
CREATE TABLE bazar_roles
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- authorization
CREATE TABLE bazar_user
(
    user_id BIGSERIAL PRIMARY KEY,
    email varchar(255) NOT NULL UNIQUE,
    username VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    phone DECIMAL NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    country_id  INTEGER NOT NULL,
    city_id  INTEGER NOT NULL,
    role_id INTEGER NOT NULL REFERENCES bazar_roles(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- is_superuser BOOLEAN NOT NULL,
 -- session_key VARCHAR(50) NOT NULL PRIMARY KEY,

CREATE TABLE bazar_session
(
    id BIGSERIAL PRIMARY KEY,
    access_token VARCHAR(100) NOT NULL UNIQUE,
    refresh_token VARCHAR(100) NOT NULL UNIQUE,
    expire_date TIMESTAMP,
    user_id INTEGER NOT NULL REFERENCES bazar_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

    -- currency_id INTEGER NOT NULL REFERENCES bazar_currency(id) ON DELETE CASCADE ON UPDATE CASCADE,
CREATE TABLE bazar_machine
(
    machine_id BIGSERIAL PRIMARY KEY,
    vin VARCHAR(200) NOT NULL UNIQUE,
    title VARCHAR(200) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    year INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    odometer DECIMAL NOT NULL,
    volume DECIMAL NOT NULL,
    horse_power INTEGER NOT NULL,
    model_id INTEGER NOT NULL,
    brand_id INTEGER NOT NULL,
    country_id INTEGER NOT NULL,
    city_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL REFERENCES bazar_category(id) ON DELETE CASCADE ON UPDATE CASCADE,
    state_id INTEGER NOT NULL REFERENCES bazar_state(id) ON DELETE CASCADE ON UPDATE CASCADE, 
    fuel_id INTEGER NOT NULL REFERENCES bazar_fuel(id) ON DELETE CASCADE ON UPDATE CASCADE,
    drive_unit_id INTEGER NOT NULL REFERENCES bazar_drive_unit(id) ON DELETE CASCADE ON UPDATE CASCADE,
    trans_type_id INTEGER NOT NULL REFERENCES bazar_trans(id) ON DELETE CASCADE ON UPDATE CASCADE,
    body_type_id INTEGER NOT NULL REFERENCES bazar_body_type(id) ON DELETE CASCADE ON UPDATE CASCADE,
    color_id INTEGER NOT NULL REFERENCES bazar_color(id) ON DELETE CASCADE ON UPDATE CASCADE,

    creator_id INTEGER NOT NULL REFERENCES bazar_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);



