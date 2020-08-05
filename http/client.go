package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, err := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)

	//request.Header.Add(
	//	"User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:",req)
			return nil
		},
	}

	resp, err := client.Do(request)

	//resp, error := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, error := httputil.DumpResponse(resp, true)

	if error != nil {
		panic(error)
	}

	fmt.Printf("%s/n",s)

}
