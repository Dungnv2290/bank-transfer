package presenter

import (
	"github.com/dungnguyen/bank-transfer/domain"
	"github.com/dungnguyen/bank-transfer/usecase"
)

type findAccountBalancePresenter struct{}

func NewFindAccountBalancePresenter() usecase.FindAccountBalancePresenter {
	return findAccountBalancePresenter{}
}

func (f findAccountBalancePresenter) Output(balance domain.Money) usecase.FindAccountBalanceOutput {
	return usecase.FindAccountBalanceOutput{Balance: balance.Float64()}
}
