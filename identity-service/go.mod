module github.com/MokhtarSMokhtar/online-wallet/identity-service

go 1.23.1

require (
	github.com/MokhtarSMokhtar/online-wallet/comman v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.28.0
)

replace github.com/MokhtarSMokhtar/online-wallet/comman => ../common
