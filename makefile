all: main

main: main/bayou_primary main/bayou_server main/bayou_bot

main/%:
	go build main/$*.go
clean:
	rm bayou_bot bayou_server bayou_primary
