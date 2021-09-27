build:
	go build -o program .

run:
	make build
	./program

clear:
	rm program

init:
	docker pull tensorflow/tensorflow
	sudo cp /vagrant/daemon-config.json /etc/docker/daemon.json
	bash scripts/install-runsc.sh
	bash scripts/install-crun.sh
	bash scripts/install-youki.sh
	sudo systemctl restart docker.service

all:
	make init
	make run
