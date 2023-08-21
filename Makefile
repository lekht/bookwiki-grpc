include .env
export

echo:
	echo ${MYSQL_PORT}

run:
	docker-compose --env-file .env up --build -d

stop:
	docker-compose down -v

con:
	mysql -h 127.0.0.1 -p=$$MYSQL_PORT -u $$MYSQL_USER --password=$$MYSQL_PASSWORD