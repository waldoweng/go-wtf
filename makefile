objects := $(patsubst %.go,%,$(wildcard *.go))

.PHONY: all clean

all: $(objects)

% : %.go
	go build $<

clean:
	rm -f $(objects)