package pkg

import "runtime"

func FileTest() (string, int, bool) {
	_, filePath, lineNum, ok := runtime.Caller(1)
	return filePath, lineNum, ok
}
