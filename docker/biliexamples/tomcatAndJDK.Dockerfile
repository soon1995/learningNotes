FROM centos:7
MAINTAINER jacksoon<abcdefg@gmail.com>

COPY readme.txt /usr/local/readme.txt
ADD jdk-8u331-linux-x64.tar.gz /usr/local/
ADD apache-tomcat-9.0.63.tar.gz /usr/local/

RUN yum -y install vim

ENV MYPATH /usr/local
WORKDIR $MYPATH

ENV JAVA_HOME /usr/local/jdk1.8.0_331
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
ENV CATALINA_HOME /usr/local/apache-tomcat-9.0.63
ENV CATALINA_BASE /usr/local/apache-tomcat-9.0.63
ENV PATH $PATH;$JAVA_HOME/bin;$CATALINA_HOME/lib;$CATALINA_HOME/bin

EXPOSE 8080

CMD /usr/local/apache-tomcat-9.0.63/bin/startup.sh && tail -F /usr/local/apache-tomcat-9.0.63/logs/catalina.out

# docker run -d -p 9090:8080 -v /home/kuangshen/build/tomcat/test:/usr/local/apache-tomcat-9.0.63/webapps/test
# -v /home/kuangshen/build/tomcat/tomcatlogs/:/usr/local/apache-tomcat-9.0.63/logs mytomcatjdk
