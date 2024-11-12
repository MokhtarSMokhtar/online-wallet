package grpcserver

import (
	"context"
	walletpb "github.com/MokhtarSMokhtar/online-wallet/online-wallet-protos/github.com/MokhtarSMokhtar/online-wallet-protos/wallet"
	"log"

	"github.com/MokhtarSMokhtar/online-wallet/wallet-service/internal/application/commands"
)

type WalletGRPCServer struct {
	walletpb.UnimplementedWalletServiceServer
	commandHandlers *commands.CommandHandlers
}

func NewWalletGRPCServer(commandHandlers *commands.CommandHandlers) *WalletGRPCServer {
	return &WalletGRPCServer{commandHandlers: commandHandlers}
}

func (s *WalletGRPCServer) UpdateWallet(ctx context.Context, req *walletpb.UpdateWalletRequest) (*walletpb.UpdateWalletResponse, error) {
	cmd := commands.UpdateWalletCommand{
		UserID: req.GetUserId(),
		Amount: req.GetAmount(),
		Reason: req.GetReason(),
	}

	err := s.commandHandlers.UpdateWallet(cmd)
	if err != nil {
		log.Printf("Failed to update wallet: %v", err)
		return &walletpb.UpdateWalletResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &walletpb.UpdateWalletResponse{
		Success: true,
		Message: "Wallet updated successfully",
	}, nil
}
