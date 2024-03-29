## Core

### Description

Cодержит код всех микросервисов с написанными к ним Dockerfile и .gitlab-ci.yaml файлами.

### Application assembly

При написании Dockerfile я использовала двухэтапную сборку, что позволило сократить общий размер образа примерно в 55 раз.

В файле .gitlab-ci.yml определила сценарии, которые должны выполняться во время конвейера CI/CD инструкции о том, где приложение должно быть развернуто. Я разделила сценарий на две стадии:

___build___: 
- Из тегов к коммитам (которые мы автоматически делаем ниже) получаем последнюю версию приложения. В случае ее отсутствия устанавливаем начальную с значением v0.1.0.
- Поднимаем ее, увеличивая значение Patch.
- Выполняем вход в Container Registry с помощью команды docker login.
- Собираем образ и с нужным тегом-версией сохраняем в Container Registry.

___deploy___:
- Проверяем, что кластер доступен
- Клонируем репозиторий с манифестами
- Выполняем развёртывание, пробрасывая env переменные, где нужно
В результате, при получении каких-либо изменений в коде у нас в кластере Kubernetes автоматически поднимется новая версия приложения на основе нового, взятого из Container Registry образа с увеличенным тегом-версией.