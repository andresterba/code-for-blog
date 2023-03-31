package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
)

// Needed for profiling see https://pkg.go.dev/runtime/pprof
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

var charArray [62]string = [62]string{
	"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6",
	"7", "8", "9",
}

func algStringCasting() string {
	// Stolen from https://blog.pratimbhosale.com/building-a-url-shortener-using-go-and-sqlite#heading-shortening-the-url
	randomPart := ""
	for i := 0; i < 6; i++ {
		randomPart += string(rune(rand.Intn(26) + 97))
	}

	return randomPart
}

func algArrayLookup() string {
	// Stolen from https://www.geeksforgeeks.org/program-generate-random-alphabets/
	res := ""

	for i := 0; i < 6; i++ {
		res = res + charArray[rand.Intn(62)]
	}

	return res
}

func main() {
	// Needed for profiling see https://pkg.go.dev/runtime/pprof
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	algOne := algStringCasting()
	algThree := algArrayLookup()

	fmt.Println(algOne)
	fmt.Println(algThree)

	// Needed for profiling see https://pkg.go.dev/runtime/pprof
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
