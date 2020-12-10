package url_test

import (
	"context"
	"testing"

	"github.com/kpr-hellofresh/tinyurl/internal/domain/url"
	"github.com/stretchr/testify/assert"
)

func TestShortener_Shorten(t *testing.T) {
	t.Run("Hashes a string", func(t *testing.T) {
		expected := "098f6bcd4621d373cade4e832627b4f6"
		actual, err := url.Shortener{}.Shorten(context.Background(), "test")

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}
