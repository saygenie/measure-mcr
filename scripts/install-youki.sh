sudo apt-get install -y \
    pkg-config         \
    libsystemd-dev     \
    libdbus-glib-1-dev \
    build-essential    \
    libelf-dev

curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source ~/.cargo/env

cd ~
git clone https://github.com/containers/youki
cd youki
bash build.sh
sudo mv target/x86_64-unknown-linux-gnu/debug/youki /usr/local/bin/
cd ..
rm -rf youki