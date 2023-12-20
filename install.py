#!/usr/bin/env python3

import json
import os

def built_link(plugins):
    if plugins != None:
        for plugin in plugins:
            if in_actions(plugin.get('actions', None), 'link'):
                links = plugin.get('links', [])
                for link in links:
                    dst = link.get('dst', None)
                    link = link.get('link', None)
                    if dst != None and link != None:
                        abslink = os.path.expanduser(link)
                        absdst = os.path.abspath(dst)
                        os.makedirs(os.path.dirname(link), exist_ok=True)
                        try:
                            os.rename(abslink, abslink+".backup")
                        except:
                            pass
                        print("Linking {} to {}".format(abslink, absdst))
                        os.symlink(absdst, abslink)
            
            built_link(plugin.get('childrens', None))

def clone(plugins):
    if plugins != None:
        for plugin in plugins:
            if in_actions(plugin.get('actions', None), 'clone'):
                name = plugin.get('name', None)
                repo = plugin.get('repo', None)

                if os.path.exists("plugins/{}".format(name)):
                    print("Plugin {} already exists, skipping clone ...".format(name))
                else:
                    if repo != None:
                        print("Cloning {}".format(repo))
                        os.system("git clone https://github.com/{}.git plugins/{}".format(repo, name))
                
            clone(plugin.get('childrens', None))

def in_actions(actions, action):
    for a in actions:
        if a == action:
            return True
    return False

with open('config.json', 'r') as config:
    data = json.load(config)
    clone(data['plugins'])
    built_link(data['plugins'])