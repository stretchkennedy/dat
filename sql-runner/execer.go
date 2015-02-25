package runner

import "github.com/mgutz/dat"

// Execer implements Executable
type Execer struct {
	runner
	builder dat.Builder
}

func NewExecer(runner runner, builder dat.Builder) *Execer {
	return &Execer{runner, builder}
}

func (ex *Execer) Exec() (*dat.Result, error) {
	return exec(ex.runner, ex.builder)
}

func (ex *Execer) QueryScalar(destinations ...interface{}) error {
	return queryScan(ex.runner, ex.builder, destinations...)
}
func (ex *Execer) QuerySlice(dest interface{}) error {
	return querySlice(ex.runner, ex.builder, dest)
}
func (ex *Execer) QueryStruct(dest interface{}) error {
	return queryStruct(ex.runner, ex.builder, dest)
}
func (ex *Execer) QueryStructs(dest interface{}) error {
	return queryStructs(ex.runner, ex.builder, dest)
}