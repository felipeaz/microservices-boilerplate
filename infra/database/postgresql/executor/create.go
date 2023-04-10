package executor

import (
	"context"
	"gorm.io/gorm"
)

type createExecutor struct{}

func NewCreateExecutor() Executor {
	return &createExecutor{}
}

func (e *createExecutor) Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error {
	return conn.WithContext(ctx).Create(args.Object).Error
}
