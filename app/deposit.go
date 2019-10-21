package app

import (
	"github.com/go-chi/chi"

	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/deposit/delivery/http"
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
)

func registerDepositHandlers(router chi.Router, l logger.Logger, r responder.Responder, duc deposit.UseCase) {
	http.RegisterDepositHandler(router, l, duc, r)
}
