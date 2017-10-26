#!/bin/bash

# Count from 1 to 10, pausing for a second between each iteration.
for COUNT in $(seq 1 10); do
  echo $COUNT
  sleep 1
done
