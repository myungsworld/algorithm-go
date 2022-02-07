package main

func main() {

	progresses := []int{95, 90, 99, 99, 80, 99}
	speeds := []int{1, 1, 1, 1, 1, 1}

	//fmt.Println(progresses[7:])

	solution(progresses, speeds)
}

func solution(progresses []int, speeds []int) []int {

	result := make([]int, 0)

	for len(progresses) > 0 {

		for i := 0; i < len(progresses); i++ {
			progresses[i] = progresses[i] + speeds[i]
		}

		deployedCnt := 0
		for i := 0; i < len(progresses); i++ {
			if progresses[i] >= 100 {
				deployedCnt++
			} else {
				break
			}
		}

		if deployedCnt > 0 {

			progresses = progresses[deployedCnt:]
			speeds = speeds[deployedCnt:]
			result = append(result, deployedCnt)
		}

	}

	return result
}
