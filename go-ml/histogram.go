//
// read the data points to prepare for plotting...
//
// @created 2018-05-15 13:56:09
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
    "strings"

	// "gonum.org/v1/plot"
	// "gonum.org/v1/plot/plotter"
	// "gonum.org/v1/plot/vg/draw"
)

type HistData struct {
    label   string
    min     float64 // the minimum value allowed in this column
    max     float64 // the maximum value allowed in this column
    minVal  float64
    maxVal  float64
    count   int64
}

func main() {
	fin := "./data-sources/histogram-data.txt"
	fout := "histogram.png"

	data, err := readData(fin)
	if err != nil {
		log.Fatalf("could not read %s: %v", fin, err)
	}

	err = plotData(fout, data)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

func readData(path string) ([]HistData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

    data := make([]HistData, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
        vals := strings.Split(s.Text(), ",")
		if len(vals) != 6 {
			log.Printf("discarding bad data line %q", s.Text())
			continue
		}
        var err error
        hg := HistData{ label:vals[0] }

        if hg.min, err = strconv.ParseFloat(vals[1], 64); err != nil {
			log.Printf("discarding bad min data line %q : %v", s.Text(), err)
			continue
        }

        if hg.max, err = strconv.ParseFloat(vals[2], 64); err != nil {
			log.Printf("discarding bad max data line %q : %v", s.Text(), err)
			continue
        }

        if hg.minVal, err = strconv.ParseFloat(vals[3], 64); err != nil {
			log.Printf("discarding bad mix val data line %q : %v", s.Text(), err)
			continue
        }

        if hg.maxVal, err = strconv.ParseFloat(vals[4], 64); err != nil {
			log.Printf("discarding bad max val data line %q : %v", s.Text(), err)
			continue
        }

        if hg.count, err = strconv.ParseInt(vals[5], 10, 32); err != nil {
			log.Printf("discarding bad count data line %q : %v", s.Text(), err)
			continue
        }

        log.Println(hg)
		data = append(data, hg)
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}

	return data, nil
}

func plotData(fout string, data []HistData) error {
    return nil
}
