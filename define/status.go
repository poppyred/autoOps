package define

import (
	"autoOps/pkg/types"
	"errors"
)

const (
	TaskSuccess types.TaskStatus = iota + 1
	TaskFail
	TaskReject
	TaskRunning
	TaskWaitingReview
)

const (
	ServiceSuccess types.ServiceStatus = 20000
	ServiceFail    types.ServiceStatus = 50000
)

var (
	UnknownError = errors.New("UNKNOWN")
	ServiceError = errors.New("ServiceError")
)
