compose:
	docker compose up -d
upmigrate:
	migrate -path internal/db/migration -database 'postgres://mobile:password@localhost:5040/profilesdb?sslmode=disable' up
downmigrate:
	migrate -path internal/db/migration -database 'postgres://mobile:password@localhost:5040/profilesdb?sslmode=disable' down
stop:
	docker stop profiles-service
clean:
	docker rm profiles-service

.PHONY: compose upmigrate downmigrate stop clean