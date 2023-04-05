package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"time"
)

// type P struct {
// 	Name string
// 	Age  int
// }

// iota

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

// context

func longprocess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	// time

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// regex

	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	// Sort

	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	// p := P{"Nancy", 20}
	fmt.Println(i, s, p)

	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

	// iota

	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	// context

	// ch := make(chan string)
	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()
	// 	ctx := context.TODO()
	// 	go longprocess(ctx, ch)

	// CTXLOOP:
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println(ctx.Err())
	// 			break CTXLOOP
	// 		case <-ch:
	// 			fmt.Println("success")
	// 			break CTXLOOP
	// 		}
	// 	}
	// 	fmt.Println("###########")

	// ioutil

	// content, err := ioutil.ReadFile("main.go")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(content))

	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }

	rr := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(rr)
	fmt.Println(string(content))
}
