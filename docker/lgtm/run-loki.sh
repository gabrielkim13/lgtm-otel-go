#!/bin/bash

cd /opt/loki

./loki-linux-amd64 --config.file=./loki.yaml |& tee /var/log/loki.log
