package model

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

// @Author linjiabao
// @Date   2022/5/23

func TestBalance(t *testing.T) {
	work := make(chan Request)
	for i := 0; i < 100; i++ {
		go Requester(work)
	}
	NewBalancer().Balance(work)
}

func TestPipelines(t *testing.T) {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := MD5All("/Users/krysta/go")
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}

func TestSubscribe(t *testing.T) {
	// STARTMAIN OMIT

	// STARTMERGECALL OMIT
	// Subscribe to some feeds, and create a merged update stream.
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")))
	// STOPMERGECALL OMIT

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	// Uncomment the panic below to dump the stack traces.  This
	// will show several stacks for persistent HTTP connections
	// created by the real RSS client.  To clean these up, we'll
	// need to extend Fetcher with a Close method and plumb this
	// through the RSS client implementation.
	//
	// panic("show me the stacks")

	// STOPMAIN OMIT
}
