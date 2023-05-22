# DevOps и Infrastructure as Code для классической трехзвенной архитектуры

Код был продублирован с GitLab в целях демонстрации. Все текущие папки проекта были отдельными группами/репозиториями.

### Purpose

Получить работающую связку микросервисов с масштабируемой и понятной архитектурой и возможностью автоматического развертывания в кластере Kubernetes.

### Project structure

`Core` - содержит код всех микросервисов, с написанными к ним Dockerfile и .gitlab-ci.yaml файлами. Предполагалось, что образы микросервисов загружаются и затем берутся из GitLab Container Regustry.

`DevOps` - содержит docker-compose.yaml файл для локальной сборки и k8s-manifests, helm-charts для развертывания приложения в кластере Kubernetes.

***

### Usage

1. Поскольку приложение развернуто в кластере Kubernetes, им можно легко воспользоваться перейдя по ссылке: http://sakurablog.ru/

2. Для **локального использования** необходимо:

* Склонировать проект с помощью команды: 
    
```
https://github.com/Morozhkaa/Project-DevOps.git
```
* Перейти в директорию DevOps/docker:
```
cd Project-DevOps/DevOps/docker
```

* Запустить проект:

```
docker-compose -f ./docker-compose.yml up -d --build
```

***

### Architecture
<p align="center"><img src="https://github.com/Morozhkaa/Project-DevOps/blob/main/images/container.jpg" width="600"></p>

***

### Demonstration

Выглядит приложение следующим образом:

<p align="center">
<img src="https://github.com/Morozhkaa/Project-DevOps/blob/main/images/login.jpg" width="600">

<img src="https://github.com/Morozhkaa/Project-DevOps/blob/main/images/overview.gif" width="600">
</p>


Более подробный видео-обзор по ___[ссылке](https://drive.google.com/file/d/1CTA_u72i51S_jvFc6jvuYvYN2yVwY7uR/view?usp=share_link)___