#!/bin/bash

BRANCH="main"

brew install libomp
brew install cmake
export CMAKE_PREFIX_PATH=/opt/homebrew/opt/openblas:/opt/homebrew/opt/libomp:/opt/homebrew
git clone --recursive --branch $BRANCH https://github.com/facebookresearch/faiss.git libfaiss-src
cd libfaiss-src
cmake -DFAISS_ENABLE_GPU=OFF -DFAISS_ENABLE_PYTHON=OFF -DBUILD_TESTING=OFF -DCMAKE_BUILD_TYPE=Release -DFAISS_ENABLE_C_API=ON -DBUILD_SHARED_LIBS=OFF -B build .
make -C build -j faiss
sudo make -C build install

arch=arm64
if [[ $(uname -m) == 'x86_64' ]]; then
  arch=x64
fi

cp build/c_api/libfaiss_c.a ../runtimes/osx-$arch/native/
cd ..