package executor

import (
	"context"
	"gorm.io/gorm"
)

type updateExecutor struct{}

func NewUpdateExecutor() Executor {
	return &updateExecutor{}
}

func (e *updateExecutor) Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error {
	return conn.WithContext(ctx).Where(idStringQuery, args.ID).Updates(args.Object).Error
}
