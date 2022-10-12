package main

import (
	"log"
	"os"

	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/import-boss/generators"
)

func main() {
	arguments := args.Default()
	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("successfully completed")
}