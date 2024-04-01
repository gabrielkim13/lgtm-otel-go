#!/bin/bash

/opt/run-grafana-agent.sh &
/opt/run-otel-go.sh &

sleep infinity
