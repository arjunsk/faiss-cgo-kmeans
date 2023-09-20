### Calling FAISS KMEANs using CGO


#### Note
- Build the FAISS library using 
```shell
cmake -DFAISS_ENABLE_GPU=OFF -DFAISS_ENABLE_PYTHON=OFF -DBUILD_TESTING=OFF -DCMAKE_BUILD_TYPE=Release -DFAISS_ENABLE_C_API=ON -DBUILD_SHARED_LIBS=OFF -B build .
```
It creates `libfaiss_c`.a file instead of `libfaiss_c.dylib`.

- Build script [here](/cgo/thirdparty/build-faiss-macos.sh)

- The build file is copied [here](/cgo/thirdparty/runtimes/osx-arm64/native)

- The Kmeans CGO code is [hear](/pkg/ivf/clustering_faiss.go)

- The test driver code is [here](/pkg/ivf/clustering_faiss_test.go)