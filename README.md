# L0
Проект представляет собой реализацию тестового задания L0 с хранилищем PostgreSQL и броекром Nats-streaming.

# Инструкция по запуску
```shell
make all
#Запуск контейнеров nats, postgresql и rest-сервера
make publisher
#Запуск скрипта публикации
```
# Запросы обрабатываемые сервисом
`GET /getById?id={model id}` - возвращает конкретную запись по id<br/>
`GET /getAll` - возвращает все имеющиеся записи

# Интерфейс отображения
В проекте есть ex.html, который создан для демонстрации полученных данных и их запроса
![Image alt](https://github.com/alexGrap/l0/blob/main/readMeImage/Screenshot%20from%202023-07-30%2015-34-05.png)
