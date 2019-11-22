package main

import "fmt"
import "time"
import "net/http"

func main() {
	fmt.Println("Running Host Ping")

	host := &PingHost {
		Url: "http://www.google.com",
		SuccessPorts: []int { 401, 403, 500 } ,
		Timeout: 5,
	}
	host.Ping()
}

func (target PingHost) Ping() {
	var timeout int
	if target.Timeout > 0 {
		timeout = target.Timeout
	} else {
		timeout = 3
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	resp, err := client.Get(target.Url)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if intInSlice(resp.StatusCode, target.SuccessPorts) == false {
		fmt.Printf("Ping Fail: Expected %v status, received %v", target.SuccessPorts, resp.StatusCode)
	}
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}