package sparse

import (
	"math"
//	"time"
//	"log"
)

type Sparse struct {
	m map[int]float64
	n float64
}

func Make(m map[int]float64) Sparse {
	return Sparse{m, 0}
}

func (s Sparse) Len() int { return len(s.m) }

func (s Sparse) Add(index int, value float64) {
	if value == 0 {
		return
	}
	s.m[index] = value
}

func (s Sparse) Get(i int) float64 {
	if r, ok := s.m[i]; ok {
		return r
	}
	return 0
}

func (s Sparse) Cos(t Sparse) float64 {
//     	start := time.Now()
	dot := s.dot(t)
//	log.Println("d: ", time.Since(start))
//	start = time.Now()
	snorm := s.norm()
//	log.Println("s: ", time.Since(start))
//	start = time.Now()
	tnorm := t.norm()
//	log.Println("t: ", time.Since(start))
	return 1 - (dot / (snorm * tnorm))
}


func (s Sparse) dot(t Sparse) float64 {
	if s.Len() <= t.Len() {
		var sum float64 = 0
		for k, v := range s.m {
			if val, ok := t.m[k]; ok {
				sum += v * val
			}
		}
		return sum
	}
	return t.dot(s)
}

func (s Sparse) norm() float64 {
	if s.n != 0 {
		return s.n
	} else {
		return s.getNorm()
	}
}

func (s Sparse) getNorm() float64 {
	var sum float64 = 0
	for _, v := range s.m {
		sum += v * v
	}
	norm := math.Sqrt(sum) 
	s.n = norm
	return norm
}

func CosMatrix(matrix []Sparse, vector Sparse) []float64 {
	var r []float64
	for i, v := range matrix {
		r[i] = v.Cos(vector)
	}
	return r
}

func NormMatrix(m []Sparse) {
     for _, w := range m {
     	 _ = w.getNorm()
     }
}

