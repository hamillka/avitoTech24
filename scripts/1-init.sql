CREATE DATABASE banner_service;

\connect "banner_service";

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

CREATE TABLE IF NOT EXISTS banners
(
    banner_id  SERIAL PRIMARY KEY,
    feature_id INT,
    FOREIGN KEY (feature_id) REFERENCES features (feature_id) ON DELETE CASCADE,
    content    TEXT,
    is_active   BOOLEAN,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bt
(
    banner_id INT,
    tag_id    INT,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banners (banner_id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id) ON DELETE CASCADE
);

CREATE INDEX idx_banner_feature ON banners (feature_id);

INSERT INTO banner_service.public.tags
VALUES (1000, 'name1000'),
       (1001, 'name1001'),
       (1002, 'name1002'),
       (1003, 'name1003'),
       (1004, 'name1004'),
       (1005, 'name1005'),
       (1006, 'name1006'),
       (1007, 'name1007'),
       (1008, 'name1008'),
       (1009, 'name1009'),
       (1010, 'name1010'),
       (1011, 'name1011');

INSERT INTO banner_service.public.features
VALUES (5000, 'name5000'),
       (5001, 'name5001'),
       (5002, 'name5002'),
       (5003, 'name5003'),
       (5004, 'name5004'),
       (5005, 'name5005'),
       (5006, 'name5006'),
       (5007, 'name5007'),
       (5008, 'name5008'),
       (5009, 'name5009'),
       (5010, 'name5010'),
       (5011, 'name5011');

INSERT INTO banner_service.public.banners
VALUES (600, 5000, '{"text":"some_text","title":"some_title","url":"some_url"}', true),
       (601, 5001, '{"text":"some_text","title":"some_title","url":"some_url"}', false),
       (602, 5002, '{"text":"some_text","title":"some_title","url":"some_url"}', true),
       (603, 5003, '{"text":"some_text","title":"some_title","url":"some_url"}', true),
       (604, 5004, '{"text":"some_text","title":"some_title","url":"some_url"}', true),
       (605, 5005, '{"text":"some_text","title":"some_title","url":"some_url"}', false),
       (606, 5006, '{"text":"some_text","title":"some_title","url":"some_url"}', true);

INSERT INTO banner_service.public.bt
VALUES (600, 1000),
       (601, 1001),
       (602, 1002),
       (602, 1003),
       (603, 1004),
       (604, 1005),
       (604, 1006),
       (605, 1007),
       (606, 1008);