FROM ubuntu
MAINTAINER zzyy<zzyybs@126.com>

ENV MYPATH /usr/local
WORKDIR $MYPATH

RUN apt-get update
RUN apt-get install net-tools
#RUN apt-get install -y iproute2
#RUN apt-get install -y inetutils-ping

EXPOSE 80

RUN echo $MYPATH
RUN echo "install ifconfig cmd into ubuntu success--------------ok"
CMD /bin/bash
