-- in furhter..
--CREATE TABLE bazar_company
--CREATE TABLE bazar_language
--CREATE TABLE bazar_currency

-- yacht, car, rocket, etc
CREATE TABLE bazar_category
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50)
);
-- new, b/u, broken
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

-- dizel, oil
CREATE TABLE bazar_fuel
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- mechanic, automat
CREATE TABLE bazar_drive_unit
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- pered, zad, oba
CREATE TABLE bazar_transmission
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- kabriolet, sedan
CREATE TABLE bazar_body_type
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- red, white
CREATE TABLE bazar_color
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
-- Kz, Oae
CREATE TABLE bazar_country
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Almaty, Dubai
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
    company VARCHAR(150) NOT NULL UNIQUE,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    country_id INTEGER NOT NULL REFERENCES bazar_country(id) ON DELETE CASCADE ON UPDATE CASCADE,
    city_id INTEGER NOT NULL REFERENCES bazar_city(id) ON DELETE CASCADE ON UPDATE CASCADE,
    role_id INTEGER NOT NULL REFERENCES bazar_roles(id) ON DELETE CASCADE ON UPDATE CASCADE
);
--src to images for Car by id
CREATE TABLE bazar_machine_image
(
    id BIGSERIAL PRIMARY KEY,
    path VARCHAR(300) NOT NULL,
    machine_id  INTEGER NOT NULL REFERENCES bazar_machine(machine_id) ON DELETE CASCADE ON UPDATE CASCADE
);
--now, not use
CREATE TABLE bazar_session
(
    id BIGSERIAL PRIMARY KEY,
    access_uuid VARCHAR(100) NOT NULL UNIQUE,
    refresh_uuid VARCHAR(100) NOT NULL UNIQUE,
    expire_date TIMESTAMP,
    user_id INTEGER NOT NULL REFERENCES bazar_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE UNIQUE
);

CREATE TABLE bazar_test
(
    id BIGSERIAL PRIMARY KEY
)

--machine
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
    model_id INTEGER NOT NULL REFERENCES bazar_model(id) ON DELETE CASCADE ON UPDATE CASCADE,
    brand_id INTEGER NOT NULL REFERENCES bazar_brand(id) ON DELETE CASCADE ON UPDATE CASCADE,
    country_id INTEGER NOT NULL REFERENCES bazar_country(id) ON DELETE CASCADE ON UPDATE CASCADE,
    city_id INTEGER NOT NULL REFERENCES bazar_city(id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id INTEGER NOT NULL REFERENCES bazar_category(id) ON DELETE CASCADE ON UPDATE CASCADE,
    state_id INTEGER NOT NULL REFERENCES bazar_state(id) ON DELETE CASCADE ON UPDATE CASCADE, 
    fuel_id INTEGER NOT NULL REFERENCES bazar_fuel(id) ON DELETE CASCADE ON UPDATE CASCADE,
    drive_unit_id INTEGER NOT NULL REFERENCES bazar_drive_unit(id) ON DELETE CASCADE ON UPDATE CASCADE,
    trans_type_id INTEGER NOT NULL REFERENCES bazar_trans(id) ON DELETE CASCADE ON UPDATE CASCADE,
    body_type_id INTEGER NOT NULL REFERENCES bazar_body_type(id) ON DELETE CASCADE ON UPDATE CASCADE,
    color_id INTEGER NOT NULL REFERENCES bazar_color(id) ON DELETE CASCADE ON UPDATE CASCADE,
    creator_id INTEGER NOT NULL REFERENCES bazar_user(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);