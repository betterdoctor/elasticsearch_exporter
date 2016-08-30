FROM scratch

COPY elasticsearch_exporter /elasticsearch_exporter
ENTRYPOINT ["/elasticsearch_exporter"]
