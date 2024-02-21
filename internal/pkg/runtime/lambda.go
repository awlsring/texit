package runtime

import "os"

func IsLambda() bool {
	_, ok := os.LookupEnv("LAMBDA_TASK_ROOT")
	return ok
}
