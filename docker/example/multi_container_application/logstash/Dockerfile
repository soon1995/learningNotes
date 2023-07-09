FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-04

RUN apt-get -yqq update
RUN apt-get -yqq install wget gpg apt-transport-https
RUN wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | gpg --dearmor -o /usr/share/keyrings/elastic-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/elastic-keyring.gpg] https://artifacts.elastic.co/packages/8.x/apt stable main" | tee -a /etc/apt/sources.list.d/elastic-8.x.list
RUN apt-get -yqq update
RUN apt-get -yqq install logstash default-jdk

ADD logstash.conf /etc/

WORKDIR /usr/share/logstash

ENTRYPOINT [ "bin/logstash" ]
CMD [ "-f", "/etc/logstash.conf", "--path.settings=/etc/logstash" ]
