# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/focal64"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "8192"
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo apt-get update &&
    curl -sSL https://get.docker.com/ | sh && # 도커 설치
    usermod -a -G docker vagrant &&
    sudo apt-get install -y gcc make && # make 설치
    sudo snap install go --classic && # go 설치
    echo "export GOPATH=~/go" >> ~/.bashrc &&
    echo "export PATH=$PATH:~/go/bin" >> ~/.bashrc &&
    source ~/.bashrc
  SHELL
end
