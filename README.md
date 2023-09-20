### Calling FAISS KMEANs using CGO


#### Note
- Build the FAISS library using 
```shell
cmake -DCMAKE_BUILD_TYPE=Release -DFAISS_ENABLE_C_API=ON -DBUILD_SHARED_LIBS=OFF -B build .
```
It creates `libfaiss_c.a` file instead of `libfaiss_c.dylib`.

- Build script [here](/cgo/thirdparty/build-faiss-macos.sh)

- The build file is copied [here](/cgo/thirdparty/runtimes/osx-arm64/native)

- The Kmeans CGO code is [hear](/pkg/ivf/clustering_faiss.go)

- The test driver code is [here](/pkg/ivf/clustering_faiss_test.go)


### Instructions

1. To build the `libfaiss_c.a` in MacOS
```shell
cd cgo/thirdparty
sh build-faiss-macos.sh
```

2. To run the test
```shell
cd pkg/ivf
go test -v -run TestFaissKmeans
```