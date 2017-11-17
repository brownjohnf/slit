// slit is like cut, but smarter.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var f = flag.String("f", "", "field(s) to include")
var F = flag.String("F", "", "field(s) to exclude")
var s = flag.Int("s", 0, "lines to skip (headers, etc)")
var d = flag.String("d", "", "delimiter on which to split columns (defaults to whitespace)")
var od = flag.String("D", " ", "delimiter to use in output")
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
		include := strings.Split(*f, ",")
		exclude := strings.Split(*F, ",")

		// Unless a delimiter was passed, split on whitespace
		if *d != "" {
			fields = strings.Split(input.Text(), *d)
		} else {
			fields = strings.Fields(input.Text())
		}

		for i, _ := range fields {
			var skip bool

			if *F != "" {
				for _, n := range exclude {
					num, err := strconv.Atoi(n)
					if err != nil {
						fmt.Println(os.Stderr, err)
						os.Exit(2)
					}

					if i == num-1 {
						skip = true
						break
					}
				}
			}

			if skip {
				continue
			}

			if *f == "" {
				out = append(out, fields[i])
				continue
			}

			for _, n := range include {
				num, err := strconv.Atoi(n)
				if err != nil {
					fmt.Println(os.Stderr, err)
					os.Exit(2)
				}

				if len(fields) < num {
					if *v {
						fmt.Fprintf(os.Stderr, "%d: Insufficient fields: %v; skipping\n", line, f)
					}
					continue
				}

				if i == num-1 {
					out = append(out, fields[i])
				}
			}
		}

		// Write-out the row's output
		fmt.Println(strings.Join(out, *od))
	}
}
