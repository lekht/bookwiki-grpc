DROP DATABASE IF EXISTS my_db;
CREATE DATABASE my_db;
USE my_db;


DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS books;

CREATE TABLE authors (
    id INT NOT NULL AUTO_INCREMENT,
    author_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- Создание таблицы 'book'
CREATE TABLE books (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

INSERT INTO authors (author_name) VALUES
    ('Дэниел Киз'),
    ('Эрнест Хемингуэй'),
    ('Александр Дюма');

-- Вставка данных в таблицу 'book'
INSERT INTO books (title, author_id) VALUES
    ('Цветы для Элджернона', 1),
    ('Старик и море', 2),
    ('Прощай, оружие!', 2),
    ('Три мушкетёра', 3),
    ('Граф Монте-Кристо', 3),
    ('Графиня де Монсоро', 3);