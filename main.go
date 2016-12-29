package main

import "log"

func main() {
	api := api{}

	log.Fatal(api.startApi(1234))
}