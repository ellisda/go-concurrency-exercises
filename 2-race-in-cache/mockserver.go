//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"strconv"
	"sync"
)

const (
	cycles        = 3
	callsPerCycle = 100
)

// RunMockServer simulates a running server, which accesses the
// key-value database through our cache
func RunMockServer(cache *KeyStoreCache) {
	var wg sync.WaitGroup

	for c := 0; c < cycles; c++ {
		wg.Add(1)
		go func(c int) {
			for i := c * callsPerCycle; i < (c+1)*callsPerCycle; i++ {

				cache.Get("Test" + strconv.Itoa(i))

			}
			wg.Done()
		}(c)
	}

	wg.Wait()
}
