package insertion

import (
	"testing"
	"algorithm"
)

func TestSort(t *testing.T) {
	item := algorithm.Generate(0,100,10)
	Sort(item)
	t.Log(item)
}
/*
go test -test.bench ^BenchmarkSort$
go test -test.bench ^BenchmarkSort$ -test.benchmem -test.cpuprofile cpu.profile -test.memprofile mem.profile -test.cpu 1,2,4
*/
func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	item := algorithm.Generate(0,10000000,5000)
	Sort(item)
}
