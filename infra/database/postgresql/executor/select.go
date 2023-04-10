package executor

import (
	"context"
	"gorm.io/gorm"
)

type selectExecutor struct{}

func NewSelectExecutor() Executor {
	return &selectExecutor{}
}

func (e *selectExecutor) Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error {
	return conn.WithContext(ctx).Find(args.Object).Error
}
