package main

//Employee Employee
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func main() {

}

//执行用时：12 ms, 在所有 Go 提交中击败了96.75%的用户
//内存消耗：6.3 MB, 在所有 Go 提交中击败了50.00%的用户
func getImportance(employees []*Employee, id int) int {
	res := 0
	for _, e := range employees {
		if e.Id == id {
			res = e.Importance
			for j := range e.Subordinates {
				res += getImportance(employees, e.Subordinates[j])
			}
		}
	}
	return res
}
