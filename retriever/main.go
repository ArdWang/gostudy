package main

import (
	"fmt"
	"github.com/gostudy/retriever/mook"
	"github.com/gostudy/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string] string) string
}


func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}

func post(poster Poster)  {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":"ccmouse",
		"course":"golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const  url = "http://www.imooc.com"

func session(s RetrieverPoster)  string{
	s.Post(url, map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mook.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}

	inspect(r)

	//fmt.Println(download(r))

	//fmt.Println(
		//download(mook.Retriever{Contents: "this is a fake imooc.com"}))

	inspect(r)

	// Type assertion
	//realRetriever := r.(*real.Retriever)
	//fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mook.Retriever); ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")

	fmt.Println(session(&retriever))

}

//
func inspect(r Retriever)  {
	fmt.Println("Inspecting",r)
	fmt.Printf(" > %T %v\n",r,r)
	fmt.Print(" > Type switch:")
	switch v :=r.(type) {
	case *mook.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
	fmt.Println()
}
