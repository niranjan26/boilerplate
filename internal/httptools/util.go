package httptools

import (
	"context"
)

func GetMuxValue(ctx context.Context) map[string]string {
	if rv := ctx.Value("vars"); rv != nil {
		return rv.(map[string]string)
	}

	return nil
}
