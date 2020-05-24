all:
	go build -o meltriv cmd/meltriv/main.go

clean:
	rm -rf meltriv
