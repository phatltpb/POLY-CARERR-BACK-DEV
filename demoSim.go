package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/slices"
)

type watch struct {
	id        int
	StudentId int
	JobId     int
	Count     int
}

type result struct {
	id   int
	rank float64
}

var listWatch []watch = []watch{
	{id: 1, JobId: 4, StudentId: 4, Count: 4},
	{id: 2, JobId: 2, StudentId: 4, Count: 3},
	{id: 3, JobId: 4, StudentId: 5, Count: 1},
	{id: 4, JobId: 3, StudentId: 6, Count: 4},
	{id: 5, JobId: 7, StudentId: 3, Count: 7},
	{id: 6, JobId: 2, StudentId: 3, Count: 2},
	{id: 7, JobId: 3, StudentId: 3, Count: 2},
	{id: 8, JobId: 5, StudentId: 7, Count: 9},
	{id: 9, JobId: 3, StudentId: 7, Count: 12},
	{id: 10, JobId: 2, StudentId: 8, Count: 2},
	{id: 11, JobId: 4, StudentId: 8, Count: 5},
	{id: 12, JobId: 2, StudentId: 9, Count: 8},
}

func main1() {
	find := 3

	listJob := []int{}
	listStudent := []int{}

	for _, v := range listWatch {
		if !slices.Contains(listJob, v.JobId) {
			listJob = append(listJob, v.JobId)
		}
		if !slices.Contains(listStudent, v.StudentId) {
			listStudent = append(listStudent, v.StudentId)
		}

	}

	fmt.Printf("%+v \n", listJob)
	fmt.Printf("%+v \n", listStudent)

	mapArr := map[int][]int{}

	for _, job := range listJob {
		newMap := []int{}

		for _, student := range listStudent {

			if index := slices.IndexFunc(listWatch, func(e watch) bool {
				return e.StudentId == student && e.JobId == job
			}); index < 0 {
				newMap = append(newMap, 0)

			} else {
				newMap = append(newMap, listWatch[index].Count)
			}
		}

		mapArr[job] = newMap
	}

	fmt.Printf("%+v \n", mapArr)

	sug2 := []result{}
	jobSlect := mapArr[find]

	for k, v := range mapArr {

		if k == find {
			continue
		}
		var (
			num1 int = 0
			num2 int = 0
			num3 int = 0
		)

		for key, val := range v {
			num1 += val * jobSlect[key]
			num2 += val * val
			num3 += jobSlect[key] * jobSlect[key]
		}

		x := math.Sqrt(float64(num2)) * math.Sqrt(float64(num3))
		var point float64 = float64(num1) / x

		a := 1 - point
		sug2 = append(sug2, result{
			id:   k,
			rank: a,
		})
	}

	slices.SortFunc(sug2, func(a result, b result) bool {
		return a.rank > b.rank
	})

	fmt.Printf("%+v \n", sug2)

}
