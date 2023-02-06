package meta

import (
	"context"
)

type Request interface {
	Path() string
	Method() string
}

type Requester interface {
	Request(ctx context.Context, r Request) ([]byte, error)
}
