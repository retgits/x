// Package main implements a csv comparator. There are two commandline flags
// -f1
// -f2
// The tool reads both files and creates a fnv64 hash of the rows. After that,
// it loops through the second file to find all records that don't exist in
// the first. If the hash of a record is found, and the records thus match,
// the instances are removed from the list. That means at the end the only
// items that don't have matching hashes are listed.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strings"
)

var (
	f1 string
	f2 string
)

func main() {
	// Parse all the flags
	flag.StringVar(&f1, "f1", "", "the first csv file to compare")
	flag.StringVar(&f2, "f2", "", "the second csv file to compare")
	flag.Parse()

	if len(f1) == 0 && len(f2) == 0 {
		fmt.Println("file flags 'f1' and/or 'f2' have not been set")
		os.Exit(1)
	}

	// Open the files and create a map from it
	masterRecordMap := mapRecords(f1)
	newRecordMap := mapRecords(f2)
	discrepancies := make(map[uint64][]string, 0)

	// Find all records that don't exist in the master map
	// where the records match, the instances are removed from
	// the master map. That means at the end the discrepancies
	// map contains all items that didn't exist in the master map
	// and the master map only contains items that no longer exist
	for key, val := range newRecordMap {
		if _, ok := masterRecordMap[key]; ok {
			delete(masterRecordMap, key)
		} else {
			discrepancies[key] = val
		}
	}

	if len(discrepancies) > 0 {
		fmt.Printf("The following items exist in %s but not in %s", f1, f2)
		for key, val := range discrepancies {
			fmt.Printf("%d -- %s\n", key, val)
		}
	} else {
		fmt.Printf("No new items exist in %s\n", f2)
	}

	if len(masterRecordMap) > 0 {
		fmt.Printf("\nThe following items no longer exist in %s\n", f2)
		for key, val := range masterRecordMap {
			fmt.Printf("%d -- %s\n", key, val)
		}
	}
}

func mapRecords(filename string) map[uint64][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %s: %s\n", filename, err)
		os.Exit(1)
	}

	r := csv.NewReader(file)

	m := make(map[uint64][]string, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		m[hash(strings.Join(record, ""))] = record
	}

	return m
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
