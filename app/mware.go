package app

import (
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/mware"
	"github.com/IamStubborN/test/pkg/mware/instance"
	"github.com/IamStubborN/test/pkg/responder"
)

func initializeMiddleWares(l logger.Logger, r responder.Responder) mware.MWare {
	return instance.NewMiddleWare(l, r)
}
