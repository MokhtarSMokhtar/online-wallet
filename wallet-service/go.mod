module github.com/MokhtarSMokhtar/online-wallet/wallet-service

go 1.23.1

require (
	github.com/MokhtarSMokhtar/online-wallet/comman v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/rabbitmq/amqp091-go v1.10.0
)

replace github.com/MokhtarSMokhtar/online-wallet/comman => ../common
