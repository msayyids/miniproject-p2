CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    deposit_amount INT DEFAULT 0
);


CREATE TABLE IF NOT EXISTS rooms (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	type VARCHAR(255),
	price INT,
	availibility BOOLEAN
);


CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    room_id INTEGER REFERENCES rooms(id),
    total_day INT,
    total_price INT,
    status VARCHAR(255) DEFAULT 'booked'
);


INSERT INTO rooms (name, type, price, availibility)
VALUES ('Single Room 01', 'single room', 100, true),
		('Single Room 02', 'single room', 100, true),
		('Single Room 03', 'single room', 100, true),
		('Single Room 04', 'single room', 100, true),
		('Single Room 05', 'single room', 100, true);


INSERT INTO rooms (name, type, price, availibility)
VALUES ('Twin Room 01', 'twin room', 150, true),
		('Twin Room 02', 'twin room', 150, true),
		('Twin Room 03', 'twin room', 150, true),
		('Twin Room 04', 'twin room', 150, true),
		('Twin Room 05', 'twin room', 150, true);


INSERT INTO rooms (name, type, price, availibility)
VALUES ('Double Room 01', 'double room', 200, true),
('Double Room 02', 'double room', 200, true),
('Double Room 03', 'double room', 200, true),
('Double Room 04', 'double room', 200, true),
('Double Room 05', 'double room', 200, true);
		

INSERT INTO rooms (name, type, price, availibility)
VALUES ('Family Room 01', 'family room', 300, true),
('Family Room 02', 'family room', 300, true),
('Family Room 03', 'family room', 300, true),
('Family Room 04', 'family room', 300, true),

('Family Room 05', 'family room', 300, true);




