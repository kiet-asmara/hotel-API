CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_type INT NOT NULL REFERENCES user_types(type_id),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    deposit_amount DECIMAL(10, 2) NOT NULL CHECK(deposit_amount >= 0) DEFAULT 0
);

CREATE TABLE deposits (
    deposit_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id),
    amount  DECIMAL(10, 2) NOT NULL,
    status varchar(10) NOT NULL,
    invoice_id VARCHAR(255) UNIQUE NOT NULL,
    url VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE user_types (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE bookings (
    booking_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id),
    room_id INT NOT NULL REFERENCES rooms(room_id),
    Checkin_date DATE NOT NULL,
    Checkout_date DATE NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    paid BOOL NOT NULL DEFAULT FALSE 
); -- tambah payment status

CREATE TABLE rooms (
    room_id SERIAL PRIMARY KEY,
    room_type_id INT NOT NULL REFERENCES room_types(room_type_id),
    available BOOL NOT NULL DEFAULT TRUE
);

CREATE TABLE room_types (
    room_type_id SERIAL PRIMARY KEY,
    room_name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    price_per_night DECIMAL(10, 2) NOT NULL,
    available_rooms INT NOT NULL CHECK(available_rooms >= 0)
);

CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    booking_id INT NOT NULL REFERENCES bookings(booking_id),
    payment_date DATE,
    payment_method varchar(10) NOT NULL,
    amount  DECIMAL(10, 2) NOT NULL,
    status varchar(10) NOT NULL,
    invoice_id VARCHAR(255) ,
    url VARCHAR(255)  
);

INSERT INTO user_types (type_name) VALUES
    ('Admin'),
    ('Guest');

INSERT INTO users (user_type, email, password, full_name) VALUES -- pass 12345
    (2,'jim@gmail.com', '$2a$10$fTkNWPLWtYXNmWNjhUlbg.Ce9uGH6Rp4C10gPFAIJ9wsdtUtCayO6','jimmy jones'),
    (2,'jane@gmail.com', '$2a$10$fTkNWPLWtYXNmWNjhUlbg.Ce9uGH6Rp4C10gPFAIJ9wsdtUtCayO6','jane doe'),
    (2,'john@gmail.com', '$2a$10$fTkNWPLWtYXNmWNjhUlbg.Ce9uGH6Rp4C10gPFAIJ9wsdtUtCayO6','john doe');

INSERT INTO room_types (room_name, description, price_per_night, available_rooms) VALUES
    ('Suite', 'A luxorious room', 100.55, 3),
    ('Deluxe', 'An above average room', 50.25, 6),
    ('Regular', 'A normal room', 50.00, 10),
    ('Express', 'A cheap, joint room with bunk beds', 15.00, 5);

INSERT INTO rooms (room_type_id, status) VALUES
    (1, TRUE),
    (1, TRUE),
    (1, TRUE),
    (2, TRUE),
    (2, TRUE),
    (2, TRUE),
    (2, TRUE),
    (2, TRUE),
    (2, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (3, TRUE),
    (4, TRUE),
    (4, TRUE),
    (4, TRUE),
    (4, TRUE),
    (4, TRUE);


INSERT INTO bookings (user_id, room_id, Checkin_date, Checkout_date, total_price) VALUES
