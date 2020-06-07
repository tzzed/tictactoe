NAME := tictactoe

.PHONY: all $(NAME) test testrace build gen

all: $(NAME)

build: $(NAME)

$(NAME):
	 go build main.go

gen: $(NAME)
	go generate ./...

test:
	go test -v -cover -timeout=1m ./...

run: $(NAME)
	go run main.go

