#!/bin/bash

cd /opt/grafana

./bin/grafana server |& tee /var/log/grafana.log
