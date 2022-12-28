REPO = $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)//;s/\.git$$//' | sed 's/:/\//g')
VERSION = 1.0.1
OS_RELEASE = $(shell awk -F= '/^NAME/{print $$2}' /etc/os-release | tr A-Z a-z)
TIMESTAMP = $(shell date +%s)
MKFILE_PATH = $(shell pwd)
RCS_DIR = appc
ANNALRC = $${HOME}/.annalrc
INSTALL_PATH = $${HOME}/.local/bin

RCS = .zshrc .zshenv .bashrc .envrc .vimrc .aliases
CONFIGS = .p10k.zsh .tmux.conf.local 
LINK_FILES = $(foreach file, $(RCS), $(MKFILE_PATH)/rcs/$(file))
LINK_FILES += $(foreach file, $(CONFIGS), $(MKFILE_PATH)/configs/$(file))

BINARIES_CMDS = $(shell ls binaries/cmd)

# 来自submodule的工具
SUBMODULE_PLUGINS = ohmyzsh ohmytmux
# 来自包管理的工具, TODO: VERSION(9.0)
INSTALL_PLUGINS = sshpass base64
PLUGINS = $(SUBMODULE_PLUGINS) $(INSTALL_PLUGINS)

ENV_TARGETS = $(LINK_FILES) $(PLUGINS)
CMD_TARGETS = $(CMDS)

OUTPUT_BINARIES = bin
INSTALL_TARGETS = $(foreach cmd, $(CMD_TARGETS), $(OUTPUT_BINARIES)/$(cmd))

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

env: $(ENV_TARGETS)
	echo "export ANNAL_ROOT=$(MKFILE_PATH)" > ${ANNALRC}

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

$(BINARIES_CMDS): build
	cp binaries/$@ .

build:
	$(MAKE) -C binaries

clean:
	$(MAKE) -C binaries $@
	rm -rf $(BINARIES_CMDS)

install: $(BINARIES_CMDS)
	mkdir -p $(INSTALL_PATH)
	cp $^ $(INSTALL_PATH)

UNFILES := $(foreach un, $(BINARIES_CMDS), $(INSTALL_PATH)/$(un))
uninstall:
	rm -f $(UNFILES)

.PHONY: $(LINK_FILES) $(ENV_TARGETS)
#$(VERBOSE).SILENT:
