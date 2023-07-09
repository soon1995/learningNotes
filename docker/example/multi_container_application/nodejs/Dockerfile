FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-01

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update
RUN apt-get -yqq install nodejs npm
# RUN ln -s /usr/bin/nodejs /usr/bin/node
RUN mkdir -p /var/log/nodeapp

ADD nodeapp /opt/nodeapp/

WORKDIR /opt/nodeapp
RUN npm install

VOLUME [ "/var/log/nodeapp" ]

EXPOSE 3000

# the server.js pulls in all the dependencies and starts an Express application
# The express app is configured to store its session information in Redis and 
# exposes a single endpoint that returns a status message
# We've configured its connection to Redis to use a host called redis_primary with
# an option to override this with an environment variable if needed.
# the app will also log to the /var/log/nodeapp/nodeapp.log file and will
# listen on port 3000.
ENTRYPOINT [ "nodejs", "server.js" ]
