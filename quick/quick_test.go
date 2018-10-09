package quick

import (
	"algorithm"
	"testing"
	_ "github.com/mkevac/debugcharts"
)

func TestSort(t *testing.T) {
	item := algorithm.Generate(0, 100, 10)
	Sort(item, 0, len(item)-1)
	t.Log(item)
}

/*
go test -test.bench ^BenchmarkSort$
go test -test.bench ^BenchmarkSort$ -test.benchmem -test.cpuprofile cpu.profile -test.memprofile mem.profile -test.cpu 1,2,4
*/
/*
BenchmarkSort-4         2000000000               0.00 ns/op
PASS
ok      algorithm/quick 0.033s
*/
func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	item := algorithm.Generate(0, 10000000, 5000)
	b.StartTimer()
	for i:=0;i<b.N ;i++ {
		Sort(item, 0, len(item)-1)
	}
}
