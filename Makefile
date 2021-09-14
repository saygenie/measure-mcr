build:
	go build -o program .

run:
	make build
	./program

clear:
	rm program

init:
	TARGET_IMAGE=tensorflow/tensorflow
	docker pull $TARGET_IMAGE
	bash install-runnc.sh
	bash install-runsc.sh
	bash install-youki.sh


all:
	make init
	make run
