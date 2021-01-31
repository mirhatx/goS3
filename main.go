package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("common_bucket_prefixes.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	osInput := os.Args[1:]
	file.Close()

	for i := 0; i < len(txtlines); i++ {
		resp, err := http.Get(fmt.Sprintf("http://s3.amazonaws.com/%s-%s/", txtlines[i], osInput[0]))
		if err != nil {
			return
		}
		if resp.StatusCode == 200 {
			fmt.Println(fmt.Sprintf("Bucket found: "+"http://s3.amazonaws.com/%s-%s/", txtlines[i], osInput[0]))
		}
	}
}
