## DevOps

### Description

содержит docker-compose.yaml файл для локальной сборки и k8s-manifests, helm-charts для развертывания приложения в кластере Kubernetes.

### Clusters info
Для проекта нам были выделены следующие ресурсы с уже настроенными кластерами:

__Кластер Kubernetes__:
- CPU:6x
- RAM:18GB
- Накопитель: 220 Gb SSD (storageClass: mts-ssd-fast) 500 IOPS
- Версия Kubernetes: 1.24.6

__Кластер PostgreSQL__:
- CPU:4GB 
- RAM:8GB 
- SSD: 30 GB 
- PSQL v15