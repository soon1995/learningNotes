FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-01

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update
RUN apt-get -yqq install ruby ruby-dev build-essential
# uwsgi write to Consul
RUN gem install --no-document json

RUN mkdir -p /opt/distributed_client
ADD client.rb /opt/distributed_client/

WORKDIR /opt/distributed_client

ENTRYPOINT [ "ruby", "/opt/distributed_client/client.rb" ]
CMD []
