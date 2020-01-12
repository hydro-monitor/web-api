# Servidor

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
2. Iniciar un container con Cassandra: `docker run --name hydromon-cassandra --net hydromon-net -d cassandra`. Esto
solo es necesario la primera vez, luego puede simplemente correrse `docker start hydromon-cassandra`.
3. Crear el keyspace en la base de datos:
    1. Conectarse a Cassandra a traves de `cqlsh`: `docker run -it --network hydromon-net --rm cassandra cqlsh hydromon-cassandra`.
    2. Crear el keyspace: `CREATE KEYSPACE hydromon WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };`
4. Compilar el servidor: `make`.
5. Ejecutar `docker build -t hydromon-server .`.
6. Iniciar el container: `docker run --network hydromon-net -p 8080:8080 hydromon-server`.