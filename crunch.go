package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/aep/feistel"
	"github.com/thanhpk/randstr"
)

func main() {
	const termSize = 99
	var max uint32 = termSize * termSize * 4

	for i := range max {
		num := feistel.Map(i, max, randstr.String(10))
		op := num % 4
		num /= 4

		term2 := (num % termSize) + 1
		term1 := (num / termSize) + 1

		var opStr string
		switch op {
			case 0:
				opStr = "+"
			case 1:
				opStr = "-"
			case 2:
				opStr = "*"
			case 3:
				opStr = "/"
		}

		fmt.Printf("%d %s %d = ", term1, opStr, term2)

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			num1 := float32(term1)
			num2 := float32(term2)

			var answer float32
			switch op {
				case 0:
					answer = num1 + num2
				case 1:
					answer = num1 - num2
				case 2:
					answer = num1 * num2
				case 3:
					answer = num1 / num2
			}

			var correct bool = false
			input, err := strconv.ParseFloat(scanner.Text(), 32)
			if err == nil {
				correct = float32(input) == answer
			}

			fmt.Println(correct)
			fmt.Println(answer)
		}

		if err := scanner.Err(); err != nil {
			// todo: format error
			log.Fatalln(err)
		}
	}
}
