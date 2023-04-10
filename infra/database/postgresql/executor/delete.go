package executor

import (
	"context"
	"gorm.io/gorm"
)

type deleteExecutor struct{}

func NewDeleteExecutor() Executor {
	return &deleteExecutor{}
}

func (e *deleteExecutor) Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error {
	return conn.WithContext(ctx).Where(idStringQuery, args.ID).Delete(args.Object).Error
}
