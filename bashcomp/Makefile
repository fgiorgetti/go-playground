all: build install

build:
	go build -ldflags "-X main.BashCompletionEncoded=`cat bash_completion.sh | base64 -w 0`" -o bashcomp bashcomp.go

install:
	./bashcomp completion > bashcomp.bash.inc
	@echo
	@echo Now you must run: "source bashcomp.bash.inc"
	@echo
	@echo "... and make sure you have 'bashcomp' binary in your PATH"
