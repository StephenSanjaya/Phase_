CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    created_at DATE,
    updated_at DATE,
);

CREATE TABLE photos (
    photo_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    caption VARCHAR(255) NOT NULL,
    photo_url VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    created_at DATE,
    updated_at DATE,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE socialmedias (
    social_media_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    social_media_url VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    created_at DATE,
    updated_at DATE,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    photo_id INT NOT NULL,
    message VARCHAR(255) NOT NULL,
    created_at DATE,
    updated_at DATE,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
    FOREIGN KEY (photo_id) REFERENCES photos(photo_id)
);
