package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var TEST_DATA string = "test.txt"

func main() {
	flagBytes := flag.Bool("c", false, "Number of Bytes")
	flagLines := flag.Bool("l", false, "Number of Lines")
	flagWords := flag.Bool("w", false, "Number of Words")
	flagChars := flag.Bool("m", false, "Number of Characters")
	flagDebug := flag.Bool("d", false, "Debug")
	flag.Parse()

	if *flagDebug {
		fmt.Println(`//   DEBUG   \\`)
		fmt.Println("Flag Bytes:", *flagBytes)
		fmt.Println("Flag Lines:", *flagLines)
		fmt.Println("Flag Words:", *flagWords)
		fmt.Println("Flag Chars:", *flagChars)
		fmt.Println(`\\ END DEBUG //`)
	}

	// Open file / read stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return
	}

	// Get file or stdin via pipe
	if len(flag.Arg(0)) > 0 {
		filepath := flag.Arg(0)
		fmt.Printf("File: %v\n", filepath)
		data, err := openFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		size := countBytesLen(data)
		_ = countBytes(size, filepath)
		return
	} else {
		log.Fatal("ERROR: No file or stdin data found")
	}
}

func openFile(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func countBytes(size int, filepath string) string {
	return fmt.Sprintf("%8d %s\n", size, filepath)
}

func countBytesLen(data []byte) int {
	return len(data)
}
