package sqlstore

import (
	"context"

	"github.com/mattermost/gorp"
)

type contextValue string

type storeContextKey string

const (
	useMaster contextValue = "useMaster"
)

func (ss *SqlStore) DbFromContext(ctx context.Context) *gorp.DbMap {
	if hasMaster(ctx) {
		return ss.GetMaster()
	}
	return ss.GetReplica()
}

func hasMaster(ctx context.Context) bool {
	if v := ctx.Value(storeContextKey(useMaster)); v != nil {
		if res, ok := v.(bool); ok && res {
			return true
		}
	}
	return false
}
