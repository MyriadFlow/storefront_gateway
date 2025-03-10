CREATE TABLE user (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    wallet_address VARCHAR(225)
    interest VARCHAR(255),
    location VARCHAR(255)
);

CREATE TABLE friend_requests (
    id SERIAL PRIMARY KEY,
    sender_id INT,
    receiver_id INT,
    status VARCHAR(50) DEFAULT 'pending', -- pending, accepted, rejected
    FOREIGN KEY (sender_id) REFERENCES people(id),
    FOREIGN KEY (receiver_id) REFERENCES people(id)
);

CREATE TABLE interests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INT,
    interest_id INT,
    location_id INT,
    FOREIGN KEY (user_id) REFERENCES people(id),
    FOREIGN KEY (interest_id) REFERENCES interests(id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    categorie VARCHAR(225),
    name VARCHAR(255)
)