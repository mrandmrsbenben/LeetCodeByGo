package main

import (
    "fmt"
    "sort"
)

type testcase struct {
    in []int
    out []int
}

func main() {
    cases := []testcase{{[]int{-4,-1,0,3,10},[]int{0,1,9,16,100}},
                    {[]int{-7,-3,2,3,11},[]int{4,9,9,49,121}}}
    for _,t := range cases{
        out := sortedSquares(t.in)
        for i := range out{
            if out[i] != t.out[i]{
                fmt.Printf("Fail.Expected:%v",t.out)
            }
        }
        fmt.Printf(" Output:%v\n",out)
    }
}

func sortedSquares(A []int) []int {
    for i,v := range A{
        A[i] = v * v
    }
    sort.Ints(A)
    return A
}
