package main

import "fmt"
import "time"
import "net/http"

func main() {
	fmt.Println("Running Host Ping")

	hosts := []PingHost { 
		PingHost {
			Url: "http://www.google.com",
			SuccessPorts: []int { 401, 403, 500 } ,
			Timeout: 5,
		},
		PingHost {
			Url: "http://www.nonexistentdomainatdotcom.com",
			SuccessPorts: []int { 401, 403, 500 } ,
			Timeout: 1,
		},
	}
	for _, h := range hosts {
		h.Ping()
	}
}

func (target PingHost) Ping() {
	success := false
	var timeout int
	if target.Timeout > 0 {
		timeout = target.Timeout
	} else {
		timeout = 3
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	fmt.Printf("Attempting to ping %v\n", target.Url)
	resp, err := client.Get(target.Url)

	if (err != nil || (intInSlice(resp.StatusCode, target.SuccessPorts) == false)) {
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			defer resp.Body.Close()
			fmt.Printf("Ping Fail: Expected %v status, received %v \n", target.SuccessPorts, resp.StatusCode)
		}
	} else {
		success = true
	}
	fmt.Printf("Success: %v\n\n", success)
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}