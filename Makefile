export APP_ENV_POSTGRES_HOST=localhost
export APP_ENV_POSTGRES_PORT=5400
export APP_ENV_POSTGRES_USERNAME=tientran
export APP_ENV_POSTGRES_PASSWORD=tien123@
export APP_ENV_POSTGRES_DATABASE=saarm_db
export APP_ENV_POSTGRES_TZ=Asia/Ho_Chi_Minh

server:
	go run cmd/main.go
swagger:
	swag init -g cmd/main.go \
	--exclude ./internal/ \
	-o ./docs/
