FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-01

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update
RUN apt-get -yqq install ruby-dev git libcurl4-openssl-dev curl build-essential python
# uwsgi write to Consul
RUN gem install --no-document uwsgi sinatra

RUN mkdir -p /opt/distributed_app
WORKDIR /opt/distributed_app

RUN uwsgi --build-plugin https://github.com/unbit/uwsgi-consul

ADD uwsgi-consul.ini /opt/distriuted_app/
ADD config.ru /opt/distributed_app/

ENTRYPOINT [ "uwsgi", "--ini", "uwsgi-consul.ini", "--ini", "uwsgi-consul.init:server1", "--ini", "uwsgi-consult.ini:server2" ]
CMD []
