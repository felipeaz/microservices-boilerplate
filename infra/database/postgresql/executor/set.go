package executor

import (
	"context"
	"gorm.io/gorm"
)

type setExecutor struct{}

func NewSetExecutor() Executor {
	return &setExecutor{}
}

func (e *setExecutor) Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error {
	return conn.WithContext(ctx).Model(args.Object).UpdateColumn(args.SetColumnArgs.Field, args.SetColumnArgs.Value).Error
}
