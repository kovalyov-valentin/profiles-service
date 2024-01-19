compose:
	docker compose up -d
upmigrate:
	migrate -path internal/db/migration -database 'postgres://mobile:password@localhost:5040/profilesdb?sslmode=disable' up
downmigrate:
	migrate -path internal/db/migration -database 'postgres://mobile:password@localhost:5040/profilesdb?sslmode=disable' down
run:
	go run cmd/main.go
stop:
	docker stop profiles-service
clean:
	docker rm profiles-service

.PHONY: compose upmigrate downmigrate stop clean run