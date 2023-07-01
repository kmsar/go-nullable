package nullable

import "testing"

func TestInt(t *testing.T) {
	a := NewNullable[int]()
	if !a.HasValue() {
		//vale is null
	}

	a.SetValue(1)
	if !a.HasValue() {
		t.Fail()
	}
	println(a.GetValue())

	b := NewNullable[int]()
	println(b.GetValueOr(2))
}
