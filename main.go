package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var env []string = []string{"dev", "development", "stage", "s3", "staging", "prod", "production", "test"}

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

	if len(osInput) == 0 {
		fmt.Println("[+] Provide a input")
		os.Exit(0)
	}
	file.Close()

	resp1, err1 := http.Get(fmt.Sprintf("http://s3.amazonaws.com/%s/", osInput[0]))
	if err1 != nil {
		return
	}
	if resp1.StatusCode == 200 {
		fmt.Println(fmt.Sprintf("Bucket found: "+"http://s3.amazonaws.com/%s/ [200]", osInput[0]))
	} else if resp1.StatusCode == 403 {
		fmt.Println(fmt.Sprintf("Bucket found: "+"http://s3.amazonaws.com/%s/ [403]", osInput[0]))
	}

	for i := 0; i < len(txtlines); i++ {
		resp, err := http.Get(fmt.Sprintf("http://s3.amazonaws.com/%s-%s/", txtlines[i], osInput[0]))
		if err != nil {
			return
		}
		if resp.StatusCode == 200 {
			fmt.Println(fmt.Sprintf("Bucket found: "+"http://s3.amazonaws.com/%s-%s/ [200]", txtlines[i], osInput[0]))
		} else if resp.StatusCode == 403 {
			fmt.Println(fmt.Sprintf("Bucket found: "+"http://s3.amazonaws.com/%s-%s/ [403]", txtlines[i], osInput[0]))
		}
	}
	for i := 0; i < len(env); i++ {
		resp2, err2 := http.Get(fmt.Sprintf("http://s3.amazonaws.com/%s-%s-%s/", env[i], osInput[0], txtlines[i]))
		if err2 != nil {
			return
		}
		if resp2.StatusCode == 200 {
			fmt.Println(fmt.Sprintf("http://s3.amazonaws.com/%s-%s-%s/ [200]", env[i], osInput[0], txtlines[i]))
		} else if resp2.StatusCode == 403 {
			fmt.Println(fmt.Sprintf("http://s3.amazonaws.com/%s-%s-%s/ [403]", env[i], osInput[0], txtlines[i]))
		}
	}
}
