package _020_kakao_blind_recruitment

import "strings"

func solution(idList []string, report []string, k int) []int {

	type reported struct {
		from string
		to   string
	}

	result := make([]int, len(idList))

	removeDupReport := map[reported]int{}

	for i := 0; i < len(report); i++ {
		arr := strings.Split(report[i], " ")
		from := arr[0]
		to := arr[1]
		r := reported{
			from: from,
			to:   to,
		}
		removeDupReport[r]++
	}

	reportToFrom := map[string][]string{}

	reportedCnt := map[string]int{}

	for key, _ := range removeDupReport {
		reportToFrom[key.to] = append(reportToFrom[key.to], key.from)
		reportedCnt[key.to]++
	}

	for reportedUser, value := range reportedCnt {
		if value >= k {

			for i := 0; i < len(idList); i++ {
				for _, reportUser := range reportToFrom[reportedUser] {
					if idList[i] == reportUser {
						result[i]++
					}
				}

			}

		}
	}

	return result
}

