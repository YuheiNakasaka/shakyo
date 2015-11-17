package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// twitter envs
	envs := GetKeyVal()
	anaconda.SetConsumerKey(envs["MyTwitterConsumerKey"])
	anaconda.SetConsumerSecret(envs["MyTwitterConsumerSecret"])
	api := anaconda.NewTwitterApi(envs["MyTwitterAccessToken"], envs["MyTwitterAccessTokenSecret"])

	// async init
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	ch := make(chan string)

	// post query
	qs := []string{"golang", "ruby", "javascript", "elixir", "Perl", "java", "C++", "python", "swift", "objective-c", "php", "scala", "groovy"}
	for _, q := range qs {
		wg.Add(1)
		go func(q2 string, api2 *anaconda.TwitterApi) {
			defer wg.Done()
			res, _ := api2.GetSearch(q2, nil)
			for _, tweet := range res.Statuses {
				ch <- tweet.Text
			}
		}(q, api)
	}

	// ツイート受け取り
	go func() {
		for text := range ch {
			fmt.Println(text)
		}
	}()

	wg.Wait()

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func GetKeyVal() map[string]string {
	raw_envs := os.Environ()
	envs := make(map[string]string)
	strs := []string{}
	for _, val := range raw_envs {
		strs = strings.SplitN(val, "=", 2)
		envs[strs[0]] = strs[1]
	}
	return envs
}
