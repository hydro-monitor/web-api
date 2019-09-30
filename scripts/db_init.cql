CREATE KEYSPACE hydromonitor WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

CREATE TABLE hydromonitor.nodes (
    id text,
    description text,
    state text,
    PRIMARY KEY (id)
);

CREATE TABLE hydromonitor.readings (
    node_id text,
    reading_time timestamp,
    water_level float,
    photo blob,
    PRIMARY KEY (node_id, reading_time)
);

CREATE TABLE hydromonitor.manual_readings (
    node_id text,
    reading_required boolean,
    PRIMARY KEY (node_id)
);

CREATE TABLE hydromonitor.states (
    name text,
    photos_per_reading int,
    ms_between_readings int,
    water_level_limit_for_previous_state float,
    water_level_limit_for_next_state float,
    previous_state text,
    next_state text,
    PRIMARY KEY (name)
);

insert into hydromonitor.nodes (id, description, state) values ('lujan-1', 'Nodo instalado debajo de un puente', 'normal');