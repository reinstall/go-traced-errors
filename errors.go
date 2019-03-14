package errors

import (
	"fmt"
	"runtime"
)

type Error struct{
	text string
}

func New(text string) Error {
	return Error {
		text: text,
	}
}

func Wrap(err Error) Error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	return New(fmt.Sprintf("%s:%d [%s]: \n%s", file, line, f.Name(), err.Error()))
}

func (e *Error) Error() string {
	return e.text
}
