package selection

import (
	"algorithm"
	"testing"
)

func TestSort(t *testing.T) {
	item := algorithm.Generate(0, 100, 10)
	Sort(item)
	t.Log(item)
}

/*
go test -test.bench ^BenchmarkSort$
go test -test.bench ^BenchmarkSort$ -test.benchmem -test.cpuprofile cpu.profile -test.memprofile mem.profile -test.cpu 1,2,4
*/
/*
BenchmarkSort-4         2000000000               0.00 ns/op
PASS
ok      algorithm/selection     0.103s
*/
func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	item := algorithm.Generate(0, 10000000, 5000)
	Sort(item)
}
