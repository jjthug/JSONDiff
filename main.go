package main

import (
	"JSONCompare/config"
	"JSONCompare/diff"
	"JSONCompare/examples"
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	// DEFINE THE LEVELS THAT ARE THE MOST DEEP OR WIDE, and need to parallelize
	levelSet := make(map[int]struct{})

	for _, level := range config.LEVELS {
		levelSet[level] = struct{}{}
	}

	json1, err1 := diff.ParseJSON(examples.JsonStr_1)
	json2, err2 := diff.ParseJSON(examples.JsonStr_2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing JSON:", err1, err2)
		return
	}

	start := time.Now()
	diff := diff.CompareJSON(json1, json2, levelSet)
	elapsed := time.Since(start)

	// Print the diff and elapsed time
	diffBytes, _ := json.MarshalIndent(diff, "", "  ")
	fmt.Println("Diff:", string(diffBytes))
	fmt.Println("Elapsed time:", elapsed)
}
