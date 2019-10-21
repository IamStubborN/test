package app

import (
	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/go-chi/chi"
)

func registerHandlers(
	router chi.Router,
	l logger.Logger,
	r responder.Responder,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase,
) {
	registerUserHandlers(router, l, r, uuc, duc, tuc)
	registerDepositHandlers(router, l, r, duc)
	registerTransactionHandlers(router, l, r, tuc)
}
