package starter

import "testing"

func TestEmpty(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Panic for empty name is not work")
		}
	}()
	var v Starter
	Register("", v)
}

type testStarter struct{}

func (t testStarter) Test() error {
	return nil
}
func (t testStarter) Set() {
}

func TestDublicate(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Panic for empty name is not work")
		}
	}()
	var v testStarter
	Register("SameName", v)
	Register("SameName", v)
	Register("SameName", v)
	Register("SameName", v)
}
