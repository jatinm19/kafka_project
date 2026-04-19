package merge

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
)

type Item struct {
	value   string
	file    *os.File
	scanner *bufio.Scanner
}

type MinHeap []*Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(*Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func RunAllMerges() {
	mergeFiles("tmp/id.chunk", "output/id_sorted.txt")
	mergeFiles("tmp/name.chunk", "output/name_sorted.txt")
	mergeFiles("tmp/continent.chunk", "output/continent_sorted.txt")
}

func mergeFiles(pattern string, output string) {
	files, _ := filepath.Glob(pattern)

	h := &MinHeap{}
	heap.Init(h)

	for _, f := range files {
		file, _ := os.Open(f)
		scanner := bufio.NewScanner(file)

		if scanner.Scan() {
			heap.Push(h, &Item{
				value:   scanner.Text(),
				file:    file,
				scanner: scanner,
			})
		}
	}

	os.MkdirAll("output", 0755)
	out, _ := os.Create(output)
	defer out.Close()

	w := bufio.NewWriter(out)
	defer w.Flush()

	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)
		fmt.Fprintln(w, item.value)

		if item.scanner.Scan() {
			item.value = item.scanner.Text()
			heap.Push(h, item)
		} else {
			item.file.Close()
		}
	}
}
