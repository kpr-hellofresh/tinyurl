package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kpr-hellofresh/tinyurl/internal/app"
	"github.com/kpr-hellofresh/tinyurl/internal/domain/url/mocks"
)

func TestGetURL_Handle(t *testing.T) {
	adder := new(mocks.Adder)
	adder.On("Add", mock.Anything, mock.Anything).Return(nil)

	srv := app.ShortenURL{
		Adder: adder,
		Shortener: app.ShortenerFunc(func(ctx context.Context, data string) (string, error) {
			return "", errors.New("fail")
		}),
	}

	_, err := srv.Handle(context.Background(), "pippo")
	assert.Error(t, err)
}
