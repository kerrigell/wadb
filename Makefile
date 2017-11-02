### Makefile for wadb

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

path_to_add := $($(GOPATH))
export PATH := $(path_to_add):$(PATH)


GOLEX   := golex
GOYACC  := goyacc
GOLINT  := golint

GO        := go
GOBUILD   := $(GO) build $(BUILD_FLAG)
GOTEST    := $(GO) test -p 3 

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"
PACKAGES  := $$(go list ./...| grep -vE 'vendor')

.PHONY:  parser goyacc golex


all: parser test 

golex:
	$(GO) get github.com/qiuyesuifeng/golex

goyacc:
	$(GOBUILD) -o bin/goyacc parser/goyacc/main.go

packges:
	$(GO) get github.com/kerrigell/wadb/util/charset
	$(GO) get github.com/qiuyesuifeng/golex

parser: goyacc
	bin/goyacc -o /dev/null parser/parser.y
	bin/goyacc -o parser/parser.go parser/parser.y 2>&1 | egrep "(shift|reduce)/reduce" | awk '{print} END {if (NR > 0) {print "Find conflict in parser.y. Please check y.output for more information."; exit 1;}}'
	rm -f y.output
	
	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		sed -i -e 's|//line.*||' -e 's/yyEofCode/yyEOFCode/' parser/parser.go; \
	elif [ $(ARCH) = $(MAC) ]; \
	then \
		/usr/bin/sed -i "" 's|//line.*||' parser/parser.go; \
		/usr/bin/sed -i "" 's/yyEofCode/yyEOFCode/' parser/parser.go; \
	fi

	@awk 'BEGIN{print "// Code generated by goyacc"} {print $0}' parser/parser.go > tmp_parser.go && mv tmp_parser.go parser/parser.go;

test:
	$(GO) test github.com/kerrigell/wadb/parser