package faiss_cgo_kmeans

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestFaissClustering_ComputeCenters(t *testing.T) {
	rowCnt := 3000
	dims := 5
	data := make([][]float32, rowCnt)
	loadData(rowCnt, dims, data)

	clusterCnt := 10
	var cluster Clustering = NewFaissClustering()
	centers, err := cluster.ComputeClusters(int64(clusterCnt), data)
	require.Nil(t, err)

	require.Equal(t, 10, len(centers))
}

func loadData(nb int, d int, xb [][]float32) {
	for r := 0; r < nb; r++ {
		xb[r] = make([]float32, d)
		for c := 0; c < d; c++ {
			xb[r][c] = rand.Float32() * 1000
		}
	}
}
