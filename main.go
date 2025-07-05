package main

import (
	"fmt"
	"gen_pass/password"
	"sync"
)


func main(){
	
	// Generate single password
	pass, err := password.GenPass(true, true, true, 20)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%v, %d \n", pass, len(pass))

	// Generating multiple passwords
	numberOfPasswords := 100
	pwChan := make(chan string)

	var wg sync.WaitGroup

	// Launch a multiple go-routine to generate passwords
	for range numberOfPasswords {
		wg.Add(1)
		go func(){
			defer wg.Done()
			pass, _ := password.GenPass(true, true, true, 20)
			// send the return down the channel
			pwChan <- pass
		}()
	}

	// Launch a goroutine to clean up after all other goroutines finish
	go func(){
		wg.Wait()
		close(pwChan)
	}()

	// Read from channel and append to list of passwords
	pwList := make([]string, 0, numberOfPasswords)
	for pw := range pwChan {
		pwList = append(pwList, pw)
	}

	fmt.Println(pwList)
	fmt.Println(len(pwList))
}