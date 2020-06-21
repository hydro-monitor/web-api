CREATE TABLE IF NOT EXISTS hydromon.users (
    email text,
    name text,
    last_name text,
    password blob,
    admin boolean,
    PRIMARY KEY (email)
);

CREATE TABLE IF NOT EXISTS hydromon.nodes (
    id text,
    description text,
    manual_reading boolean,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS hydromon.readings (
    node_id text,
    reading_id timeuuid,
    reading_time timestamp,
    water_level double,
    manual_reading boolean,
    PRIMARY KEY (node_id, reading_id)
) WITH CLUSTERING ORDER BY (reading_id DESC);

CREATE TABLE IF NOT EXISTS hydromon.photos (
    reading_time timeuuid,
    number int,
    picture blob,
    PRIMARY KEY (reading_time, number)
);

CREATE TABLE IF NOT EXISTS hydromon.configurations (
    node_id text,
    configuration text,
    PRIMARY KEY (node_id)
);