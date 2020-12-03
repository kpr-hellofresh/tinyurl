package app_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kpr-hellofresh/tinyurl/internal/app"
	"github.com/kpr-hellofresh/tinyurl/internal/domain/url"
	"github.com/kpr-hellofresh/tinyurl/internal/domain/url/mocks"
)

func TestShortenURL_Handle(t *testing.T) {
	t.Run("Shorten Fail", func(t *testing.T) {
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
	})

	t.Run("Adder Fails", func(t *testing.T) {
		now := time.Now()

		adder := new(mocks.Adder)
		adder.On("Add", mock.Anything, url.URL{
			ID:        "mockId",
			Data:      "pippo",
			CreatedAt: now,
		}).Return(errors.New("fail"))

		srv := app.ShortenURL{
			Adder: adder,
			Shortener: app.ShortenerFunc(func(ctx context.Context, data string) (string, error) {
				return "mockId", nil
			}),
			Now: func() time.Time {
				return now
			},
		}

		_, err := srv.Handle(context.Background(), "pippo")
		assert.Error(t, err)
	})

	t.Run("Happy Path", func(t *testing.T) {
		now := time.Now()
		testURL := url.URL{
			ID:        "mockId",
			Data:      "pippo",
			CreatedAt: now,
		}

		adder := new(mocks.Adder)
		adder.On("Add", mock.Anything, testURL).Return(nil)

		srv := app.ShortenURL{
			Adder: adder,
			Shortener: app.ShortenerFunc(func(ctx context.Context, data string) (string, error) {
				return "mockId", nil
			}),
			Now: func() time.Time {
				return now
			},
		}

		u, err := srv.Handle(context.Background(), "pippo")
		assert.NoError(t, err)
		assert.Equal(t, testURL, u)
	})
}

func TestGetURL_Handle(t *testing.T) {
	t.Run("Getter fails", func(t *testing.T) {
		getter := new(mocks.Getter)
		getter.On("Get", mock.Anything, mock.Anything).Return(url.URL{}, errors.New("Getter fail"))

		srv := app.GetURL{
			Getter: getter,
		}

		_, err := srv.Handle(context.Background(), "mockId")
		assert.Error(t, err)
	})

	t.Run("Happy Path", func(t *testing.T) {
		now := time.Now()
		mockId := "mockId"

		testURL := url.URL{
			ID:        mockId,
			Data:      "pippo",
			CreatedAt: now,
		}

		getter := new(mocks.Getter)
		getter.On("Get", mock.Anything, mockId).Return(testURL, nil)

		srv := app.GetURL{
			Getter: getter,
		}

		u, err := srv.Handle(context.Background(), mockId)
		assert.NoError(t, err)
		assert.Equal(t, testURL, u)
	})
}
