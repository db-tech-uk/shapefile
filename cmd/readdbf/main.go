package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/db-tech-uk/shapefile/dbf"
)

func readWithLib(src string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	dbf, err := dbf.NewReader(r)
	if err != nil {
		return err
	}

	fmt.Printf("Day = %s\n", dbf.Day)
	fmt.Printf("NumberOfRecords = %d\n", dbf.NumberOfRecords)
	fmt.Printf("NumberOfFields = %d\n", dbf.NumberOfFields)
	fmt.Printf("RecordLength = %d\n", dbf.RecordLength)

	for _, field := range dbf.Fields {
		fmt.Printf("  %#v\n", field)
	}

	for i := 0; i < int(dbf.NumberOfRecords); i++ {
		r, err := dbf.Next()
		if err != nil {
			return err
		}
		fmt.Printf("%#v\n", r.Attrs())
	}

	return nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if err := readWithLib(arg); err != nil {
			log.Panic(err)
		}
	}
}
