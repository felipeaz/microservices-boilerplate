package executor

import (
	"context"
	"gorm.io/gorm"
)

type rawExecutor struct{}

func NewRawExecutor() Executor {
	return &rawExecutor{}
}

func (e *rawExecutor) Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error {
	return conn.WithContext(ctx).Raw(args.QueryArgs.QueryString).Scan(args.Object).Error
}
