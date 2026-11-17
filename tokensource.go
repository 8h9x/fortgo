package fortgo

import "context"

type TokenSource interface {
    Token(ctx context.Context) (string, error)
}
