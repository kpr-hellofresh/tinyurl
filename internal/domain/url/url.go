package url

import (
	"context"
	"time"
)

type URL struct {
	ID        string
	Data      string
	CreatedAt time.Time
}

type Getter interface {
	Get(ctx context.Context, id string) (URL, error)
}

type Adder interface {
	Add(ctx context.Context, url URL) error
}
