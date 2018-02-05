#!/bin/bash

for number in {1..10000}
do
    echo $(date --iso-8601=seconds) $number
    sleep 1
done
exit 0