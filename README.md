# Servidor

[![Build Status](https://travis-ci.org/hydro-monitor/web-api.svg?branch=master)](https://travis-ci.org/hydro-monitor/web-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/hydro-monitor/web-api)](https://goreportcard.com/report/github.com/hydro-monitor/web-api)
[![codebeat badge](https://codebeat.co/badges/bec1313f-1ff4-4ea7-bc75-6839183ce20e)](https://codebeat.co/projects/github-com-hydro-monitor-web-api-master)

Repositorio del servidor del sistema de medicion y estudios de rios realizado como parte del trabajo profesional de los alumnos Agustina Barbetta y Manuel Porto.

Dentro del servidor se encuentra:
- Una instancia de la base de datos Cassandra para guardar las tablas de usuarios y mediciones. Para ejecutarla es necesario cumplir los requisitos mínimos, detallados en la documentación de Cassandra.
- Una API
  - Con endpoints para utilización de los nodos de medición. Como pueden ser:
    - Consulta de configuración actual
    - Envío de nueva medición (tiempo, altura de agua, foto)
  - (web) Con endpoints para utilización del panel de administración y vista de mediciones. Como pueden ser:
    - Login de usuario
    - Get de mediciones
    - Post de configuración de medición
- Un módulo que sirva como cliente del cluster de Cassandra, el cual pueda ser invocado desde las APIs para hacer queries a cualquier réplica del cluster de Cassandra.

La replicación del servidor es invisible al usuario, lo que se expone es un load balancer. Tanto los nodos de medición como el panel harán requests contra el load balancer que este último redirigirá hacia algún servidor.

Para ejecutar el servidor se deben realizar los siguientes pasos:

1. Crear una Docker network: `docker network create hydromon-net`.
2. Iniciar un container con Cassandra: `docker run --name hydromon-cassandra-1 --net hydromon-net -d cassandra`. Esto
solo es necesario la primera vez, luego puede simplemente correrse `docker start hydromon-cassandra`.
También puede exponerse el puerto 9042 para poder conectarse con el servidor sin que esté corriendo en Docker.
3. Crear el keyspace en la base de datos:
    1. Conectarse a Cassandra a través de `cqlsh`: `docker run -it --network hydromon-net --rm cassandra cqlsh hydromon-cassandra`.
    2. Crear el keyspace: `CREATE KEYSPACE hydromon WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`
4. Compilar el servidor: `make web-api-linux`.
5. Ejecutar `docker build -t hydromon-server .`.
6. Iniciar el container: `docker run --net hydromon-net -p 8080:8080 -d hydromon-server`.


## Levantar un cluster de Cassandra simple
1. Iniciar un container de Cassandra: `docker run --name hydromon-cassandra-1 --net hydromon-net -d cassandra`.
2. Para agregar nuevos containers ejecutar: 
`docker run --name hydromon-cassandra-2 -d -net hydromon-net -e CASSANDRA_SEEDS=hydromon-cassandra-1 cassandra`
3. Si se desean agregar más nodos al cluster repetir el paso 2 cambiando el nombre de cada nodo.
4. Crear el keyspace en la base de datos: `CREATE KEYSPACE hydromon WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };`

### Reemplazar un nodo fallido del Cluster
1. Ingresar a algún nodo activo y ejecutar: `nodetool status`. Registrar la ip del nodo que falló.
2. Ejecutar `docker rm hydromon-cassandra-x` donde x es el número del nodo que falló.
3. Ejecutar `docker run --name hydromon-cassandra-2 -d -net hydromon-net -e CASSANDRA_SEEDS=hydromon-cassandra-1 cassandra -Dcassandra.replace_address_first_boot=<dead_node_ip>`

## Levantar un cluster de Cassandra con volúmenes persistentes
1. Crear una Docker network: `docker network create hydromon-net`.
2. Crear el o los directorios donde residirá toda la información de los nodos de Cassandra.
3. Iniciar un container de Cassandra: `docker run -v /my/own/datadir:/var/lib/cassandra --name hydromon-cassandra-1 --net hydromon-net -d cassandra`.
4. Para agregar nuevos containers ejecutar: 
`docker run -v /my/own/datadir-2:/var/lib/cassandra --name hydromon-cassandra-2 -d -net hydromon-net -e CASSANDRA_SEEDS=hydromon-cassandra-1 cassandra`
5. Si se desean agregar más nodos al cluster repetir el paso 2 cambiando el nombre de cada nodo.
6. Crear el keyspace en la base de datos: `CREATE KEYSPACE hydromon WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };`

### ¿Cómo revivir un nodo caído con su volumen?
### ¿Cómo revivir un nodo caído sin su volumen?
### ¿Cómo revivir tres nodos caídos con sus volúmenes?
