dev:
	docker compose --env-file .env -f docker-compose.yml up

prod:
	docker run -p 3030:3030 -e DB_USER=gaze_news -e DB_PASSWORD=gaze_news -e DB_NAME=gaze_news -e PORT=3030 -e APP_ENV=development gaze-news:main

build:
	docker build --progress=plain -t gaze-news --no-cache .

test:
	go test ./...
# Create a new migration file with => migrate create -ext sql -dir migrations migration_file_name
# make sure migrator docker image is built
migrate-up:
	migrate -path "./migrations/" -database postgres://gaze_news:gaze_news@0.0.0.0:5433/gaze_news?sslmode=disable up

# make sure migrator docker image is built
migrate-down:
	migrate -path "./migrations/" -database postgres://gaze_news:gaze_news@0.0.0.0:5433/gaze_news?sslmode=disable down

# docker run -p 3030:3030 -e DB_USER=gaze_news -e DB_PASSWORD=gaze_news -e DB_NAME=gaze_news -e PORT=3030 -e APP_ENV=development gaze_news-api