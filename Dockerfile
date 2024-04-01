FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./otel-go

FROM ubuntu:22.04

ENV GRAFANA_AGENT_VERSION="0.40.3"

RUN apt-get update && apt-get install -y \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sOL https://github.com/grafana/agent/releases/download/v$GRAFANA_AGENT_VERSION/grafana-agent-linux-amd64.zip && \
    unzip grafana-agent-linux-amd64.zip -d /bin && \
    mv /bin/grafana-agent-linux-amd64 /bin/grafana-agent && \
    rm grafana-agent-linux-amd64.zip

COPY docker/otel-go/grafana-agent.yaml /etc/grafana-agent.yaml

COPY --from=builder /app/otel-go /usr/bin/otel-go

COPY docker/otel-go/*.sh /opt/

EXPOSE 8080

CMD ["/opt/run-all.sh"]
