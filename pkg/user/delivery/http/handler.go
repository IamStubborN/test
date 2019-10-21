package http

import (
	"encoding/json"
	"net/http"

	"github.com/IamStubborN/test/pkg/deposit"
	"github.com/IamStubborN/test/pkg/transaction"

	"github.com/IamStubborN/test/models"

	"github.com/IamStubborN/test/pkg/responder"

	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/user"
	"github.com/go-chi/chi"
)

type userHandler struct {
	logger      logger.Logger
	user        user.UseCase
	deposit     deposit.UseCase
	transaction transaction.UseCase
	responder   responder.Responder
}

func RegisterUserHandler(
	router chi.Router,
	l logger.Logger,
	r responder.Responder,
	uuc user.UseCase,
	duc deposit.UseCase,
	tuc transaction.UseCase) {
	handler := &userHandler{
		logger:      l,
		responder:   r,
		user:        uuc,
		deposit:     duc,
		transaction: tuc,
	}

	router.Post("/user/create", handler.addUser)
	router.Post("/user/get", handler.getUser)
}

func (uh *userHandler) addUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		uh.logger.Warn(ErrEmptyBody)
		uh.responder.ResponseWithError(w, ErrEmptyBody, http.StatusBadRequest)
		return
	}
	defer uh.checkError(r.Body.Close)

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		uh.logger.Warn(err)
		uh.responder.ResponseWithError(w, ErrUnmarshalBody, http.StatusBadRequest)
		return
	}

	err = uh.user.AddUser(&u)
	if err != nil {
		uh.logger.Warn(err)
		uh.responder.ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	uh.responder.ResponseOK(w, http.StatusOK)
}

func (uh *userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		uh.logger.Warn(ErrEmptyBody)
		uh.responder.ResponseWithError(w, ErrEmptyBody, http.StatusBadRequest)
		return
	}
	defer uh.checkError(r.Body.Close)

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		uh.logger.Warn(err)
		uh.responder.ResponseWithError(w, ErrUnmarshalBody, http.StatusBadRequest)
		return
	}

	us, err := uh.user.GetUser(u.ID)
	if err != nil {
		uh.logger.Warn(err)
		uh.responder.ResponseWithError(w, ErrGetUser, http.StatusBadRequest)
		return
	}

	depCount, depSum := uh.deposit.GetDepositCountAndSum(us.ID)
	winCount, winSum := uh.transaction.GetWinCountAndSum(us.ID)
	betCount, betSum := uh.transaction.GetBetCountAndSum(us.ID)

	ur := userResponse{
		User:         us,
		DepositCount: depCount,
		DepositSum:   depSum,
		BetCount:     betCount,
		BetSum:       betSum,
		WinCount:     winCount,
		WinSum:       winSum,
	}

	uh.responder.ResponseGETWithObject(w, ur, http.StatusOK)
}

func (uh userHandler) checkError(f func() error) {
	if err := f(); err != nil {
		uh.logger.Warn(err)
	}
}
