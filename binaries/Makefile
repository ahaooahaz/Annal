ifeq ($(OS), Windows_NT)
SHELL = cmd.exe
# Windows
TARGETS = $(shell dir cmd /b)
GO_SRCS = $(shell for /r . %%i in (*.go) do @echo %%i)

else
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Linux)
# Linux
REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')
TIMESTAMP = $(shell date +%s)
BUILT_TS ?= $(shell TZ='Asia/Shanghai' date '+%Y-%m-%d %H:%M:%S')

GO_SRCS = $(shell find  .  -type f -regex  ".*.go$$")
TARGETS = $(shell ls cmd)
endif
endif

v:
	echo $(GO_SRCS)

COMMIT_ID ?= $(shell git rev-parse --short HEAD)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

GO = go
VERSION ?= 1.0.1

LDFLAGS += -X "$(REPO)/version.BuildTS=$(BUILT_TS)"
LDFLAGS += -X "$(REPO)/version.GitHash=$(COMMIT_ID)"
LDFLAGS += -X "$(REPO)/version.Version=$(VERSION)"
LDFLAGS += -X "$(REPO)/version.GitBranch=$(BRANCH)"

build: $(TARGETS)

$(TARGETS): $(GO_SRCS)
	${CGO_BUILD_OP} $(GO) build -ldflags '${LDFLAGS} -X "$(REPO)/version.App=$@"' -tags='$(TAGS)' -o $@ $(REPO)/cmd/$@/

test:
	go test ./... -coverprofile=${COVERAGE_REPORT} -covermode=atomic -tags='$(TAGS)'

clean:
	-rm -rf $(TARGETS)
