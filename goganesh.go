package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	Score int
	Title string
}

func (it *Item) String() string {
	return fmt.Sprintf("%d:%s", it.Score, it.Title)
}

// Transforms a stringified item to an actual Item instance.
func Itemify(s string) (*Item, error) {
	bits := append(strings.SplitN(s, ":", 2), "")
	p, err := strconv.Atoi(bits[0])
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

// Update the rank of the item at the specified position and move it up
// to its correct position.
func rankUp(items []*Item, itemPos int) {
	items[itemPos].Score++
	for i := itemPos - 1; i >= 0; i-- {
		if items[i].Score > items[itemPos].Score {
			break
		} else {
			items[i], items[i+1] = items[i+1], items[i]
			itemPos--
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
	itemPos := searchByTitle(items, title)
	// If the item was not found, set it at the end with a null score.
	// Doing so will allow it to go up the lowest scored items.
	if itemPos == -1 {
		itemPos = len(items)
		items = append(items, &Item{0, title})
	}
	rankUp(items, itemPos)
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
