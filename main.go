package main

import "github.com/linesmerrill/fitgirl-repacks-ui/db"

func main() {
	// initialize db connection
	mydb := db.Init()

	// scraper that hits the URL and saves data to the db
	scraper{mydb}.execute()
	defer mydb.Session.Close()
}
