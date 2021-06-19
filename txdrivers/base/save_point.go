package base

import "fmt"

type SavePointIface interface {
	Create(id string) string
	Release(id string) string
	Rollback(id string) string
}

type defaultSavePoint struct{}

func (dsp *defaultSavePoint) Create(id string) string {
	return fmt.Sprintf("SAVEPOINT %s", id)
}
func (dsp *defaultSavePoint) Release(id string) string {
	return fmt.Sprintf("RELEASE SAVEPOINT %s", id)
}
func (dsp *defaultSavePoint) Rollback(id string) string {
	return fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", id)
}
