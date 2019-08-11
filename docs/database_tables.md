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