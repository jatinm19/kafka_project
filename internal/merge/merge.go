package merge

import (
	"bufio"
	"container/heap"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"kafka-pipeline/internal/kafkautil"
)

type Item struct {
	value   string
	file    *os.File
	scanner *bufio.Scanner
	mode    string
}

type MinHeap []*Item

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	a := strings.Split(h[i].value, ",")
	b := strings.Split(h[j].value, ",")

	switch h[i].mode {

	case "id":
		x, _ := strconv.Atoi(a[0])
		y, _ := strconv.Atoi(b[0])
		return x < y

	case "name":
		return strings.ToLower(a[1]) <
			strings.ToLower(b[1])

	case "continent":
		return strings.ToLower(a[3]) <
			strings.ToLower(b[3])
	}

	return h[i].value < h[j].value
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func RunAllMerges() {
	mergeAndPublish("tmp/id_*.chunk", "id", "id")
	mergeAndPublish("tmp/name_*.chunk", "name", "name")
	mergeAndPublish("tmp/continent_*.chunk", "continent", "continent")

	kafkautil.CloseWriters()
}

func mergeAndPublish(pattern, topic, mode string) {
	files, _ := filepath.Glob(pattern)

	h := &MinHeap{}
	heap.Init(h)

	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			continue
		}

		scanner := bufio.NewScanner(file)

		if scanner.Scan() {
			heap.Push(h, &Item{
				value:   scanner.Text(),
				file:    file,
				scanner: scanner,
				mode:    mode,
			})
		} else {
			file.Close()
		}
	}

	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)

		_ = kafkautil.Publish(topic, item.value)

		if item.scanner.Scan() {
			item.value = item.scanner.Text()
			heap.Push(h, item)
		} else {
			item.file.Close()
		}
	}
}
