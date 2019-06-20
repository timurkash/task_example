package consts

import "github.com/timurkash/task_example/common/config"

var (
	TASKS_PORT = config.GetEnv("TASKS_PORT", "3000")
)
