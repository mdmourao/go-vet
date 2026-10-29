package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"name"`
}

func main() {
	nome := "Gophers"
	fmt.Printf("Olá %d\n", nome)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	<-ctx.Done()

	resp, err := http.Get("https://gophers.pt")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	_, err = os.Open("exemplo.txt")
	var pathErr *fs.PathError
	if errors.As(err, pathErr) {
		fmt.Println(pathErr.Path)
	}
}
