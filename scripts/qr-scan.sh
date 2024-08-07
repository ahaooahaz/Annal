#!/bin/bash

flameshot gui -r | zbarimg - 2>/dev/null | sed "s/^QR-Code://g" | xsel --clipboard --input 2>/dev/null
