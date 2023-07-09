FROM ubuntu:latest
MAINTAINER James Turnbull "james@example.com"
ENV REFRESHED_AT 2016-06-01

RUN apt-get update -y; apt-get -y install ruby ruby-dev build-essential redis-tools
# RUN gem install --no-rdoc --no-ri sinatra json redis
RUN gem install --no-document sinatra json redis

RUN mkdir -p /opt/webapp

EXPOSE 4567

CMD [ "/opt/webapp/bin/webapp" ]
