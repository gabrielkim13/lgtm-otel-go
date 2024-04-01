FROM ubuntu:22.04

ENV LOKI_VERSION="2.9.6"
ENV GRAFANA_VERSION="10.4.1"
ENV PROMETHEUS_VERSION="2.51.1"
ENV TEMPO_VERSION="2.4.1"

RUN apt-get update && apt-get install -y \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sOL https://github.com/grafana/loki/releases/download/v$LOKI_VERSION/loki-linux-amd64.zip && \
    mkdir /opt/loki && \
    unzip loki-linux-amd64.zip -d /opt/loki && \
    rm loki-linux-amd64.zip

RUN curl -sOL https://dl.grafana.com/oss/release/grafana-$GRAFANA_VERSION.linux-amd64.tar.gz && \
    tar zxf grafana-$GRAFANA_VERSION.linux-amd64.tar.gz -C /opt && \
    mv /opt/grafana-v$GRAFANA_VERSION /opt/grafana && \
    rm grafana-$GRAFANA_VERSION.linux-amd64.tar.gz

RUN curl -sOL https://github.com/prometheus/prometheus/releases/download/v$PROMETHEUS_VERSION/prometheus-$PROMETHEUS_VERSION.linux-amd64.tar.gz && \
    tar zxf prometheus-$PROMETHEUS_VERSION.linux-amd64.tar.gz -C /opt && \
    mv /opt/prometheus-$PROMETHEUS_VERSION.linux-amd64 /opt/prometheus && \
    rm prometheus-$PROMETHEUS_VERSION.linux-amd64.tar.gz

RUN curl -sOL https://github.com/grafana/tempo/releases/download/v$TEMPO_VERSION/tempo_${TEMPO_VERSION}_linux_amd64.tar.gz && \
    mkdir /opt/tempo && \
    tar zxf tempo_${TEMPO_VERSION}_linux_amd64.tar.gz -C /opt/tempo && \
    rm tempo_${TEMPO_VERSION}_linux_amd64.tar.gz

COPY loki.yaml /opt/loki
COPY run-loki.sh /opt

COPY grafana-datasources.yaml /opt/grafana/conf/provisioning/datasources/
COPY run-grafana.sh /opt

COPY prometheus.yaml /opt/prometheus
COPY run-prometheus.sh /opt

COPY tempo.yaml /opt/tempo
COPY run-tempo.sh /opt

COPY run-all.sh /opt

EXPOSE 3100 3000 9090 3200 9096 4317 4318

CMD ["/opt/run-all.sh"]
