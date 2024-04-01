#!/bin/bash

cd /opt/prometheus

./prometheus \
    --web.enable-remote-write-receiver \
    --enable-feature=exemplar-storage \
    --enable-feature=native-histograms \
    --config.file=./prometheus.yaml |& tee /var/log/prometheus.log
