package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"time"
)

var bufferSize int = 1024

func main() {

	strBufferSize := os.Args[1]

	if strBufferSize != "" {
		bufSize, err := strconv.Atoi(strBufferSize)
		if err == nil {
			bufferSize = bufSize
		} else {
			log.Printf("Unable to use %s as a buffer size. Using %d instead.", strBufferSize, bufSize)
		}
	}

	log.Printf("Starting with buffer size %d", bufferSize)


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
	for i := 0; i < bufferSize; i++ {
		randBuffer = append(randBuffer, rand.Float64())
	}
	//sort them in place
	sort.Float64s(randBuffer)
	return &randBuffer
}
