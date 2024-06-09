# export APP_ENV_SECRET_KEY=secret-key
# export APP_ENV_POSTGRESQL_HOST=local_home
# export APP_ENV_POSTGRES_PORT=5400
# export APP_ENV_POSTGRES_USERNAME=tientran
# export APP_ENV_POSTGRES_PASSWORD=tien123@
# export APP_ENV_POSTGRES_DATABASE=saarm_db
# export APP_ENV_POSTGRES_TZ=Asia/Ho_Chi_Minh
# export APP_ENV_MINIO_ENDPOINT=172.28.0.0:9000
# export APP_ENV_MINIO_ACCESS_KEY=minio-root
# export APP_ENV_MINIO_SECRET_KEY=tien19217

.PHONY: run
run:
	go run cmd/main.go

swagger:
	swag init -g cmd/main.go \
	--exclude ./internal/ \
	-o ./docs/
