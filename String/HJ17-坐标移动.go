package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hj17() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	x, y := removeCoord(s)
	fmt.Printf("%v,%v", x, y)
}

func removeCoord(s string) (int, int) {
	x, y := 0, 0
	s1 := strings.Split(s, ";")
	for i := 0; i < len(s1); i++ {
		s2 := s1[i]
		if len(s2) >= 2 {
			first := s2[0]
			if first == 'A' || first == 'D' || first == 'S' || first == 'W' {
				s2 = s2[1:]
				num, err := strconv.Atoi(s2)
				if err == nil {
					switch first {
					case 'A':
						x -= num
					case 'S':
						y -= num
					case 'W':
						y += num
					case 'D':
						x += num
					}
				}
			}
		}
		continue
	}
	return x, y
}
