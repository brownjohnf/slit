// Addup adds integers in columns from stdin
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var f = flag.String("f", "1", "field(s) to select")
var s = flag.Int("s", 0, "lines to skip (headers, etc")
var d = flag.String("d", "", "delimiter on which to split columns")
var v = flag.Bool("v", false, "verbose errors")

func main() {
	flag.Parse()

	var line int

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		line++

		// Skip this line if it's part of the preamble count
		if *s >= line {
			continue
		}

		var out []string
		var fields []string
		delim := "\t"

		// Unless a delimiter was passed, split on whitespace
		if *d != "" {
			fields = strings.Split(input.Text(), *d)
			delim = *d
		} else {
			fields = strings.Fields(input.Text())
		}

		// Split the columns passed in f, and loop through them
		for _, i := range strings.Split(*f, ",") {
			index, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(os.Stderr, err)
				os.Exit(2)
			}

			// Skip this row if there aren't enough columns in it
			if len(fields) < index {
				if *v {
					fmt.Fprintf(os.Stderr, "%d: Insufficient fields: %v; skipping\n", line, f)
				}
				continue
			}

			// Add the field to the output slice
			out = append(out, fields[index-1])
		}

		// Write-out the row's output
		fmt.Println(strings.Join(out, delim))
	}
}
