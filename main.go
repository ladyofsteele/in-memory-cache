package main

import (
	"fmt"
	"time"
	"github.com/coocood/freecache"
)

// Decided to proceed with freecache for the sake of trying out an in-memory caching solution
// that is reported to have excellent performance and low-impact on resources. I wanted to use
// an in-memory solution so that I could use the Horizontal Pod Autobalancer to scale up
// after a certain amount of memory was consumed by the in-memory cache, that coincided
// with having 10 stored keys in-memory

func main(){
	// In bytes, where 1024 * 1024 represents a single Megabyte, and 100 * 1024*1024 represents 100 Megabytes.
	cacheSize := 100 * 1024 * 1024

	// Instantiate a new cache of that size
	cache := freecache.NewCache(cacheSize)

	// Create a test key/value pair just to see how this thing works
	key := []byte("abc")
	val := []byte("def")

	// Set that key/value pair to expire in 1 second and enter it into cache
	expire := 1 // change this value to test my error statements!
	set := cache.Set(key, val, expire)
	if set != nil {
		fmt.Println("\nKey/value pair was not successfully entered into cache. Error: ", set)
	} else {
		fmt.Println("\nKey/value pair was successfully entered into cache...")
	}

	// Prove that the key was entered into the cache
	fmt.Println("\nIn-memory cache entry count: ", cache.EntryCount())

	// Calling Sleep method to see if the expiry works
	fmt.Println("Sleep for 1 seconds to allow the key to expire...")
	time.Sleep(1 * time.Second)

	// Prove that the key cannot be retrieved because it has expired.
	// Success criteria: return an error if this key expires in the allotted time.
	// Failure criteria: return the key, :sad_panda_emoji:
	got, err := cache.Get(key)
	if err != nil {
		fmt.Println(err, "-- the key expired! YAY!")
	} else {
		fmt.Printf("Oh no! We can still retrieve the key! Value for key: %s\n", got)
	}

	// This next block will verify that we cannot delete the selected key
	// because ideally it will have expired. Else, it'll clean up our cache
	// so we can continue to troubleshoot.
	affected := cache.Del(key)
	if affected {
		fmt.Println("Sad day -- the key didn't expire. They key has been deleted. Keep trying!\n")
	} else {
		fmt.Println("There was no key to delete! This is the desired outcome! \n")
	}
	fmt.Println("In-memory cache entry count: ", cache.EntryCount()) // Should show 0 if the keys are expiring
}