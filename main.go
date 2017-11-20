package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"math/rand"
)

func main() {

	var (
		sampleSize = flag.Int("n", 5, "Sample size ")
	)

	var reader = bufio.NewReader(os.Stdin)

	buff := make([]byte, 8096, 8096)

	counters := make([]int64, 256)

	for read, err := reader.Read(buff); err != io.EOF; read, err = reader.Read(buff) {
		for i := 0; i < read; i++ {
			b := int(buff[i])
			if b != 10 {
				counters[b] = counters[b] + 1
			}
		}
	}

	sorted := sortMap(counters)

	samplesCount := int64(0)

	for _, p := range sorted {
		samplesCount += p.Value
	}

	indexes := make([]int64, *sampleSize, *sampleSize)

	for i := 0; i < len(indexes); i++ {
		indexes[i] = rand.Int63n(samplesCount)
	}


	result := make([]byte, *sampleSize, *sampleSize)

	currentCount := int64(0)
	for _, p := range sorted {
		for i, r := range indexes {
			if r >= currentCount && r < currentCount + p.Value {
				result[i] = p.Key
			}
		}
		currentCount += p.Value
	}

	for _, b := range result {
		fmt.Printf("%s", string(b))
	}
}

func sortMap(m map[byte]int64) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}

// A data structure to hold key/value pairs
type Pair struct {
	Key   byte
	Value int64
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }


