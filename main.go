package main

import (
	"fmt"
	"github.com/temoto/robotstxt"
	// "golang.org/x/net/context"
	// "github.com/thingful/thingfulx"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	// checkThisUrl("https://www.car2go.com/api/kvfdjflk/cdsd.txt")

	urls := []string{
		"https://www.car2go.com/api/kvfdjflk/cdsd.txt",
		"https://www.car2go.com/api1/kvfdjflk/cdsd.txt",
	}
	checkURLs(urls)
}

func checkURLs(urls []string) (bool, error) {

	robotsAddress := ""
	allAllowed := true
	robots, err := robotstxt.FromString("User-agent: *\nDisallow:")
	if err != nil {
		return false, err
	}
	// for every urls, check if it's the same robot, if not request this specific robot

	for _, u1 := range urls {

		u, err := url.Parse(u1)
		if err != nil {
			return false, err
		}

		newRobotsAddress := u.Scheme + "://" + u.Host + "/robots.txt"

		if newRobotsAddress != robotsAddress {

			// fmt.Println("this is new robot")
			robotsAddress = newRobotsAddress
			resp, err := http.Get(robotsAddress)
			if err != nil {
				return false, err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			// fmt.Printf("body = %s\n", body)
			robots, err = robotstxt.FromBytes(body)

		}

		allow := robots.TestAgent(u.Path, "thingful")
		if !allow {
			fmt.Printf("%s is NOT allowed\n", u.Path)
			allAllowed = false
		}

	}

	return allAllowed, nil

}

func checkThisUrl(url1 string) {

	u, err := url.Parse(url1)
	if err != nil {
		fmt.Println("something")
	}

	// fmt.Println(u.Scheme)
	// fmt.Println(u.Opaque)
	robotsAddress := u.Scheme + "://" + u.Host + "/robots.txt"
	fmt.Printf("robotsAddress = %s\n", robotsAddress)
	fmt.Printf("path to check  = %s\n", u.Path)

	resp, err := http.Get(robotsAddress)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// fmt.Printf("body = %s\n", body)
	robots, err := robotstxt.FromBytes(body)

	allow := robots.TestAgent(u.Path, "thingful")
	if allow {
		fmt.Println("this path is allowed")
	} else {
		fmt.Println("this path is NOT allowed")
	}

}
