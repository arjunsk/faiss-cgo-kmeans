#!/bin/bash

BRANCH="main"

brew install libomp
brew install cmake
export CMAKE_PREFIX_PATH=/opt/homebrew/opt/libomp:/opt/homebrew
git clone --recursive --branch $BRANCH https://github.com/facebookresearch/faiss.git libfaiss-src
cd libfaiss-src
git reset --hard  d87888b13e7eb339bb9c45825e9d20def6665171
cmake -DFAISS_ENABLE_GPU=OFF -DFAISS_ENABLE_PYTHON=OFF -DBUILD_TESTING=OFF -DCMAKE_BUILD_TYPE=Release -DFAISS_ENABLE_C_API=ON -DBUILD_SHARED_LIBS=OFF -B build .
sudo make -C build -j faiss
sudo make -C build install

arch=arm64
if [[ $(uname -m) == 'x86_64' ]]; then
  arch=x64
fi

cp build/c_api/libfaiss_c.a ../runtimes/osx-$arch/native/
cd ..