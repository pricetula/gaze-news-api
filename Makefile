# dev is used to start the development environment using Docker Compose
dev:
	docker compose --env-file .env -f docker-compose.yml up

# build is used to build the Docker image for the project
build:
	docker build --progress=plain -t gaze-news --no-cache .

# test is used to run the tests in the project
test:
	go test ./...

# migrate-up is used to apply migrations to the database
migrate-up:
	migrate -path "./migrations/" -database postgres://gaze_news:gaze_news@0.0.0.0:5433/gaze_news?sslmode=disable up

# migrate-down is used to revert the last migration applied to the database
migrate-down:
	migrate -path "./migrations/" -database postgres://gaze_news:gaze_news@0.0.0.0:5433/gaze_news?sslmode=disable down

# Create a new migration file with => migrate create -ext sql -dir migrations migration_file_name
