package bubble

import (
	"testing"
	"algorithm"
)

func TestSort(t *testing.T) {
	item := algorithm.Generate(0,100,10)
	Sort(item)
	t.Log(item)
}

func TestSortUsingSortPackage(t *testing.T) {
	item := algorithm.Generate(0,100,10)
	SortUsingSortPackage(array(item))
	t.Log(item)
}

/*
go test -bench ^BenchmarkSort$
go test -test.bench=^BenchmarkSort$ -test.benchmem -test.cpuprofile cpu.profile -test.memprofile mem.profile -test.cpu 1,2,4
*/
func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	item := algorithm.Generate(0,10000000,5000)
	b.StartTimer()
	for i:=0;i<b.N;i++  {
		Sort(item)
	}
}

func BenchmarkSortUsingSortPackage(b *testing.B) {
	b.StopTimer()
	item := array(algorithm.Generate(0,10000000,5000))
	b.StartTimer()
	for i:=0;i<b.N;i++ {
		SortUsingSortPackage(item)
	}
}
