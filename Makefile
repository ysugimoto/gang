.PHONY: brew clean

brew:
	@echo "Building for Darwin amd64"
	GOOS=darwin GOARC=amd64 go build -o ./build/amd64/gang cmd/gang/gang.go
	cd ./build/amd64 && tar cvfz gang_amd64.tar.gz gang

	@echo "Building for Darwin i386"
	GOOS=darwin GOARC=386 go build -o ./build/i386/gang cmd/gang/gang.go
	cd ./build/i386 && tar cvfz gang_i386.tar.gz gang

clean:
	rm -rf ./build


