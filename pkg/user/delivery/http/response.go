package http

import "github.com/IamStubborN/test/models"

type userResponse struct {
	*models.User
	DepositCount uint64  `json:"depositCount"`
	DepositSum   float64 `json:"depositSum"`
	BetCount     uint64  `json:"betCount"`
	BetSum       float64 `json:"betSum"`
	WinCount     uint64  `json:"winCount"`
	WinSum       float64 `json:"winSum"`
}
