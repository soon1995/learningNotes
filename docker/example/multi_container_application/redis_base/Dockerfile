FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-03

RUN apt-get -yqq update
RUN apt-get -yqq install software-properties-common
RUN add-apt-repository ppa:chris-lea/redis-server
RUN apt-get -yqq update
RUN apt-get -yqq install redis-server redis-tools

VOLUME [ "/var/lib/redis", "/var/log/redis" ]

EXPOSE 6379

# Install latest version of Redis (from PPA rather than using
# the older packages shipped with Ubuntu), specifies two VOLUMES
# , and expose Redis default port
# it does not have an ENTRYPOINT or CMD because we're not actually 
# going to run this image, we're just going to build on top of it
CMD []
