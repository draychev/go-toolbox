package isdev

import "runtime"

func IsDev() bool {
	return runtime.GOOS == "darwin"
}
