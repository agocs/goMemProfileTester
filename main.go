package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/pprof"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/readMem0/", useNeither)
	http.HandleFunc("/readMem1/", useMemReadMemStats)
	http.HandleFunc("/readMem2/", useMemPprof)
	http.ListenAndServe(":8080", nil)

}

func useNeither(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(*wasteMem())
	log.Print("done")
	w.Write(response)
}

func useMemReadMemStats(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(*wasteMem())
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)
	log.Print(ms.Alloc)
	w.Write(response)
}

func useMemPprof(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(*wasteMem())
	buf := bytes.Buffer{}
	pprof.Lookup("heap").WriteTo(&buf, 1)
	//	log.Print(buf.String())
	log.Print("done")
	w.Write(response)
}

func wasteMem() *[]float64 {
	randBuffer := []float64{}
	//fill up a very large array with random values
	for i := 0; i < 204800; i++ {
		randBuffer = append(randBuffer, rand.Float64())
	}
	//sort them in place
	sort.Float64s(randBuffer)
	return &randBuffer
}
