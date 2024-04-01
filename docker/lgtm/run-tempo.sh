#!/bin/bash

cd /opt/tempo

./tempo --config.file=./tempo.yaml |& tee /var/log/tempo.log
