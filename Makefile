include .env
export

all: run testconn

run:
	docker-compose up --build -d

stop:
	docker-compose down -v

testconn:
	bash ./check-mysql.sh