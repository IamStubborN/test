package http

import (
	"encoding/json"
	"net/http"

	"github.com/IamStubborN/test/pkg/transaction"

	"github.com/IamStubborN/test/models"

	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/go-chi/chi"
)

type transactionHandler struct {
	logger      logger.Logger
	transaction transaction.UseCase
	responder   responder.Responder
}

func RegisterTransactionHandler(router chi.Router, l logger.Logger, t transaction.UseCase, r responder.Responder) {
	handler := &transactionHandler{
		logger:      l,
		transaction: t,
		responder:   r,
	}

	router.Post("/transaction", handler.addTransaction)
}

func (th *transactionHandler) addTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		th.logger.Warn(ErrEmptyBody)
		th.responder.ResponseWithError(w, ErrEmptyBody, http.StatusBadRequest)
		return
	}
	defer th.checkError(r.Body.Close)

	var t models.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		th.logger.Warn(err)
		th.responder.ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	err = th.transaction.AddTransaction(&t)
	if err != nil {
		th.logger.Warn(err)
		th.responder.ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	res := response{
		Error:   "",
		Balance: t.BalanceAfter,
	}

	th.responder.ResponsePOSTWithObject(w, res, http.StatusOK)
}

func (th *transactionHandler) checkError(f func() error) {
	if err := f(); err != nil {
		th.logger.Warn(err)
	}
}
