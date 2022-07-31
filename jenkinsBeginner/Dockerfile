FROM openjdk:8-jdk-alpine
VOLUME /tmp
COPY ./target/demoJenkins.jar demoJenkins.jar
ENTRYPOINT ["java","-jar","/demoJenkins.jar","&"]
