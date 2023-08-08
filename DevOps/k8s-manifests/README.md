## K8s-manifests

Для каждого микросервиса я создала необходимые ему виды манифестов:

- __Pods__ - наименьшие развертываемые вычислительные единицы, которые вы можете создавать и которыми можете управлять в Kubernetes. По сути, это обертка над группой из одного или нескольких контейнеров с общим хранилищем и сетевыми ресурсами, а также спецификацией того, как запускать контейнеры.
- __Service__ - абстракция, помогающая предоставлять доступ к группам Pod по сети.
- __ConfigMap__ - объект API, используемый для хранения неконфиденциальных данных в парах ключ-значение. Поды могут использовать ConfigMaps как переменные среды, аргументы командной строки или как файлы конфигурации в томе.
- __Secret__ - объект, похожий на ConfigMaps, но содержащий конфиденциальные данные, таких как пароль, токен или ключ.

Также, требовалось создать __Ingress__ - объект API, предоставляющий правила маршрутизации для управления доступом внешних пользователей к службам в кластере Kubernetes, обычно через HTTPS/HTTP.

Помимо этого, чтобы получить образы из приватного реестра (в моем случае из Gitlab Container Registry), нужно было предоставить Kubernetes необходимые секреты для аутен- тификации в этом реестре. Для этого достаточно было создать Deploy Tokens и выполнить небольшую настройку согласно
___[инструкции](https://chris-vermeulen.com/using-gitlab-registry-with-kubernetes/)___.