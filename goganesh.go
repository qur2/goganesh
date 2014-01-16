package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var file = flag.String("file", "goganesh.list", "The file in which the rankins are stored.")
var capacity = flag.Int("capacity", 50, "The number of entry to handle.")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		// capacity-1 is to always strip the last item, so that
		// new ones get a chance at being ranked
		updateRanking(*file, *capacity-1, args[0])
	} else {
		listRanking(*file, *capacity)
	}
}

type Item struct {
	Score float64
	Title string
}

func (it *Item) String() string {
	return fmt.Sprintf("%.6f:%s", it.Score, it.Title)
}

// Some inspiration found there: http://math.stackexchange.com/questions/57429/functions-similar-to-log-but-with-results-between-0-and-1
func Grow(score float64) float64 {
	newScore := score + (math.Exp(-score/12 + 1))
	return newScore
}
func Shrink(score float64) float64 {
	return score * .86
}

// Transforms a stringified item to an actual Item instance.
func Itemify(s string) (*Item, error) {
	bits := append(strings.SplitN(s, ":", 2), "")
	p, err := strconv.ParseFloat(bits[0], 64)
	return &Item{p, bits[1]}, err
}

// Searches for a title in an array of items.
func searchByTitle(items []*Item, needle string) (pos int) {
	pos = -1
	for i := range items {
		if items[i].Title == needle {
			pos = i
			break
		}
	}
	return
}

func decay(items []*Item) {
	for _, it := range items {
		it.Score = Shrink(it.Score)
	}
}

// Update the rank of the item at the specified position and move it up
// to its correct position.
func rankUp(items []*Item, k int) {
	items[k].Score = Grow(items[k].Score)
	for i := k - 1; i >= 0; i-- {
		if items[i].Score > items[k].Score {
			break
		} else {
			items[i], items[i+1] = items[i+1], items[i]
			k--
		}
	}
}

// Outputs the titles of the ranking items.
func listRanking(file string, capacity int) {
	items, err := readLines(file, capacity)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, item := range items {
		fmt.Println(item.Title)
	}
}

// Update an item in the ranking.
func updateRanking(file string, capacity int, title string) {
	items, err := readLines(file, capacity)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	k := searchByTitle(items, title)
	decay(items[:k])
	decay(items[k+1:])
	// If the item was not found, set it at the end with a very low score.
	// Doing so will allow it to go up the lowest scored items.
	if k == -1 {
		k = len(items)
		items = append(items, &Item{0.05, title})
	}
	rankUp(items, k)
	if err := writeLines(items, file); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}

// Reads a whole file into memory and return an array of Items
func readLines(path string, maxCount int) ([]*Item, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	count := 0
	var items []*Item
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		it, err := Itemify(scanner.Text())
		if err != nil {
			return items, err
		}
		items = append(items, it)
		count++
		if count >= maxCount {
			break
		}
	}
	return items, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(items []*Item, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, item := range items {
		fmt.Fprintln(w, item)
	}
	return w.Flush()
}
