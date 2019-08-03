package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cloud struct {
	val  int64
	next *Cloud
}

func main() {
	var cloudNum int64   // number of clouds
	var cloudList *Cloud // linked list
	var cloudHead *Cloud // head element of the linked list

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		cloudNum, _ = strconv.ParseInt(s, 10, 64)
	}

	// fill the list
	if scanner.Scan() {
		s := scanner.Text()
		args := strings.Fields(s)
		for i := 0; i < int(cloudNum); i++ {
			v, _ := strconv.ParseInt(args[i], 10, 64)
			if cloudList == nil {
				cloudList = &Cloud{v, nil}
				cloudHead = cloudList
			} else {
				cloudList.next = &Cloud{v, nil}
				cloudList = cloudList.next
			}
		}
	}

	jumps := 0
	cloud := cloudHead

	for {
		if cloud == nil {
			break
		}

		// get next cloud to see if we can jump
		next := cloud.next
		if next == nil {
			break
		}

		// if next cloud is thunder, jump to the one after it
		if next.val == 1 {
			cloud = next.next
			jumps += 1
			continue
		}

		nextnext := cloud.next.next

		// if next next cloud is not thunder, jump directly to it
		if nextnext != nil && nextnext.val == 0 {
			cloud = nextnext
		} else {
			cloud = next
		}

		jumps++
	}

	fmt.Println(jumps)
}
