CREATE DATABASE banner_service;

\connect "banner_service";

CREATE TABLE IF NOT EXISTS banners
(
    banner_id  SERIAL PRIMARY KEY,
    feature_id INT,
    content    TEXT,
    is_active   BOOLEAN,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tags
(
    tag_id SERIAL PRIMARY KEY,
    name   text NOT NULL
);

CREATE TABLE IF NOT EXISTS features
(
    feature_id SERIAL PRIMARY KEY,
    name       TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS bt
(
    banner_id INT,
    tag_id    INT,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banners (banner_id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users
(
    user_id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    role INT NOT NULL
);

CREATE INDEX idx_banner_feature ON banners (feature_id);
