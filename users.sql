CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    profile_picture_url TEXT, -- URL for uploaded profile picture
    bio TEXT, -- Optional biography or description
    phone_number VARCHAR(20), -- Optional phone number
    role VARCHAR(20) DEFAULT 'user', -- 'user', 'seller', 'admin'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
