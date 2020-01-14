package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()

	for scanner.Scan() {
		cnt := 0
		t := scanner.Text()
		shouldTry := true
		for shouldTry {
			shouldTry = false
			tokyoIdx := strings.Index(t, "tokyo")
			kyotoIdx := strings.Index(t, "kyoto")

			//no tokyo, no kyoto
			if tokyoIdx == -1 && kyotoIdx == -1 {
				break
			}

			//tokyo, no kyoto
			if 0 <= tokyoIdx && kyotoIdx == -1 {
				t = t[tokyoIdx+5:]
				cnt++
				shouldTry = true
				continue
			}

			//no tokyo, kyoto
			if -1 == tokyoIdx && 0 < kyotoIdx {
				t = t[kyotoIdx+5:]
				cnt++
				shouldTry = true
				continue
			}

			//tokyo => kyoto
			if tokyoIdx < kyotoIdx {
				t = t[tokyoIdx+5:]
				cnt++
				shouldTry = true
				continue
			}

			//kyoto => tokyo
			if kyotoIdx < tokyoIdx {
				t = t[kyotoIdx+5:]
				cnt++
				shouldTry = true
				continue
			}
		}
		fmt.Println(cnt)
	}
}
