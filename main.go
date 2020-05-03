package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/goMidas/midas"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const CLOCKS_PER_SEC = 1000000

func load_data(src, dst, times *[]int, inputFile string, undirected bool) {
	infile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	reader := csv.NewReader(infile)
	var s, d, t int
	if undirected == false {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			if s, err = strconv.Atoi(record[0]); err != nil {
				log.Fatal(err)
			}
			if d, err = strconv.Atoi(record[1]); err != nil {
				log.Fatal(err)
			}
			if t, err = strconv.Atoi(record[2]); err != nil {
				log.Fatal(err)
			}
			*src = append(*src, s)
			*dst = append(*dst, d)
			*times = append(*times, t)
		}
	} else {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			if s, err = strconv.Atoi(record[0]); err != nil {
				log.Fatal(err)
			}
			if d, err = strconv.Atoi(record[1]); err != nil {
				log.Fatal(err)
			}
			if t, err = strconv.Atoi(record[2]); err != nil {
				log.Fatal(err)
			}
			*src = append(*src, s)
			*dst = append(*dst, d)
			*times = append(*times, t)
			*src = append(*src, d)
			*dst = append(*dst, s)
			*times = append(*times, t)
		}
	}
}

func main() {
	inputFile := flag.String("input", "", "Input File. (Required)")
	outputFile := flag.String("output", "scores.txt", "Output File. Default is scores.txt")
	rows := flag.Int("rows", 2, "Number of rows/hash functions. Default is 2")
	buckets := flag.Int("buckets", 769, "Number of buckets. Default is 769")
	alpha := flag.Float64("alpha", 0.6, "Alpha: Temporal Decay Factor. Default is 0.6")
	norelations := flag.Bool("norelations", false, "To run Midas instead of Midas-R.")
	undirected := flag.Bool("undirected", false, "If graph is undirected.")

	flag.Parse()
	if *inputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *rows < 1 {
		log.Fatal("Number of hash functions should be positive.")
	}
	if *buckets < 2 {
		log.Fatal("Number of buckets should be at least 2")
	}
	if *alpha <= 0 || *alpha >= 1 {
		log.Fatal("Alpha: Temporal Decay Factor must be between 0 and 1.")
	}

	var src, dst, times []int

	if *undirected == true {
		load_data(&src, &dst, &times, *inputFile, true)
	} else {
		load_data(&src, &dst, &times, *inputFile, false)
	}
	fmt.Println("Finished Loading Data from", *inputFile)

	if *norelations == true {
		startTime := time.Now()
		scores := midas.Midas(&src, &dst, &times, *rows, *buckets)
		fmt.Println("Time taken:", (time.Now().Sub(startTime))/CLOCKS_PER_SEC, "s")

		fmt.Println("Writing Anomaly Scores to", *outputFile)
		file, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		for _, score := range *scores {
			writer.WriteString(fmt.Sprintf("%f\n", score))
		}
		writer.Flush()
	} else {
		startTime := time.Now()
		scores := midas.MidasR(&src, &dst, &times, *rows, *buckets, *alpha)
		fmt.Println("Time taken:", (time.Now().Sub(startTime))/CLOCKS_PER_SEC, "s")

		fmt.Println("Writing Anomaly Scores to", *outputFile)
		file, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		for _, score := range *scores {
			writer.WriteString(fmt.Sprintf("%f\n", score))
		}
		writer.Flush()
	}
}
