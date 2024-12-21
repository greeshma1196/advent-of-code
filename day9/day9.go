// package day9

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// )

package day9

import (
	"bufio"
	"fmt"
	"os"
)

func ComputeAOCDay9_1(name string) {

	input := readFile(name)

	// two pointer solution
	length := len(input)
	left := 0
	right := length - 1

	for {
		for left < length && input[left] != -1 {
			left += 1
		}

		for right >= 0 && input[right] == -1 {
			right -= 1
		}

		if left >= right {
			break
		}

		input[left], input[right] = input[right], input[left]
		left += 1
		right -= 1
	}

	res := 0

	for i, s := range input {
		if s == -1 {
			continue
		}
		res += i * s
	}

	fmt.Printf("Result of day 9 part 1: %d\n", res)
}

func ComputeAOCDay9_2(name string) {
	input := readFile(name)

	length := len(input)

	occurances := make(map[int]int)

	for _, n := range input {
		if n == -1 {
			continue
		}
		occurances[n] += 1
	}

	fileIDList := make([]int, 0)
	fileID := -1
	for i := length - 1; i >= 0; i-- {
		if input[i] != -1 && input[i] != fileID {
			fileID = input[i]
			fileIDList = append(fileIDList, fileID)
		}
	}

	startFileIdx := make(map[int]int)
	fileID = -1
	for i := 0; i < length; i++ {
		if input[i] != -1 && input[i] != fileID {
			fileID = input[i]
			startFileIdx[fileID] = i
		}
	}

	for _, fileID := range fileIDList {
		fileOcc := occurances[fileID]
		idx := startFileIdx[fileID]
		freeOcc := 0

		if idx == 0 {
			break
		}

		i := 0
		for i < idx {
			if input[i] == -1 {
				freeOcc = 1
				j := i + 1
				for j < idx {
					if input[j] != -1 {
						break
					}
					freeOcc += 1
					j += 1
				}

				if fileOcc <= freeOcc {
					j = i
					count := 0
					for count < fileOcc {
						input[j] = fileID
						j += 1
						count += 1
					}

					count = 0
					for idx < length && count < fileOcc {
						input[idx] = -1
						idx += 1
						count += 1
					}
					break
				}
			}
			i += 1
		}

	}

	res := 0

	for i, s := range input {
		if s == -1 {
			continue
		}
		res += i * s
	}

	fmt.Printf("Result of day 9 part 2: %d\n", res)

}

func readFile(name string) []int {
	fi, err := os.Open(name)
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	data := make([]int, 0)

	// idNumb := 0
	// flag := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			digit := (line[i] - '0')
			data = append(data, int(digit))
			// data = populateData(idNumb, int(digit), data, flag)
			// flag++
			// if flag%2 == 0 {
			// 	idNumb++
			// }
		}
	}
	data = parseData(data)
	return data
}

func parseData(input []int) []int {
	numID := 0
	data := make([]int, 0)
	for i := 0; i < len(input); i++ {
		n := input[i]
		j := 0
		if i%2 == 0 {
			for j < n {
				data = append(data, numID)
				j += 1
			}
			numID += 1
		} else {
			for j < n {
				data = append(data, -1)
				j += 1
			}
		}
	}
	return data
}
