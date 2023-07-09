FROM ubuntu:20.04
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-02

RUN apt-get -yqq update
RUN apt-get -yqq install tomcat9 default-jdk

ENV CATALINA_HOME /usr/share/tomcat9
ENV CATALINA_BASE /var/lib/tomcat9
ENV CATALINA_PID /var/run/tomcat9.pid
ENV CATALINA_SH /usr/share/tomcat9/bin/catalina.sh
ENV CATALINA_TMPDIR /tmp/tomcat9-tomcat7-tmp

RUN mkdir -p $CATALINA_TMPDIR

VOLUME [ "/var/lib/tomcat9/webapps/" ]

EXPOSE 8080

ENTRYPOINT [ "/usr/share/tomcat9/bin/catalina.sh", "run"]
