## Tablas

### Nodos
```
CREATE TABLE hydromonitor.nodes (
    node_id text,
    description text,
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
    node_id text,
    name text,
    photos_per_reading int,
    ms_between_readings int,
    water_level_limit_for_previous_state float,
    water_level_limit_for_next_state float,
    previous_state text,
    next_state text,
    PRIMARY KEY (node_id, name)
)
```

### Estados configuraciones
```
CREATE TABLE hydromonitor.configurations_status (
    node_id text,
    name text,
    active bool,
    PRIMARY KEY (node_id, name)
)
```