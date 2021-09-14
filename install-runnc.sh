go get github.com/nabla-containers/runnc
sudo apt install genisoimage
sudo apt install jq

# Go to the repo
cd $GOPATH/src/github.com/nabla-containers/runnc

# Get the neceesary binaries for the runtime
make build

# Install libseccomp on the host
sudo apt install libseccomp-dev

# Install the appropriate binaries/libraries
make install