## In-Memory Cache with freecache: A Playground!
---
### Library used: [freecache](https://github.com/coocood/freecache)
---
### What Is This?
A simple Go application for writing key/value pairs to an in-memory cache and setting them to expire after 1 second.

Reasons for tinkering with this solution: 
- It's 100% Go!
- The example provided in the project's readme gave me a great jumping off point
- Built-in functions were pretty clearly defined, which was great because I'm still getting used to Go
- It doesn't have any garbage collection overhead
- Seems to have good performance reviews
- I want to learn how to use in-memory caching with Go for some very special projects :)

### How to Build This
- NOTE: This project requires Go v1.17
1. Fork repo and clone to your machine's desired working directory
2. In that working directory, `cd` into `in-memory-cache`. You should see three files:
- go.mod (module defined with package requirements shown)
- go.sum (checksums from required packages)
- main.go (where all the in-cache memory magic happens!)
3. Retrieve the required packages with `go get github.com/coocood/freecache`
4. Confirm checksums with `go mod tidy`
5. Run `main.go` with `go run main.go`

Good feedback with the default arguments looks like:
```
Key/value pair was successfully entered into cache...

In-memory cache entry count:  1
Sleep for 1 seconds...
Entry not found -- the key expired! YAY!
There was no key to delete! This is the desired outcome! 

In-memory cache entry count:  0
```

### What This Is Doing
- Imports the following libraries:
   - `fmt` for printing lines to the console
   - `time` for implementing a sleep function (which tests the key expiration in the cache)
   - `github.com/coocood/freecache`, which is freecache itself. [Docs](https://github.com/coocood/freecache#readme)
- Sets the cache maximum size to 100MB
- Instantiates a new in-memory cache of that size
- Creates a key/value pair for testing
- Sets the in-memory expiration to 1 second for that test key (with error handling)
- Confirm that the key is in the cache (In-memory cache entry count should be 1)
- Sleeps for 1 second to allow the key to expire
- Try to retrieve the key, with success and failure criterion defined and with error handling
- Try to delete the key, with error handling. No key to delete? Huzzah!
- Confirm that the key is NOT in the cache (In-memory cache entry count should be 0)

### What's Next?
- Consume a message queue and stream keys into this cache. From there, they could be cached or destroyed based on set expiration.
- Use the library's server metrics to observe performance metrics
- Baseline based on generic projects from there 
