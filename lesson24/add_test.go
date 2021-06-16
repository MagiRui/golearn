package lesson24

import "testing"

//go test -v
//go test -cover
func TestAdd(t *testing.T) {

	if add(1, 2) != 3 {

		t.Error("error")
	}
}

//go test -bench -v
//go test -bench . -benchmem -gcflags "-N -l "
func BenchmarkAdd(b *testing.B){

	for i:=0; i<b.N;i++{

		add(1,2)
	}
}
