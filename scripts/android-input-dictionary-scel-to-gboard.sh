#!/bin/bash
set -ex
scelfile=$(realpath $1)
scelfiledir=$(dirname $scelfile)
scelfilename=$(basename $scelfile)

docker run --rm -it -v $scelfiledir:/dict imewlconverter -i:scel /dict/$scelfilename -o:ggpy /dict/$scelfilename.tmp.txt
iconv -f GBK -t UTF-8 $scelfiledir/$scelfilename.tmp.txt -o $scelfiledir/$scelfilename.tmp.utf8.txt
echo "# Gboard Dictionary version:1" > $scelfiledir/$scelfilename.txt
awk -F'\t' '{ gsub(/[\r\n]+/, "", $3); print $3 "\t" $1 "\tzh-CN"}' $scelfiledir/$scelfilename.tmp.utf8.txt > $scelfiledir/$scelfilename.txt

