package grpcclient

import (
	"log"
	"sync"

	walletpb "github.com/MokhtarSMokhtar/online-wallet/online-wallet-protos/github.com/MokhtarSMokhtar/online-wallet-protos/wallet"
	"google.golang.org/grpc"
)

var (
	clientInstance walletpb.WalletServiceClient
	once           sync.Once
)

func GetWalletServiceClient() walletpb.WalletServiceClient {
	once.Do(func() {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to Wallet Service: %v", err)
		}
		clientInstance = walletpb.NewWalletServiceClient(conn)
	})
	return clientInstance
}
