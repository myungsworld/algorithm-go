package main

func main(){

	arr1 := [][]int{{1,4},{3,2},{4,1}}
	arr2 := [][]int{{3,2},{1,4}}

	solution(arr1,arr2)

}

// 행렬 곱 구하기
func solution(arr1 [][]int, arr2 [][]int) [][]int {


	result := make([][]int,0)

	for i := 0 ; i < len(arr1) ; i++ {
		tmp := make([]int,0)
		for j := 0 ; j < len(arr2[0]); j++ {
			tmp = append(tmp,0)
		}
		result = append(result,tmp)
	}

	for i := 0 ; i < len(arr1) ; i++ {
		for j :=0 ; j < len(arr2[0]) ; j++ {
			for k := 0 ; k < len(arr1[0]); k++ {
				result[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}

	return result
}

