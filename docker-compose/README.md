# Docker Compose Setup

This Docker Compose configuration allows you to run statuz, Prometheus, and Grafana in a Dockerized environment.

## Prerequisites

Before you can use this Docker Compose setup, ensure you have the following prerequisites installed on your system:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)

## Configuration

To configure the behavior of the statuz, you need to set environment variables in the `docker-compose.yaml` file for the `statuz` service. These environment variables determine how statuz checks and reports the health status:
Refer to the [README](../README.md)  for additional details and examples on how to configure Statuz for your specific use case.

Example configuration:

```yaml
environment:
  TYPE: webpage
  ENDPOINT: https://www.google.com
  INTERVAL: 15s
```

## Usage

1. Navigate to the directory containing the `docker-compose.yml` file.
2. Customize the `docker-compose.yml` file with your desired Statuz configuration (as described above) and adjust Prometheus and Grafana configurations if necessary.
3. Start the services in detached mode (background) using the following command:
   ```bash 
    docker-compose up -d
   ```
4. Access statuz, Prometheus, and Grafana in your web browser:
   - statuz: http://localhost:18080 (or the configured port)
   - Prometheus: http://localhost:19090 (or the configured port)
   - Grafana: http://localhost:13000 (or the configured port)
### Cleanup
When you're finished using the services or need to stop them, follow these steps:
1. Open a terminal and navigate to the directory containing the `docker-compose.yml` file.
2. To stop and remove the running containers, use the following command:
   ```bash
   docker-compose down
   ```

