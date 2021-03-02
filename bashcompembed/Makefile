all: build install

build:
	go build -o bashcompembed bashcompembed.go

install:
	./bashcompembed completion > bashcompembed.bash.inc
	@echo
	@echo Now you must run: "source bashcompembed.bash.inc"
	@echo
	@echo "... and make sure you have 'bashcompembed' binary in your PATH"
