CREATE TABLE IF NOT EXISTS hydromon.nodes (
    id text,
    description text,
    configuration text,
    state text,
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
);

CREATE TABLE IF NOT EXISTS hydromon.photos (
    reading_time timeuuid,
    number int,
    picture blob,
    PRIMARY KEY (reading_time, number)
);

CREATE TABLE IF NOT EXISTS hydromon.states (
    node_id text,
    name text,
    photos_per_reading int,
    reading_interval int,
    lower_limit double, -- meters
    upper_limit double, --meters
    previous_state text,
    next_state text,
    PRIMARY KEY (node_id, name)
);