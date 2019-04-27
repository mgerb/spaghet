build:
	go build -o ./dist/spaghet ./main.go
	
windows:
	GOOS=windows GOARCH=386 go build -o ./dist/spaghet.exe ./main.go

clean:
	rm -rf ./dist

all: build windows
