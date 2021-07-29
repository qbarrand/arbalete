all: arbalete

arbalete: $(shell find . -type f -name '*.go') go.*
	go build -o $@

clean:
	rm arbalete
