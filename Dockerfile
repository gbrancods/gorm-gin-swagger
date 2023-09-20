FROM debian

COPY ./build/linux/ggs /ggs
COPY ./logs /logs
COPY ./config.toml /config.toml

ENV TZ="America/Sao_Paulo"

ENTRYPOINT /ggs