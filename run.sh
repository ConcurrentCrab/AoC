#!/bin/sh

e="$1";
i="$2";
cat "inputs/${i}.txt" | go run "solutions/${e}.go";
cat answers.txt | grep "$e $i: "
