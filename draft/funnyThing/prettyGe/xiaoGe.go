package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	// MYWORD My word
	var head string
	var tail string
	var MYWORD string
	var sep string
	var zoom float64
	flag.StringVar(&head, "head", "浮世万千，吾爱有三，日、月和卿，日为朝，月为暮，卿为朝朝暮暮!", "A sentence printed on the head")
	flag.StringVar(&tail, "tail", "\t\t\t\t--- pretty XiaoHua", "A sentence printed on the tail")
	flag.StringVar(&MYWORD, "words", "I love three things in this world, sun moon and you! - Sun for morning, moon for night and you forever!", "The words you want to talk")
	flag.StringVar(&sep, "sep", "-", "The separator")
	flag.Float64Var(&zoom, "zoom", 1.0, "Zoom setting")
	flag.Parse()

	chars := strings.Split(MYWORD, sep)

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(head)
	fmt.Println()
	time.Sleep(time.Duration(1) * time.Second)
	for _, char := range chars {
		allChar := make([]string, 0)

		for y := 12 * zoom; y > -12*zoom; y-- {
			lst := make([]string, 0)
			lstCon := ""
			for x := -30 * zoom; x < 30*zoom; x++ {
				x2 := float64(x)
				y2 := float64(y)
				formula := math.Pow(math.Pow(x2*0.04/zoom, 2)+math.Pow(y2*0.1/zoom, 2)-1, 3) - math.Pow(x2*0.04/zoom, 2)*math.Pow(y2*0.1/zoom, 3)
				if formula <= 0 {
					index := int(x) % len(char)
					if index >= 0 {
						lstCon += string(char[index])
					} else {
						lstCon += string(char[int(float64(len(char))-math.Abs(float64(index)))])
					}

				} else {
					lstCon += " "
				}
			}
			lst = append(lst, lstCon)
			allChar = append(allChar, lst...)
		}

		for _, text := range allChar {
			fmt.Printf("%s\n", text)
		}
	}
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("\t\t\t\t", tail)
}
