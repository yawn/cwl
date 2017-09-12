BUILD 	:= $(shell git rev-parse --short HEAD)
FLAGS		:= "-s -w -X=main.build=$(BUILD) -X=main.time=`TZ=UTC date '+%FT%TZ'` -X=main.version=$(VERSION)"
REPO 		:= cwl
TOKEN 	= `cat .token`
USER 		:= yawn
VERSION := "1.0.0"

.PHONY: build clean release retract

build:
	nice gox -parallel=2 -osarch="darwin/amd64 linux/amd64 linux/arm" -ldflags $(FLAGS) -output "build/{{.OS}}-{{.Arch}}/cwl" ./bin/
	find build -type f -print0 | xargs -I '{}' -0 -P 16 nice upx --best --brute -q {} > /dev/null

clean:
	rm -rf build

release:
	git tag $(VERSION) -f && git push --tags -f
	github-release release --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN)
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name cwl-linux --file build/linux-amd64/cwl
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name cwl-linux-arm --file build/linux-arm/cwl
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name cwl-osx --file build/darwin-amd64/cwl

retract:
	github-release delete --tag $(VERSION) -s $(TOKEN)
