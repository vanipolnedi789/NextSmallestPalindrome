package main

import (
    "fmt"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup
var inputCount = 1000

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	startTime := time.Now()
	inputChannel := make(chan int)
	for i := 0 ; i < inputCount ; i ++ {
		wg.Add(1)
		go fetchRandomInput(inputCount,inputChannel)
		go func() {
			randomValue := <- inputChannel
			fmt.Println("next smallest palindrome after ",randomValue," is: ",nextSmallestPalindrome(randomValue))
			defer wg.Done()
		}()    	
	}
	// blocks/waits for waitgroup
    wg.Wait()
	endTime := time.Now()
	fmt.Println("Time taken for palindrome process : ",endTime.Sub(startTime))
}

//fetchRandomInput - fetches random value between 1 to 1000
func fetchRandomInput(max int ,inputChannel chan int) {
	inputChannel <- rand.Intn(max-1) + 1
}

//nextSmallestPalindrome -- finds next smallest palindrome after given input
func nextSmallestPalindrome(num int) int{
    for i:=num+1;i<num+15;i++ {
        if fmt.Sprint(i) == stringReverse(fmt.Sprint(i))  {
			return i
        }
    }   
	return 0 
}

//stringReverse -- returns the reverse string for the given string
func stringReverse(stringChannel string)  string{
    rns := []rune(stringChannel) 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
	return string(rns)
}
