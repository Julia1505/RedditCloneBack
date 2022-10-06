# RedditCloneBack
## Bachkend для клона реддита
* Задание для прокачки скилов на Go

---
### API:
1. POST /api/register - регистрация
2. POST /api/login - логин
3. GET /api/posts/ - список всех постов
4. POST /api/posts/ - добавление поста (есть 2 типа: с урлом и с текстом)
5. GET /api/posts/{CATEGORY_NAME} - список постов категории 
6. GET /api/post/{POST_ID} - конкретный пост с комментариями 
7. POST /api/post/{POST_ID} - добавление комментария
8. DELETE /api/post/{POST_ID}/{COMMENT_ID} - удаление комментария
9. GET /api/post/{POST_ID}/upvote - рейтинг поста вверх 
10. GET /api/post/{POST_ID}/downvote - рейтинг поста вниз
11. GET /api/post/{POST_ID}/unvote - удаление голоса
11. DELETE /api/post/{POST_ID} - удаление поста
12. GET /api/user/{USER_LOGIN} - список постов пользователя

---
### Сущности
1. Пользователь
2. Сессия (получается при авторизации)
3. Пост
4. Коммент к посту

