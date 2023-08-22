include .env
export

echo:
	echo ${MYSQL_PORT}

run:
	docker-compose up --build 

stop:
	docker-compose down -v

conn:
	mysql -h 127.0.0.1 -p=$$MYSQL_PORT -u $$MYSQL_USER --password=$$MYSQL_PASSWORD