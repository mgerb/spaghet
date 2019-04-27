build:
	go build -o ./dist/spaghet ./main.go
	
windows:
	GOOS=windows GOARCH=386 go build -o ./dist/spaghet.exe ./main.go

clean:
	rm -rf ./dist

zip:
	zip -j ./dist/windows.zip ./dist/spaghet.exe
	zip -j ./dist/linux.zip ./dist/spaghet

all: build windows zip
