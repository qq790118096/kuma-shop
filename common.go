package main

import "log"

func halt(data ...interface{}) {
	log.Fatalln(data)
}
