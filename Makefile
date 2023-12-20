REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')
VERSION = 1.0.1
OS_RELEASE = $(shell awk -F= '/^NAME/{print $$2}' /etc/os-release | tr A-Z a-z)
TIMESTAMP = $(shell date +%s)
MKFILE_PATH = $(shell pwd)
RCS_DIR = appc
ANNALRC = $${HOME}/.annalrc
INSTALL_PATH = $${HOME}/.local/bin
SHELL = /bin/bash
VERBOSE ?= 1

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

.PHONY: env build clean
$(VERBOSE).SILENT: