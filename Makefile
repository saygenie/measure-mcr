build:
	go build -o program .

run:
	make build
	./program

clear:
	rm program

init:
	docker pull tensorflow/tensorflow
	bash install-runnc.sh
	bash install-runsc.sh
	bash install-youki.sh


all:
	make init
	make run
