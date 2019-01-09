package main

import (
	"flag"
	"github.com/lempiy/privatax/lib"
	"log"
	"strings"
)

func main() {
	dbfs := flag.String("dbf", "", "Comma-separated paths to Privat Bank *.dbf source files")
	flag.Parse()
	if dbfs == nil || *dbfs == "" {
		log.Fatal("'dbf' is required flag")
	}
	dbfSources := strings.Split(*dbfs, ",")
	count, err := lib.Count(dbfSources)
	if err != nil {
		log.Fatalf("cannot count tax amount. Err: %s", err)
	}
	log.Printf("Tax amount is %.2f UAH", count)
}
