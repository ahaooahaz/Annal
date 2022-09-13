REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')
VERSION = 0.0.1
OS_RELEASE = $(shell awk -F= '/^NAME/{print $$2}' /etc/os-release | tr A-Z a-z)
TIMESTAMP = $(shell date +%s)
MKFILE_PATH = $(shell pwd)
RCS_DIR = appc
GO = go
GO_SRCS = $(shell find  .  -type f -regex  ".*.go$$")
CMDS = $(shell ls cmd)
ANNALRC = $${HOME}/.annalrc
INSTALL_PATH = $${HOME}/.local/bin

RCS = .zshrc .zshenv .bashrc .envrc .vimrc .aliases
CONFIGS = .p10k.zsh .tmux.conf.local 
LINK_FILES = $(foreach file, $(RCS), $(MKFILE_PATH)/rcs/$(file))
LINK_FILES += $(foreach file, $(CONFIGS), $(MKFILE_PATH)/configs/$(file))

# 来自submodule的工具
SUBMODULE_PLUGINS = ohmyzsh ohmytmux
# 来自包管理的工具, TODO: VERSION(9.0)
INSTALL_PLUGINS = sshpass base64
PLUGINS = $(SUBMODULE_PLUGINS) $(INSTALL_PLUGINS)

ENV_TARGETS = $(LINK_FILES) $(PLUGINS)
CMD_TARGETS = $(CMDS)

OUTPUT_BINARIES = bin
INSTALL_TARGETS = $(foreach cmd, $(CMD_TARGETS), $(OUTPUT_BINARIES)/$(cmd))

ifeq ($(ARCH), arm64)
	CGO_BUILD_OP := CGO_ENABLED=1 GOOS=linux CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=$(ARCH)
endif 

COMMIT_ID ?= $(shell git rev-parse --short HEAD)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
LDFLAGS += -X "$(REPO)/version.BuildTS=$(shell TZ='Asia/Shanghai' date '+%Y-%m-%d %H:%M:%S')"
LDFLAGS += -X "$(REPO)/version.GitHash=$(COMMIT_ID)"
LDFLAGS += -X "$(REPO)/version.Version=$(VERSION)"
LDFLAGS += -X "$(REPO)/version.GitBranch=$(BRANCH)"

ifneq ($(findstring "ubuntu", $(OS_RELEASE)),)
	PKG_MANAGER := apt
endif

ifneq ($(findstring "centos", $(OS_RELEASE)),)
	PKG_MANAGER := yum
endif

ifneq ($(USER), "root")
	SUDO := sudo
endif

# while gocv already installed, build with video tools.
ifneq ($(shell pkg-config --cflags -- opencv4 2>/dev/null),)
	TAGS += gocv
endif

all: cmd

env: $(ENV_TARGETS)
	echo "export ANNAL_ROOT_PATH=$(MKFILE_PATH)" > ${ANNALRC}
cmd: $(CMD_TARGETS)

$(INSTALL_PLUGINS):
	if ! type $@ 2>/dev/null; then $(SUDO) $(PKG_MANAGER) install $@ -y; fi

$(LINK_FILES):
	-mv ~/$(notdir $@) ~/$(notdir $@).bak.$(TIMESTAMP)
	ln -sf $@ ~/

ZSH_PLUGINS = zsh-autosuggestions  zsh-syntax-highlighting
ZSH_THEMES = powerlevel10k

ohmyzsh: $(ZSH_PLUGINS) $(ZSH_THEMES)
	-mv ~/.oh-my-zsh ~/.oh-my-zsh.bak.$(TIMESTAMP)
	ln -sr plugins/$@ ~/.oh-my-zsh

ohmytmux:
	-mv ~/.tmux ~/.tmux.bak.$(TIMESTAMP)
	ln -sr plugins/.tmux ~/.tmux
	-mv ~/.tmux.conf ~/.tmux.conf.bak.$(TIMESTAMP)
	ln -sf ~/.tmux/.tmux.conf ~/

$(ZSH_PLUGINS):
	-ln -sr plugins/$@ plugins/ohmyzsh/custom/plugins

powerlevel10k:
	-ln -sr plugins/$@ plugins/ohmyzsh/custom/themes

# ssh login echo info.
welcome:
	$(SUDO) cp scripts/60-my-welcome-info /etc/update-motd.d

$(CMD_TARGETS): $(GO_SRCS)
	${CGO_BUILD_OP} $(GO) build -ldflags '${LDFLAGS} -X "$(REPO)/version.App=$@"' -tags='$(TAGS)' -o $(OUTPUT_BINARIES)/$@ $(REPO)/cmd/$@/

install: $(INSTALL_TARGETS) scripts/jt
	mkdir -p $(INSTALL_PATH)
	mv $(INSTALL_TARGETS) $(INSTALL_PATH)

test:
	go test ./... -coverprofile=${COVERAGE_REPORT} -covermode=atomic -tags='$(TAGS)'

clean:
	-rm -rf $(OUTPUT_BINARIES)

.PHONY: $(LINK_FILES) $(CMD_TARGETS) $(ENV_TARGETS)
#$(VERBOSE).SILENT:
