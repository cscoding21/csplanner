FROM ubuntu:22.04

WORKDIR /app

RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y wget
RUN apt-get install -y net-tools
RUN apt-get install -y iputils-ping

RUN curl -sSf https://install.surrealdb.com | sh
RUN curl -sf https://binaries.nats.dev/nats-io/natscli/nats@latest | sh
RUN mv nats /usr/local/bin
RUN nats context add nats --server nats.nats.svc:4222 --description "NATS Instance" --select

ENTRYPOINT ["/bin/bash"]