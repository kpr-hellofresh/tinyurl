package url

import (
	"context"
	"crypto/md5"
	"encoding/hex"
)

type Shortener struct{}

func (Shortener) Shorten(ctx context.Context, data string) (string, error) {
	b := md5.Sum([]byte(data))
	return hex.EncodeToString(b[:]), nil
}
