exeName=dummy
sourceName=main.go

all: compileLinux compileWindows

compileLinux:
	go build -o bin/linux/$(exeName) $(sourceName)

compileWindows:
	GOOS=windows GOARCH=386 go build -o bin/windows/$(exeName).exe $(sourceName)

install:
	cp bin/linux/$(exeName) ${GOPATH}/bin/

clean:
	rm -f bin/linux/$(exeName) ; rm -f bin/windows/$(exeName).exe
