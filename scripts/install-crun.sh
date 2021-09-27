sudo apt-get install -y make git gcc build-essential pkgconf libtool \
   libsystemd-dev libcap-dev libseccomp-dev libyajl-dev \
   go-md2man libtool autoconf python3 automake

cd ~
git clone https://github.com/containers/crun
cd crun
./autogen.sh
./configure
make
sudo mv crun /usr/local/bin/crun
cd ..
rm -rf crun