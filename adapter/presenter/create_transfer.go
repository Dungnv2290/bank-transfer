package presenter

import (
	"time"

	"github.com/dungnguyen/bank-transfer/domain"
	"github.com/dungnguyen/bank-transfer/usecase"
)

type createTransferPresenter struct{}

func NewCreateTransferPresenter() usecase.CreateTransferPresenter {
	return createTransferPresenter{}
}

func (c createTransferPresenter) Output(transfer domain.Transfer) usecase.CreateTransferOutput {
	return usecase.CreateTransferOutput{
		ID:                   transfer.ID().String(),
		AccountOriginID:      transfer.AccountOriginID().String(),
		AccountDestinationID: transfer.AccountDestinationID().String(),
		Amount:               transfer.Amount().Float64(),
		CreatedAt:            transfer.CreatedAt().Format(time.RFC3339),
	}
}
