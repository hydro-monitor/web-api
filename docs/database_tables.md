## Tablas

### Nodos
```
CREATE TABLE hydromonitor.nodes (
    id text,
    description text,
    state text,
    PRIMARY KEY (node_id)
)
```

### Mediciones
```
CREATE TABLE hydromonitor.readings (
    node_id text,
    reading_time timestamp,
    water_level float,
    photo blob,
    PRIMARY KEY (node_id, reading_time)
)
```

### Mediciones manuales
```
CREATE TABLE hydromonitor.manual_readings (
    node_id text,
    reading_required boolean,
    PRIMARY KEY (node_id)
)
```

### Estados
```
CREATE TABLE hydromonitor.states (
    name text,
    photos_per_reading int,
    ms_between_readings int,
    water_level_limit_for_previous_state float,
    water_level_limit_for_next_state float,
    previous_state text,
    next_state text,
    PRIMARY KEY (name)
)
```