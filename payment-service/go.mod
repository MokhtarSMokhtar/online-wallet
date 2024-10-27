module github.com/MokhatrSMokhtar/online-wallet/payment-service

go 1.23.1

require (
	github.com/joho/godotenv v1.5.1
	github.com/oklog/ulid/v2 v2.1.0
	go.mongodb.org/mongo-driver v1.17.1
)

require (
	github.com/MokhtarSMokhtar/online-wallet/comman v0.0.0-00010101000000-000000000000
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.17.0 // indirect

)

replace github.com/MokhtarSMokhtar/online-wallet/comman => ../common
