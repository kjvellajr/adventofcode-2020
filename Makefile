all:
	@find . -name '*.go' -print0 | sort -z | xargs -r0 -L 1 go run
