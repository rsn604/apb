SHELL=/bin/bash
GOBUILD=go build
GORUN=go run
MAIN_PROG=apb
SRCS=main.go
DB=ApptDB.boltdb
#DB=util/HPL.boltdb
# ---------------------------------------------------
clean:
	-@rm $(MAIN_PROG)

# ---------------------------------------------------
gofmt:
	@gofmt -l -s -w .

# ---------------------------------------------------
run:
	$(GORUN) $(SRCS) $(DB)

# ---------------------------------------------------
build:
	$(GOBUILD) -o $(MAIN_PROG) $(SRCS)

# ---------------------------------------------------
build-win64:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(MAIN_PROG)64.exe $(SRCS)

# ---------------------------------------------------
build-win32:
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(MAIN_PROG)32.exe $(SRCS)

# ---------------------------------------------------
build-arm6:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(MAIN_PROG)_arm6 $(SRCS)

# ---------------------------------------------------
build-arm7:
	GOOS=linux GOARCH=arm GOARM=7 $(GOBUILD) -o $(MAIN_PROG)_arm7 $(SRCS)

# ---------------------------------------------------
build-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(MAIN_PROG)_arm64 $(SRCS)
