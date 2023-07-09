FROM jenkins/jenkins
MAINTAINER james@examples.com
ENV REFRESHED_AT 2016-06-01

USER root
# install sudo for jenkins user
RUN apt-get -y update; apt-get install -y sudo wget
RUN echo "jenkins ALL=NOPASSWD: ALL" >> /etc/sudoers
# install Docker binary, we'll use this to connect to our Docker host and run containers for our builds
RUN wget http://get.docker.com/builds/Linux/x86_64/docker-latest.tgz
RUN tar -xvzf docker-latest.tgz
RUN mv docker/* /usr/bin/

USER jenkins
RUN jenkins-plugin-cli --plugins junit git git-client ssh-slaves greenballs chucknorris ws-cleanup
