package app

import (
	"context"

	"github.com/kpr-hellofresh/tinyurl/internal/domain/url"
)

// type Service struct {
// 	URLAdder url.Adder
// }

// func (srv Service) ShortenURL() {

// }

// func (srv Service) GetURL() {

// }

type ShortenURL struct {
	Adder url.Adder
}

func (srv ShortenURL) Handle() {

}

type GetURL struct {
	Getter url.Getter
}

func (srv GetURL) Handle(ctx context.Context, id string) (url.URL, error) {
	return srv.Getter.Get(ctx, id)
}
