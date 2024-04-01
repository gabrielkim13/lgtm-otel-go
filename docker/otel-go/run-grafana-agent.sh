#!/bin/bash

/bin/grafana-agent \
    -config.file=/etc/grafana-agent.yaml \
    |& tee /var/log/grafana-agent.log
