package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

//[LeetCodeByGo: 更多LeetCode题库Go语言题解](https://github.com/mrandmrsbenben/LeetCodeByGo)
func main() {
	makeCommentsForMethod()
	// expect := []string{"bob,649,842,prague", "alex,175,1127,mexico", "iris,164,119,paris", "lee,991,1570,mexico", "lee,895,1876,taipei", "iris,716,754,moscow", "chalicefy,19,592,singapore", "chalicefy,820,71,newdelhi", "maybe,231,1790,paris", "lee,158,987,mexico", "iris,803,691,milan", "xnova,786,804,guangzhou", "lee,734,1915,prague", "bob,836,1904,dubai", "iris,666,231,chicago", "iris,677,1451,milan", "maybe,860,517,toronto", "iris,344,1452,bangkok", "lee,664,463,frankfurt", "chalicefy,95,1222,montreal", "lee,293,1102,istanbul", "maybe,874,36,hongkong", "maybe,457,1802,montreal", "xnova,535,270,munich", "iris,39,264,istanbul", "chalicefy,548,363,barcelona", "lee,373,184,munich", "xnova,405,957,mexico", "chalicefy,517,266,luxembourg", "iris,25,657,singapore", "bob,688,451,beijing", "bob,263,1258,tokyo", "xnova,852,330,barcelona", "xnova,589,837,budapest", "lee,152,981,mexico", "alex,893,1976,shenzhen", "xnova,560,825,prague", "iris,967,1119,guangzhou", "alex,924,223,milan", "chalicefy,212,1865,chicago", "alex,443,537,taipei", "bob,510,1923,madrid", "bob,798,343,hongkong", "iris,643,1703,madrid", "bob,478,928,barcelona", "maybe,75,1980,shanghai", "iris,176,268,milan", "maybe,560,587,milan", "alex,406,776,istanbul", "maybe,481,1504,munich", "maybe,685,602,madrid", "iris,678,788,madrid", "chalicefy,36,1984,paris", "iris,749,200,amsterdam", "iris,406,433,bangkok", "bob,777,542,taipei", "maybe,230,1434,barcelona", "iris,420,1818,zurich", "lee,622,194,amsterdam", "maybe,545,608,shanghai", "xnova,201,1375,madrid", "lee,432,520,dubai", "bob,150,1634,singapore", "maybe,467,1178,munich", "iris,45,904,beijing", "maybe,607,1953,tokyo", "maybe,636,558,milan", "bob,568,1674,toronto", "iris,825,484,madrid", "iris,951,930,dubai", "bob,465,1080,taipei", "chalicefy,16,176,rome", "xnova,836,153,jakarta", "bob,436,530,warsaw", "alex,354,1328,luxembourg", "iris,928,1565,paris", "xnova,627,834,budapest", "xnova,640,513,jakarta", "alex,119,16,toronto", "xnova,443,1687,taipei", "chalicefy,867,1520,montreal", "alex,456,889,newdelhi", "lee,166,3,madrid", "bob,65,1559,zurich", "maybe,668,572,mexico", "bob,402,922,montreal"}
	// output := []string{"alex,119,16,toronto", "alex,175,1127,mexico", "alex,354,1328,luxembourg", "alex,406,776,istanbul", "alex,443,537,taipei", "alex,456,889,newdelhi", "alex,893,1976,shenzhen", "alex,924,223,milan", "bob,65,1559,zurich", "bob,150,1634,singapore", "bob,263,1258,tokyo", "bob,402,922,montreal", "bob,436,530,warsaw", "bob,465,1080,taipei", "bob,478,928,barcelona", "bob,510,1923,madrid", "bob,568,1674,toronto", "bob,649,842,prague", "bob,688,451,beijing", "bob,777,542,taipei", "bob,798,343,hongkong", "bob,836,1904,dubai", "chalicefy,16,176,rome", "chalicefy,19,592,singapore", "chalicefy,36,1984,paris", "chalicefy,95,1222,montreal", "chalicefy,212,1865,chicago", "chalicefy,517,266,luxembourg", "chalicefy,548,363,barcelona", "chalicefy,820,71,newdelhi", "chalicefy,867,1520,montreal", "iris,25,657,singapore", "iris,39,264,istanbul", "iris,45,904,beijing", "iris,164,119,paris", "iris,176,268,milan", "iris,344,1452,bangkok", "iris,406,433,bangkok", "iris,420,1818,zurich", "iris,643,1703,madrid", "iris,666,231,chicago", "iris,677,1451,milan", "iris,678,788,madrid", "iris,716,754,moscow", "iris,749,200,amsterdam", "iris,803,691,milan", "iris,825,484,madrid", "iris,928,1565,paris", "iris,951,930,dubai", "iris,967,1119,guangzhou", "lee,158,987,mexico", "lee,166,3,madrid", "lee,293,1102,istanbul", "lee,373,184,munich", "lee,432,520,dubai", "lee,622,194,amsterdam", "lee,664,463,frankfurt", "lee,734,1915,prague", "lee,895,1876,taipei", "lee,991,1570,mexico", "maybe,75,1980,shanghai", "maybe,230,1434,barcelona", "maybe,231,1790,paris", "maybe,457,1802,montreal", "maybe,467,1178,munich", "maybe,481,1504,munich", "maybe,545,608,shanghai", "maybe,560,587,milan", "maybe,607,1953,tokyo", "maybe,636,558,milan", "maybe,668,572,mexico", "maybe,685,602,madrid", "maybe,860,517,toronto", "maybe,874,36,hongkong", "xnova,201,1375,madrid", "xnova,405,957,mexico", "xnova,443,1687,taipei", "xnova,535,270,munich", "xnova,560,825,prague", "xnova,589,837,budapest", "xnova,627,834,budapest", "xnova,640,513,jakarta", "xnova,786,804,guangzhou", "xnova,836,153,jakarta", "xnova,852,330,barcelona"}
	// findDifference(expect, output)
}

func makeCommentsForMethod() {
	text, err := clipboard.ReadAll()
	if err == nil {
		s1 := strings.Split(text, "\n")
		fmt.Println(len(s1))
		s2 := "//"
		for i := 0; i < len(s1); i++ {
			// fmt.Printf(s1[i])
			if i == 5 {
				s2 = s2 + "\n"
				s2 = s2 + "//"
			}
			s2 += s1[i]
		}
		fmt.Printf(s2)
		clipboard.WriteAll(s2)
	}
}

func findDifference(expect, output []string) {
	// expectStrs := strings.Fields(expect)
	// outputStrs := strings.Fields(output)
	diffMap := make(map[string]int)
	expectMap := make(map[string]int)
	for _, s := range expect {
		expectMap[s] = 0
	}
	outputMap := make(map[string]int)
	for _, s := range output {
		outputMap[s] = 0
		if _, ok := expectMap[s]; !ok {
			diffMap[s] = 0
		}
	}
	for _, s := range expect {
		if _, ok := outputMap[s]; !ok {
			diffMap[s] = 0
		}
	}
	fmt.Println("Begin *************")
	for s := range diffMap {
		fmt.Println(s)
	}
	fmt.Println("Finished **********")
}
