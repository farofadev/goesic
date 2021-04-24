up:
	docker-compose up -d

down:
	docker-compose down

prune:
	docker-compose down -v

logs:
	docker-compose logs --tail=120 -f

logs-app:
	docker-compose logs --tail=120 --no-log-prefix -f app

logs-db:
	docker-compose logs --tail=120 --no-log-prefix -f database

logs-database:
	docker-compose logs --tail=120 --no-log-prefix -f database