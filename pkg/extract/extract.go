package extract

import (
	"context"

	"github.com/office-aska/lwgc/pkg/ctxkeys"
)

// UserID - ContextからUserIDを抽出する
func UserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(ctxkeys.UserIDKey{}).(string)
	return id, ok
}
