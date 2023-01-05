ifeq ($(OS), Windows_NT)
SHELL = cmd.exe
# Windows
TARGETS = $(shell dir cmd /b)
GO_SRCS = $(shell for /r . %%i in (*.go) do @echo %%i)
REPO = github.com/AHAOAHA/Annal/binaries
BUILT_TS ?= $(shell echo %date:~0,4%%date:~5,2%%date:~8,2%%time:~0,2%%time:~3,2%%time:~6,2%)
BINARY_SUFFIX = .exe

else
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S), Linux)
# Linux
REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')/binaries
TIMESTAMP = $(shell date +%s)
BUILT_TS ?= $(shell TZ='Asia/Shanghai' date '+%Y-%m-%d %H:%M:%S')

GO_SRCS = $(shell find  .  -type f -regex  ".*.go$$")
TARGETS = $(shell ls cmd)
endif
endif

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
	${CGO_BUILD_OP} $(GO) build -ldflags '${LDFLAGS} -X "$(REPO)/version.App=$@"' -tags='$(TAGS)' -o $@$(BINARY_SUFFIX) $(REPO)/cmd/$@/

test:
	go test ./... -coverprofile=${COVERAGE_REPORT} -covermode=atomic -tags='$(TAGS)'

clean:
	-rm -rf $(TARGETS)
