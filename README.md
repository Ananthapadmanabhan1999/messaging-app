This is a sample event-driven, real-time messaging application that follows a
microservices architecture.

ToDO : Add more testcases. Refactor k8s

| Service               | Description                                                                                                          |
| --------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `auth-service`        | Stateless user authentication (e.g., login/signup)                                                                   |
| `websocket-service`   | WebSocket server: receives messages from clients, publishes to Kafka, and relays messages back from Kafka to clients |
| `persistence-service` | Kafka consumer that stores messages in PostgreSQL                                                                    |

Update the .env with your database details and also update kafka details.


How to Deploy (Kubernetes)

------- Set Kubernetes context--------

kubectl config use-context your-cluster

---- Apply namespace ------

kubectl apply -f k8s/namespace.yaml

------- Start Kafka and Zookeeper -------

kubectl apply -f k8s/zookeeper-deployment.yaml

kubectl apply -f k8s/kafka-deployment.yaml

-----------Deploy PostgreSQL-----

kubectl apply -f k8s/postgres-deployment.yaml

-------------Deploy Services-------------

kubectl apply -f k8s/auth-deployment.yaml

kubectl apply -f k8s/websocket-deployment.yaml

kubectl apply -f k8s/persistence-deployment.yaml



>>>>>>>  Testing the App >>>>>>>>


Use curl or Postman to call the Auth API

Use a WebSocket client to connect to ws://websocket-service:8081/ws

Check messages in PostgreSQL using a tool like pgAdmin or CLI



 