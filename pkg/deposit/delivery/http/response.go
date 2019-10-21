package http

import "github.com/IamStubborN/test/models"

type getResponse struct {
	models.User
	DepositCount uint64 `json:"depositCount"`
	DepositSum   uint64 `json:"depositSum"`
	BetCount     uint64 `json:"betCount"`
	BetSum       uint64 `json:"betSum"`
	WinCount     uint64 `json:"winCount"`
	WinSum       uint64 `json:"winSum"`
}

type response struct {
	Error   string  `json:"error"`
	Balance float64 `json:"balance"`
}
