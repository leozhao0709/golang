package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
)

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	var arr []int

	source := rand.NewSource(666)
	rand := rand.New(source)

	cpuProfile, _ := os.Create("./debug/profile/cpu.prof")   // ignore error here
	memProfile1, _ := os.Create("./debug/profile/mem1.prof") // ignore error here

	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()
	pprof.WriteHeapProfile(memProfile1)

	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	bubbleSort(arr)

	memProfile2, _ := os.Create("./debug/profile/mem2.prof") // ignore error here
	pprof.WriteHeapProfile(memProfile2)
	defer memProfile2.Close()
}
