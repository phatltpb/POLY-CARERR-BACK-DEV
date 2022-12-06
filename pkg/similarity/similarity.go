package similarity

import (
	"errors"
	"math"

	"github.com/tuongnguyen1209/poly-career-back/apis/models"
	"golang.org/x/exp/slices"
)

type Similarity struct {
	Id   int
	Rank float32
}

func GetSimilarity(stuWatch []models.StudentWatch, idJob int, idStudent int) ([]Similarity, error) {

	listJob := []int{}
	listStudent := []int{}
	mapListJob := map[int][]int{}

	haveIdJob := false

	for _, v := range stuWatch {
		if !slices.Contains(listJob, v.JobId) {
			listJob = append(listJob, v.JobId)
		}
		if !slices.Contains(listStudent, v.StudentId) {
			listStudent = append(listStudent, v.StudentId)
		}
		if v.JobId == idJob {
			haveIdJob = true
		}
		if idStudent > 0 && v.StudentId == idStudent {
			idJob = v.JobId
		}
	}

	if !haveIdJob {
		return nil, errors.New("no sim")
	}

	for _, job := range listJob {
		newMap := []int{}

		for _, student := range listStudent {

			if index := slices.IndexFunc(stuWatch, func(e models.StudentWatch) bool {
				return e.StudentId == student && e.JobId == job
			}); index < 0 {
				newMap = append(newMap, 0)

			} else {
				newMap = append(newMap, int(stuWatch[index].Count))
			}
		}

		mapListJob[job] = newMap
	}

	listSim := []Similarity{}
	currentJob := mapListJob[idJob]
	for k, v := range mapListJob {

		if k == idJob {
			continue
		}
		var (
			num1 int = 0
			num2 int = 0
			num3 int = 0
		)

		for key, val := range v {
			num1 += val * currentJob[key]
			num2 += val * val
			num3 += currentJob[key] * currentJob[key]
		}

		x := math.Sqrt(float64(num2)) * math.Sqrt(float64(num3))
		var point float64 = float64(num1) / x

		listSim = append(listSim, Similarity{
			Id:   k,
			Rank: float32(1 - point),
		})
	}

	slices.SortFunc(listSim, func(a Similarity, b Similarity) bool {
		return a.Rank > b.Rank
	})

	return listSim, nil
}
