package executor

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ExecArgs struct {
	ExecutorType
	ID     uuid.UUID
	Object interface{}
	QueryArgs
	SetColumnArgs
}

type QueryArgs struct {
	QueryString string
}

type SetColumnArgs struct {
	Field string
	Value interface{}
}

type Executor interface {
	Exec(ctx context.Context, conn *gorm.DB, args ExecArgs) error
}
