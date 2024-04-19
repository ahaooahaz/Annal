
TARGETS := client server

all: $(TARGETS)

$(TARGETS):
	@echo "Building $@"
	@go build -o $@ cmd/$@/main.go

clean:
	@rm -f $(TARGETS)