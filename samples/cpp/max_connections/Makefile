objects = server client

all: $(objects)

$(objects):%:%.cc
	g++ -o $@ $< -std=c++11

clean:
	-rm -rf $(objects)
.PHONY: clean