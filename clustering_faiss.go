package faiss_cgo_kmeans

/*
#cgo CPPFLAGS: -Ithirdparty/libfaiss-src/c_api
#cgo CFLAGS: -Ithirdparty/libfaiss-src/c_api
#cgo darwin LDFLAGS: -Lthirdparty/runtimes/osx-arm64 -lfaiss_c -lfaiss -lomp
#cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
#cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#include <stdlib.h>
#include <Clustering_c.h>
#include <impl/AuxIndexStructures_c.h>
#include <index_factory_c.h>
#include <error_c.h>
*/
import "C"
import "errors"

// CGO code for https://github.com/facebookresearch/faiss/blob/main/c_api/Clustering_c.h functions
type Kmeans struct {
}

func New() *Kmeans {
	return &Kmeans{}
}

func (f *Kmeans) ComputeClusters(clusterCnt int64, data [][]float32) (centroids [][]float32, err error) {
	if len(data) == 0 {
		return nil, errors.New("empty rows")
	}
	if len(data[0]) == 0 {
		return nil, errors.New("zero dimensions")
	}

	rowCnt := int64(len(data))
	dims := int64(len(data[0]))

	// flatten data from 2D to 1D
	vectorFlat := make([]float32, dims*rowCnt)

	//TODO: optimize
	for r := int64(0); r < rowCnt; r++ {
		for c := int64(0); c < dims; c++ {
			vectorFlat[(r*dims)+c] = data[r][c]
		}
	}

	//TODO: do memory de-allocation if any
	centroidsFlat := make([]float32, dims*clusterCnt)
	var qError float32
	c := C.faiss_kmeans_clustering(
		C.ulong(dims),                 // d dimension of the data
		C.ulong(rowCnt),               // n nb of training vectors
		C.ulong(clusterCnt),           // k nb of output centroids
		(*C.float)(&vectorFlat[0]),    // x training set (size n * d)
		(*C.float)(&centroidsFlat[0]), // centroids output centroids (size k * d)
		(*C.float)(&qError),           // q_error final quantization error
		//@return error code
	)
	if c != 0 {
		return nil, getLastError()
	}

	if qError <= 0 {
		//final quantization error
		return nil, errors.New("final quantization error >0")
	}

	centroids = make([][]float32, clusterCnt)
	for r := int64(0); r < clusterCnt; r++ {
		centroids[r] = centroidsFlat[r*dims : (r+1)*dims]
	}
	return
}

func (f *Kmeans) Close() {
	//TODO implement me
	panic("implement me")
}

func getLastError() error {
	return errors.New(C.GoString(C.faiss_get_last_error()))
}
