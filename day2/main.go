package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(input string) int {
	
	idRanges := strings.Split(input, ",");
	answer := 0;

	for i := 0; i < len(idRanges); i++ {
		idRanges[i] = strings.TrimSpace(idRanges[i]);
		idRange := strings.Split(idRanges[i], "-");
		fmt.Println(idRanges[i]);

		lower,_ := strconv.Atoi(idRange[0]);
		higher,_ := strconv.Atoi(idRange[1]);

		fmt.Println(lower, "_", higher);

		for j := lower; j <= higher; j++ {
			if repeating(j) {
				//fmt.Println("Repeating:", j, "in range", idRanges[i]);
				answer += j;
			}
		} 
	}

	return answer;
}

func repeating(check int) bool {

	//Part 2 Solution
	checkAsAString := strconv.Itoa(check);
	s := checkAsAString + checkAsAString;
	r := s[1:len(s)-1];

	return strings.Contains(r, checkAsAString);

	//Part 1 Solution
	// if strings.Index(strconv.Itoa(check), "0") == 0 {
	// 	return true;
	// }

	// if len(strconv.Itoa(check)) % 2 != 0 {
	// 	return false;
	// }

	// s := strconv.Itoa(check);

	// lowerHalf := s[:len(s)/2];
	// higherHalf := s[len(s)/2:];

	// return lowerHalf == higherHalf;

}

// func findFactors(n int) []int {
// 	if n <= 0 {
// 		return []int{}
// 	}
// 	factors := make([]int, 0)
// 	// Iterate from 1 up to the square root of n
// 	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
// 		if n%i == 0 {
// 			// If i is a factor, add it
// 			factors = append(factors, i)
// 			// Also add the other factor (n/i), unless it's the same as i (for perfect squares)
// 			if i != n/i {
// 				factors = append(factors, n/i)
// 			}
// 		}
// 	}
// 	// Sort the factors for a cleaner output
// 	sort.Ints(factors)
// 	return factors
// }


func fileInput() string {
	file,error := os.Open("input.txt");

	if error != nil {
		panic("Wrong file path");
	}

	input := bufio.NewReader(file);

	fileContents,err := input.ReadString('\n')

	if err != nil {
		panic("Something wrong with the file");
	}

	return fileContents;
}

func main() {
	input := fileInput();

	fmt.Println(solution(input))
}