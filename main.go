package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Fedy1507/18-uyga-vazifa/faktorial"
)

func main() {

	var check int
	fmt.Printf("1 Faktorial \n")
	fmt.Printf("2 Fayl \n")
	fmt.Printf("Qaysi birini tanlaysiz? ->")
	fmt.Scan(&check)
	if check == 1 {
		c := make(chan int)
		var n int
		fmt.Printf("qaysi sonning faktorialini hisoblashni hohlaysiz? ")
		fmt.Scan(&n)
		go faktorial.Faktorial(n, c)

		fmt.Printf("%d ning faktoriali: %d\n", n, <-c) // natijani channel dan olib konsolga chiqaramiz
	} else if check == 2 {
		files := []string{"fayl/file1.txt", "fayl/file2.txt", "fayl/file3.txt"}

		results := make(chan string)

		var wg sync.WaitGroup

		for _, file := range files {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()

				content, err := os.ReadFile(file)
				if err != nil {
					fmt.Println(err)
					return
				}

				results <- string(content)
			}(file)
		}

		outputFile, err := os.Create("combined_results.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outputFile.Close()

		for i := 0; i < len(files); i++ {
			result := <-results
			outputFile.WriteString(result)
		}

		close(results)

		wg.Wait()
	}

}
