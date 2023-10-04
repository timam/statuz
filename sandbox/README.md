# Sandbox Setup

---

This sandbox configuration allows you to run statuz microservices with Prometheus, and Grafana in a Dockerized environment.
It also gives statuz microservices access to necessary kubernetes api with minikube. 

## Prerequisites

Before you can use this Docker Compose setup, ensure you have the following prerequisites installed on your system:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

### Versions
We have tested this sandbox with following tools and version, if you face any issue feel free to raise an issue. 

| name           | version             |
|----------------|---------------------|
| go             | `go1.20.6`          |
| docker         | `24.0.6`            |
| docker-compose | `v2.22.0-desktop.2` |
| minikube       | `v1.31.2`           |
| kubernetes     | `v1.27.4`           |
| kubectl        | `v1.28.2`           |



## How to setup sandbox environment?
As this sandbox is highly dependent on docker and minikube, please make sure docker and minikube is up and running.

We have a script prepared to get started. simply run this command to create/update necessary configs.
```bash
$ bash setup.sh
```

## How to start sandbox environment?
Once `setup.sh` updates/sets necessary configs, you can run following command to start statuz sandbox environment.
```bash
$ bash start.sh
```
- Kubernetes API, Prometheus, and Grafana will be available on:
   - Kubernetes API: http://localhost:18001
   - Prometheus: http://localhost:19090
   - Grafana: http://localhost:13000
- statuz microservices will be available on:
   - watcher: http://localhost:18080
   - genie: http://localhost:18081

## How to stop sandbox environment?
We have a bash script prepared to stop statuz sandbox environment. Simply run
```bash
$ bash stop.sh
```

---
## watcher configuration 
*NB: this manual config update will be deprecated soon*

To configure the behavior of the watcher, you need to set environment variables in the `docker-compose.yaml`. These 
environment variables determine how watcher checks and reports the health status.
Refer to the [README](../README.md)  for additional details and examples. Example configuration:

```yaml
environment:
  TYPE: webpage
  ENDPOINT: https://www.google.com
  INTERVAL: 15s
```

