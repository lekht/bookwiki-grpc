# bookwiki-grpc

```
git clone https://github.com/lekht/bookwiki-grpc.git
```

Следующая команда запустит 2 контейнера: приложение и mysql container. Необходимо подождать минуту (или больше). \
Сброка контейнеров осуществляется при помощи docker-compose.

```
make all
```

Чтобы проверить работоспособность, предлагаю воспользоваться следующим инструментом [grpcurl](https://github.com/fullstorydev/grpcurl).

Метод GetBooksByAuthor по автору осуществляет поиск его книг.
```
grpcurl --plaintext -d '{"author_name":"Александр Дюма"}' 0.0.0.0:8088 wiki_grpc.WikiHandler/GetBooksByAuthor
```

Метод GetAuthorsByBook по книге возвращает авторов.
```
grpcurl --plaintext -d '{"book_title":"Три мушкетёра"}' 0.0.0.0:8088 wiki_grpc.WikiHandler/GetAuthorsByBook
```
