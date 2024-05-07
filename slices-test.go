package main

import "fmt"

func main() {
	months := [...]string{0: "oct", 1: "jan", 2: "feb", 3: "mar", 4: "apr", 5: "dec"}
	days := [...]string{"1", "2", "3", "4", "5", "6"}
	fmt.Println("months:", months, "cap:", cap(months))
	fmt.Println("days:", days)

	s1 := months[2:6]
	fmt.Println("s1:", s1, s1[0], "cap:", cap(s1))
	s1[0] = "FEB"
	fmt.Println("s1:", s1, s1[0], "cap:", cap(s1))
	fmt.Println("months:", months)
	s1 = append(s1, "nig")
	fmt.Println("s1:", s1, "cap:", cap(s1))
	fmt.Println("months:", months, "cap:", cap(months))
	s1 += days[2:5]
	fmt.Println("s1:", s1, "cap:", cap(s1))
	s1[len(s1)-2] = "new"
	fmt.Println("s1:", s1, "cap:", cap(s1))
	fmt.Println("days:", days, "cap:", cap(days))
	s1 = append(s1, "some")
	fmt.Println("s1:", s1, "cap:", cap(s1))
	fmt.Println("days:", days, "cap:", cap(days))
}
