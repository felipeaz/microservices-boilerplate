package executor

type ExecutorType string

var (
	CreateType ExecutorType = "create"
	UpdateType ExecutorType = "update"
	SetType    ExecutorType = "set"
	SelectType ExecutorType = "select"
	RawType    ExecutorType = "raw"
	DeleteType ExecutorType = "delete"
)

func NewExecutor(executorType ExecutorType) Executor {
	switch executorType {
	case CreateType:
		return NewCreateExecutor()
	case UpdateType:
		return NewUpdateExecutor()
	case SetType:
		return NewSetExecutor()
	case SelectType:
		return NewSelectExecutor()
	case RawType:
		return NewRawExecutor()
	case DeleteType:
		return NewDeleteExecutor()
	}
	return nil
}
