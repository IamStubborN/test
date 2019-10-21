package daemon

import (
	"context"
)

type Daemon interface {
	Run(ctx context.Context)
}
