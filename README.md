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
