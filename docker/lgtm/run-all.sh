#!/bin/bash

/opt/run-loki.sh &
/opt/run-grafana.sh &
/opt/run-prometheus.sh &
/opt/run-tempo.sh &

sleep infinity
