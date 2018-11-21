BINARY_NAME="go-crawler"

build:
	@go build -o $(BINARY_NAME) -v

run:
	@go build -o $(BINARY_NAME) -v
	@./$(BINARY_NAME)

clean: 
	@go clean
	@rm -f $(BINARY_NAME)

deps:
	@go get github.com/jawher/mow.cli
	@go get github.com/parnurzeal/gorequest
	@go get github.com/PuerkitoBio/goquery
	@go get gopkg.in/go-playground/pool.v3
