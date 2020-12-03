package app

import (
	"context"
	"time"

	"github.com/kpr-hellofresh/tinyurl/internal/domain/url"
)

type Shortener interface {
	Shorten(ctx context.Context, data string) (string, error)
}

type ShortenerFunc func(ctx context.Context, data string) (string, error)

func (fn ShortenerFunc) Shorten(ctx context.Context, data string) (string, error) {
	return fn(ctx, data)
}

type GetterFunc func(ctx context.Context, id string) (url.URL, error)

func (fn GetterFunc) Get(ctx context.Context, id string) (url.URL, error) {
	return fn(ctx, id)
}

// type Service struct {
// 	URLAdder url.Adder
// }

// func (srv Service) ShortenURL() {

// }

// func (srv Service) GetURL() {

// }

type ShortenURL struct {
	Adder     url.Adder
	Shortener Shortener
	Now       func() time.Time
}

func (srv ShortenURL) Handle(ctx context.Context, longUrl string) (url.URL, error) {
	id, err := srv.Shortener.Shorten(ctx, longUrl)
	if err != nil {
		return url.URL{}, err
	}

	u := url.URL{
		ID:        id,
		Data:      longUrl,
		CreatedAt: srv.Now(),
	}

	if err := srv.Adder.Add(ctx, u); err != nil {
		return url.URL{}, err
	}

	return u, nil
}

type GetURL struct {
	Getter url.Getter
}

func (srv GetURL) Handle(ctx context.Context, id string) (url.URL, error) {
	return srv.Getter.Get(ctx, id)
}
