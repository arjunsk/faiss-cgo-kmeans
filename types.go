package faiss_cgo_kmeans

type Clustering interface {
	ComputeClusters(clusterCnt int64, data [][]float32) (centroids [][]float32, err error)
	Close()
}
