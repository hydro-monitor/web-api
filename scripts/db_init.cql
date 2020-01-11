CREATE KEYSPACE hydromon WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

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
    reading_time timeuuid,
    water_level float,
    PRIMARY KEY (node_id, reading_time)
);

CREATE TABLE IF NOT EXISTS hydromon.photos (
    reading_id uuid,
    photo_id timeuuid,
    photo blob,
    PRIMARY KEY (reading_id, photo_id)
);

CREATE TABLE IF NOT EXISTS hydromon.configurations (
    id text,
    description text,
    states list<text>,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS hydromon.states (
    id text,
    photos_per_reading int,
    ms_between_readings int,
    water_level_limit_for_previous_state float,
    water_level_limit_for_next_state float,
    previous_state text,
    next_state text,
    PRIMARY KEY (id)
);

insert into hydromon.nodes (id, description, configuration, state, manual_reading) values
    ('lujan-1',
     'Nodo instalado debajo de un puente',
     'configuracion-1',
     'normal',
     false);