package main
import (
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"os"
	"syscall"
	log "github.com/sirupsen/logrus"
	"time"
	"fmt"
	"runtime"
	"sync"
)

func dummyCPUUsage() {
	var a uint64
	var t = time.Now()
	for {
		t = time.Now()
		a += uint64(t.Unix())
	}
}

func dummyAllocations() {
	var d []uint64

	for {
		for i := 0; i < 2*1024*1024; i++ {
			d = append(d, 42)
		}
		time.Sleep(time.Second * 10)
		fmt.Println(len(d))
		d = make([]uint64, 0)
		runtime.GC()
		time.Sleep(time.Second * 10)
	}
}

func Counter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 200)
		counter++
	}
	wg.Done()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	go dummyAllocations()
	go dummyCPUUsage()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()
	log.Printf("you can now open http://localhost:8080/debug/pprof in your browser \n" +
		"cmd go tool pprof http://localhost:8080/debug/pprof/heap")
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")
}