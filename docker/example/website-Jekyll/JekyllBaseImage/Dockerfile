FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-01

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update
RUN apt-get -yqq install ruby ruby-dev build-essential nodejs
RUN gem install jekyll -v 2.5.3

# hold new website source code
VOLUME /data
# hold compiled Jekyll site
VOLUME /var/www/html
WORKDIR /data

# build Jekyll site it finds in the /data/ into /var/www/html/ directory
ENTRYPOINT [ "jekyll", "build", "--destination=/var/www/html" ]
