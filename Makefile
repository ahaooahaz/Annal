REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')
VERSION = 1.0.1
OS_RELEASE = $(shell awk -F= '/^NAME/{print $$2}' /etc/os-release | tr A-Z a-z)
TIMESTAMP = $(shell date +%s)
MKFILE_PATH = $(shell pwd)
RCS_DIR = appc
ANNALRC = $${HOME}/.annalrc
INSTALL_PATH = $${HOME}/.local/bin
SHELL = /bin/bash
.SHELLFLAGS := -e -u -o pipefail -c
VERBOSE ?= 1

RCS = .zshrc .zshenv .bashrc .envrc .vimrc .aliases .bash_profile .profile
CONFIGS_RCS = .p10k.zsh .tmux.conf.local
HOME_LINK_FILES = $(foreach file, $(RCS), $(MKFILE_PATH)/rcs/$(file))
HOME_LINK_FILES += $(foreach file, $(CONFIGS), $(MKFILE_PATH)/configs/$(file))
SCRIPTS = $(shell find scripts -type f)

PACKAGE_PLUGINS = sshpass base64 at xsel

ifneq ($(findstring "ubuntu", $(OS_RELEASE)),)
	PKG_MANAGER := apt
endif

ifneq ($(findstring "centos", $(OS_RELEASE)),)
	PKG_MANAGER := yum
endif

ifneq ($(USER), "root")
	SUDO := sudo
endif

all: build

env:
	python3 install.py

$(PACKAGE_PLUGINS):
	if ! type $@ 2>/dev/null; then $(SUDO) $(PKG_MANAGER) install $@ -y; fi

$(LINK_FILES):
	-mv ~/$(notdir $@) ~/$(notdir $@).bak.$(TIMESTAMP)
	ln -sf $@ ~/

wechat wechat.work:
	-rm ~/.local/bin/$@
	ln -sr scripts/docker-$@.sh ~/.local/bin/$@

$(ZSH_PLUGINS):
	-ln -sr plugins/$@ plugins/ohmyzsh/custom/plugins

powerlevel10k:
	-ln -sr plugins/$@ plugins/ohmyzsh/custom/themes

$(BINARIES_CMDS): build
	cp binaries/$@ .

build:
	$(MAKE) -C binaries

clean:
	$(MAKE) -C binaries $@
	rm -rf $(BINARIES_CMDS)

BINARIES_CMDS = $(shell ls binaries/cmd)
install: $(BINARIES_CMDS)
	mkdir -p $(INSTALL_PATH)
	if [ ! -h $(INSTALL_PATH)/annal ]; then ln -s $(shell pwd)/binaries/annal $(INSTALL_PATH)/annal; fi

UNFILES := $(foreach un, $(BINARIES_CMDS), $(INSTALL_PATH)/$(un))
uninstall:
	rm -f $(UNFILES)

.PHONY: $(LINK_FILES) $(ENV_TARGETS) $(RIME_CONFIGS) $(RIME_DICTS) $(RIME_EMOJI)
$(VERBOSE).SILENT: