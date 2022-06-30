OS_RELEASE = $(shell awk -F= '/^NAME/{print $$2}' /etc/os-release | tr A-Z a-z)
TIMESTAMP = $(shell date +%s)
MKFILE_PATH = $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
RCS_DIR = appc

RCS = .zshrc .zshenv .bashrc .envrc .vimrc
CONFIGS = .p10k.zsh .tmux.conf.local 
LINK_FILES = $(foreach file, $(RCS), $(MKFILE_PATH)/rcs/$(file))
LINK_FILES += $(foreach file, $(CONFIGS), $(MKFILE_PATH)/configs/$(file))

# 来自submodule的工具
SUBMODULE_PLUGINS = ohmyzsh ohmytmux
# 来自包管理的工具
INSTALL_PLUGINS = vim
PLUGINS = $(SUBMODULE_PLUGINS) $(INSTALL_PLUGINS)

ENV_TARGETS = $(LINK_FILES) $(PLUGINS)

ifeq ($(findstring "ubuntu", $(OS_RELEASE)),)
	PKG_MANAGER := apt
endif

ifeq ($(findstring "centos", $(OS_RELEASE)),)
	PKG_MANAGER := yum
endif

ifeq ($(USER), "root")
	SUDO := sudo
endif

v:
	echo $(USER)

env: $(ENV_TARGETS)

$(INSTALL_PLUGINS):
	$(SUDO) $(PKG_MANAGER) install $(INSTALL_PLUGINS) -y

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
	-mv plugins/ohmyzsh/custom/plugins/$@ plugins/ohmyzsh/custom/plugins/$@.bak.$(TIMESTAMP)
	ln -sr plugins/$@ plugins/ohmyzsh/custom/plugins

powerlevel10k:
	-mv plugins/ohmyzsh/custom/themes/$@ plugins/ohmyzsh/custom/themes/$@.bak.$(TIMESTAMP)
	ln -sr plugins/$@ plugins/ohmyzsh/custom/themes

.PHONY: $(LINK_FILES)
