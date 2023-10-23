package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	totalmem "github.com/pbnjay/memory"
)

func main() {

	// Automatically set GOMAXPROCS to the number of your CPU cores.
	// Increase performance by allowing Golang to use multiple processors.
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs) // Sets the GOMAXPROCS value

	fmt.Println("------")
	fmt.Println("Available CPUs: " + fmt.Sprintf("%v", numCPUs))
	fmt.Println("Available OS Memory: " + fmt.Sprintf("%v MB | %v GB", (totalmem.TotalMemory()/Megabyte), (totalmem.TotalMemory()/Gigabyte)))
	fmt.Println("------")

	// support for openai
	rand.New(rand.NewSource(time.Now().UnixNano()))

	cliFlags()
	fmt.Println()

}
