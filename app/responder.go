package app

import (
	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
	"github.com/IamStubborN/test/pkg/responder/instance"
)

func initializeResponder(l logger.Logger) responder.Responder {
	return instance.NewJSONResponder(l)
}
