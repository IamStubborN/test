package http

import (
	"encoding/json"
	"net/http"

	"github.com/IamStubborN/test/pkg/deposit"

	"github.com/IamStubborN/test/models"

	"github.com/IamStubborN/test/pkg/responder"

	"github.com/IamStubborN/test/pkg/logger"
	"github.com/go-chi/chi"
)

type depositHandler struct {
	logger    logger.Logger
	deposit   deposit.UseCase
	responder responder.Responder
}

func RegisterDepositHandler(router chi.Router, l logger.Logger, d deposit.UseCase, r responder.Responder) {
	handler := &depositHandler{
		logger:    l,
		deposit:   d,
		responder: r,
	}

	router.Post("/user/deposit", handler.addDeposit)
}

func (uh *depositHandler) addDeposit(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		uh.logger.Warn(ErrEmptyBody)
		uh.responder.ResponseWithError(w, ErrEmptyBody, http.StatusBadRequest)
		return
	}
	defer uh.checkError(r.Body.Close)

	var d models.Deposit
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		uh.logger.Warn(err)
		uh.responder.ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	err = uh.deposit.AddDeposit(&d)
	if err != nil {
		uh.logger.Warn(err)
		uh.responder.ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	res := response{
		Error:   "",
		Balance: d.BalanceAfter,
	}

	uh.responder.ResponsePOSTWithObject(w, res, http.StatusOK)
}

func (uh depositHandler) checkError(f func() error) {
	if err := f(); err != nil {
		uh.logger.Warn(err)
	}
}
