package app

import (
	"github.com/go-chi/chi"

	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/transaction"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/IamStubborN/test/pkg/user/delivery/http"
)

func registerUserHandlers(
	router chi.Router,
	l logger.Logger,
	r responder.Responder,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase) {
	http.RegisterUserHandler(router, l, r, uuc, duc, tuc)
}
