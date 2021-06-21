package base

import "fmt"

type SavePointIface interface {
	Create(id string) string
	Release(id string) string
	Rollback(id string) string
}

type DefaultSavePoint struct{}

func (dsp *DefaultSavePoint) Create(id string) string {
	return fmt.Sprintf("SAVEPOINT %s", id)
}
func (dsp *DefaultSavePoint) Release(id string) string {
	return fmt.Sprintf("RELEASE SAVEPOINT %s", id)
}
func (dsp *DefaultSavePoint) Rollback(id string) string {
	return fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", id)
}
