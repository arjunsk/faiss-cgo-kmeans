package main

import (
	"faiss-go/pkg/ivf"
	"fmt"
	"math/rand"
)

func main() {
	rowCnt := 3000
	dims := 5
	data := make([][]float32, rowCnt)
	loadData(rowCnt, dims, data)

	clusterCnt := 10
	var cluster = ivf.NewFaissClustering()
	centers, err := cluster.ComputeClusters(int64(clusterCnt), data)
	if err != nil {
		panic(err)
	}

	fmt.Println(centers)
}

func loadData(nb int, d int, xb [][]float32) {
	for r := 0; r < nb; r++ {
		xb[r] = make([]float32, d)
		for c := 0; c < d; c++ {
			xb[r][c] = rand.Float32() * 1000
		}
	}
}
