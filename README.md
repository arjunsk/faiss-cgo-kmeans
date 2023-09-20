## Calling FAISS KMEANs using CGO

### Instructions

1. To build the `libfaiss_c.a` in MacOS
```shell
cd cgo/thirdparty
sh build-faiss-macos.sh
```
2. Build go project
```shell
CGO_LDFLAGS="-L/cgo/thirdparty/runtimes/osx-arm64/native -lfaiss_c" \
go build ./pkg/ivf
```

3. To run the test
```shell
cd pkg/ivf
go test -v -run TestFaissKmeans
```


### Note
- Build the FAISS library using
```shell
cmake -DCMAKE_BUILD_TYPE=Release -DFAISS_ENABLE_C_API=ON -DBUILD_SHARED_LIBS=OFF -B build .
```
It creates `libfaiss_c.a` file instead of `libfaiss_c.dylib`.

- Build script [here](/cgo/thirdparty/build-faiss-macos.sh)

- The build file is copied [here](/cgo/thirdparty/runtimes/osx-arm64/native)

- The Kmeans CGO code is [hear](/pkg/ivf/clustering_faiss.go)
```cgo
/*
#cgo LDFLAGS: ${SRCDIR}/../../cgo/thirdparty/runtimes/osx-arm64/native/libfaiss_c.a

#include <stdlib.h>
#include <faiss/c_api/Clustering_c.h>
#include <faiss/c_api/impl/AuxIndexStructures_c.h>
#include <faiss/c_api/index_factory_c.h>
#include <faiss/c_api/error_c.h>
*/
```
The Kmeans code fails due this error: `Undefined symbols for architecture arm64:`.

However, this code works when we build `libfaiss_c.dylib` using `-DBUILD_SHARED_LIBS=ON` and the `libfaiss_c.dylib` is added to the `/usr/local/lib`.
Then, we add the CGO flag in the go code
````cgo
/*
#cgo LDFLAGS: -lfaiss_c

#include <stdlib.h>
#include <faiss/c_api/Clustering_c.h>
#include <faiss/c_api/impl/AuxIndexStructures_c.h>
#include <faiss/c_api/index_factory_c.h>
#include <faiss/c_api/error_c.h>
*/
````

But I am trying to package FAISS C_API as a static library and link it to the go code. Hence, I am building `libfaiss_c.a` locally and packaging as a library.

- The test driver code is [here](/pkg/ivf/clustering_faiss_test.go)


### Issue

1. FAISS CGO throwing "Undefined symbols for architecture arm64" when using `libfaiss_c.a` in MacOS M2

<details>
<summary> Error Log </summary>

```log
# faiss-go/pkg/ivf.test
/usr/local/go/pkg/tool/darwin_arm64/link: running clang failed: exit status 1
Undefined symbols for architecture arm64:
  "faiss::Clustering::Clustering(int, int)", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
  "faiss::Clustering::Clustering(int, int, faiss::ClusteringParameters const&)", referenced from:
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
  "faiss::kmeans_clustering(unsigned long, unsigned long, unsigned long, float const*, float*)", referenced from:
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "faiss::ClusteringParameters::ClusteringParameters()", referenced from:
      _faiss_ClusteringParameters_init in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
  "std::exception_ptr::exception_ptr(std::exception_ptr const&)", referenced from:
      _faiss_get_last_error in libfaiss_c.a(error_impl.cpp.o)
  "std::exception_ptr::~exception_ptr()", referenced from:
      _faiss_get_last_error in libfaiss_c.a(error_impl.cpp.o)
      thread-local wrapper routine for faiss_last_exception in libfaiss_c.a(error_impl.cpp.o)
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering.cold.1 in libfaiss_c.a(Clustering_c.cpp.o)
      ...
  "std::exception_ptr::operator=(std::exception_ptr const&)", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering.cold.1 in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering.cold.2 in libfaiss_c.a(Clustering_c.cpp.o)
  "std::runtime_error::runtime_error(char const*)", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "std::runtime_error::runtime_error(std::runtime_error const&)", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "std::runtime_error::~runtime_error()", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering.cold.2 in libfaiss_c.a(Clustering_c.cpp.o)
  "std::__1::basic_string<char, std::__1::char_traits<char>, std::__1::allocator<char>>::basic_string(std::__1::basic_string<char, std::__1::char_traits<char>, std::__1::allocator<char>> const&)", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::FaissException(faiss::FaissException const&) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "std::exception::~exception()", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering.cold.1 in libfaiss_c.a(Clustering_c.cpp.o)
  "std::exception::~exception()", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::FaissException(faiss::FaissException const&) in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::~FaissException() in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "std::current_exception()", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "std::rethrow_exception(std::exception_ptr)", referenced from:
      _faiss_get_last_error in libfaiss_c.a(error_impl.cpp.o)
  "std::terminate()", referenced from:
      ___clang_call_terminate in libfaiss_c.a(Clustering_c.cpp.o)
  "typeinfo for faiss::FaissException", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table21 in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table28 in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table29 in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table31 in libfaiss_c.a(Clustering_c.cpp.o)
      ...
  "typeinfo for std::runtime_error", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "typeinfo for std::exception", referenced from:
      GCC_except_table0 in libfaiss_c.a(error_impl.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table21 in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table28 in libfaiss_c.a(Clustering_c.cpp.o)
      GCC_except_table29 in libfaiss_c.a(Clustering_c.cpp.o)
      ...
  "vtable for faiss::FaissException", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::FaissException(faiss::FaissException const&) in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::~FaissException() in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  NOTE: a missing vtable usually means the first non-inline virtual member function has no definition.
  "vtable for std::exception", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  NOTE: a missing vtable usually means the first non-inline virtual member function has no definition.
  "operator delete(void*)", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::~FaissException() in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "operator new(unsigned long)", referenced from:
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
  "___cxa_allocate_exception", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "___cxa_begin_catch", referenced from:
      _faiss_get_last_error in libfaiss_c.a(error_impl.cpp.o)
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      ___clang_call_terminate in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      ...
  "___cxa_end_catch", referenced from:
      _faiss_get_last_error in libfaiss_c.a(error_impl.cpp.o)
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      ...
  "___cxa_free_exception", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "___cxa_throw", referenced from:
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_train in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_kmeans_clustering in libfaiss_c.a(Clustering_c.cpp.o)
  "___gxx_personality_v0", referenced from:
      _faiss_get_last_error in libfaiss_c.a(error_impl.cpp.o)
      _faiss_Clustering_new in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::runtime_error>(std::runtime_error) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<std::exception>(std::exception) in libfaiss_c.a(Clustering_c.cpp.o)
      std::exception_ptr std::make_exception_ptr[abi:v15006]<faiss::FaissException>(faiss::FaissException) in libfaiss_c.a(Clustering_c.cpp.o)
      faiss::FaissException::FaissException(faiss::FaissException const&) in libfaiss_c.a(Clustering_c.cpp.o)
      _faiss_Clustering_new_with_params in libfaiss_c.a(Clustering_c.cpp.o)
      ...
ld: symbol(s) not found for architecture arm64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```


</details>

