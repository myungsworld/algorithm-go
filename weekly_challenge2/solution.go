package weekly_challenge2

import "fmt"

// 대학 교수인 당신은, 상호평가를 통하여 학생들이 제출한 과제물에 학점을 부여하려고 합니다. 아래는 0번부터 4번까지 번호가 매겨진 5명의 학생들이 자신과 다른 학생의 과제를 평가한 점수표입니다.

// No.	0	1	2	3	4
// 0	100	90	98	88	65
// 1	50	45	99	85	77
// 2	47	88	95	80	67
// 3	61	57	100	80	65
// 4	24	90	94	75	65
// 평균	45.5	81.25	97.2	81.6	67.8
// 학점	F	B	A	B	D
// 위의 점수표에서, i행 j열의 값은 i번 학생이 평가한 j번 학생의 과제 점수입니다.

// 0번 학생이 평가한 점수는 0번 행에담긴 [100, 90, 98, 88, 65]입니다.
// 0번 학생은 자기 자신에게 100점, 1번 학생에게 90점, 2번 학생에게 98점, 3번 학생에게 88점, 4번 학생에게 65점을 부여했습니다.
// 2번 학생이 평가한 점수는 2번 행에담긴 [47, 88, 95, 80, 67]입니다.
// 2번 학생은 0번 학생에게 47점, 1번 학생에게 88점, 자기 자신에게 95점, 3번 학생에게 80점, 4번 학생에게 67점을 부여했습니다.
// 당신은 각 학생들이 받은 점수의 평균을 구하여, 기준에 따라 학점을 부여하려고 합니다.
// 만약, 학생들이 자기 자신을 평가한 점수가 유일한 최고점 또는 유일한 최저점이라면 그 점수는 제외하고 평균을 구합니다.

// 0번 학생이 받은 점수는 0번 열에 담긴 [100, 50, 47, 61, 24]입니다. 자기 자신을 평가한 100점은 자신이 받은 점수 중에서 유일한 최고점이므로, 평균을 구할 때 제외합니다.
// 0번 학생의 평균 점수는 (50+47+61+24) / 4 = 45.5입니다.
// 4번 학생이 받은 점수는 4번 열에 담긴 [65, 77, 67, 65, 65]입니다. 자기 자신을 평가한 65점은 자신이 받은 점수 중에서 최저점이지만 같은 점수가 2개 더 있으므로, 유일한 최저점이 아닙니다. 따라서, 평균을 구할 때 제외하지 않습니다.
// 4번 학생의 평균 점수는 (65+77+67+65+65) / 5 = 67.8입니다.
// 제외할 점수는 제외하고 평균을 구한 후, 아래 기준에 따라 학점을 부여합니다.

// 평균	학점
// 90점 이상	A
// 80점 이상 90점 미만	B
// 70점 이상 80점 미만	C
// 50점 이상 70점 미만	D
// 50점 미만	F
// 학생들의 점수가 담긴 정수형 2차원 배열 scores가 매개변수로 주어집니다. 이때, 학생들의 학점을 구하여 하나의 문자열로 만들어서 return 하도록 solution 함수를 완성해주세요.

func solution(scores [][]int) string {

	result := ""

	for i := 0; i < len(scores); i++ {
		cnt := 0
		tmpSlice := make([]int, 0)
		for j := 0; j < len(scores); j++ {
			tmpSlice = append(tmpSlice, scores[j][i])
		}
		fmt.Println("tmpSlice :", tmpSlice)
		fmt.Println("tmpSlice[i] = ", tmpSlice[i])
		max := Max(tmpSlice)
		if max == tmpSlice[i] {
			tmpSlice[i] = 0
			cnt++

		}
		min := Min(tmpSlice)
		if min == tmpSlice[i] {
			tmpSlice[i] = 0
			cnt++
		}
		fmt.Println("tmpSlice[i] = ", tmpSlice[i])
		fmt.Println("max:", max, "min:", min)

		sum := 0
		for k := 0; k < len(tmpSlice); k++ {
			sum += tmpSlice[k]
		}
		avg := 0

		fmt.Println("sum:", sum)

		if cnt == 0 {
			avg = sum / len(tmpSlice)
		} else {
			avg = sum / (len(tmpSlice) - 1)
		}

		if avg >= 90 {
			result += "A"
		} else if avg >= 80 && avg < 90 {
			result += "B"
		} else if avg >= 70 && avg < 80 {
			result += "C"
		} else if avg >= 50 && avg < 70 {
			result += "D"
		} else if avg < 50 {
			result += "F"
		}
		fmt.Println("avg : ", avg)
		fmt.Println("result : ", result)

	}

	return result

}

func Max(n []int) (a int) {

	a = n[0]
	for _, each := range n {

		if each > a {
			a = each
		}

	}

	cnt := 0
	for _, each := range n {
		if a == each {
			cnt++
		}
	}

	if cnt > 1 {
		return 0
	}
	return
}

func Min(n []int) (a int) {

	a = n[0]
	for _, each := range n {

		if each < a {
			a = each

		}

	}
	cnt := 0
	for _, each := range n {
		if a == each {
			cnt++
		}
	}

	if cnt > 1 {
		return 0
	}

	return
}
