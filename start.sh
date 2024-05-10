echo "GO download"
go mod download

CGO_ENABLED=0 GOOS=linux GOARCH=amd64

export GOOGLE_APPLICATION_CREDENTIALS="/opt/google-cloud/service-account.json"

echo "Run Application"
go run cmd/main.go

echo "Running go appp"
