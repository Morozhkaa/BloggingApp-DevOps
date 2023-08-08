## Post-service

Post-service отвечает за публикацию постов.
- Видеть все посты может каждый посетитель сайта, но публиковать/редактировать/удалять свои посты – только зарегистрированный пользователь.
- Зарегистрирован пользователь или нет мы понимаем с помощью auth-middleware: клиент обращается к auth-service, передавая токены через cookies, там их проверяют на валидность и, если все хорошо и такой пользователь есть в системе, возвращают информацию о нем. После этого post-service проверяет, есть ли у нас права на запрашиваемое действие.
- Удалять и редактировать чужие записи могут только участники с ролью admin или moderator.
- Все посты хранятся в отдельной базе данных (вместе с комментариями к ним).