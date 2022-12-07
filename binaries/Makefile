REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')/binaries
TIMESTAMP = $(shell date +%s)
GO = go
GO_SRCS = $(shell find  .  -type f -regex  ".*.go$$")
TARGETS = $(shell ls cmd)
VERSION ?= v0.1.0

COMMIT_ID ?= $(shell git rev-parse --short HEAD)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS += -X "$(REPO)/version.BuildTS=$(shell TZ='Asia/Shanghai' date '+%Y-%m-%d %H:%M:%S')"
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

migrate:
	migrate -source file://./migrations/sqlite3/ -database sqlite3://../storage/ up 2