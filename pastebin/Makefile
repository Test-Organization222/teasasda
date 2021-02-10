local-test-env-up:
	docker-compose -f ./docker-compose-local-test.yml up -d
	migrate -source "file://./external/storage/migrations/postgres" -database "postgres://pastebin:pastebin@localhost:5444/pastebin?sslmode=disable" up

local-test-env-down:
	docker-compose -f ./docker-compose-local-test.yml down

local-test:
	make local-test-env-up
	go test ./... -p=1 --cover
