# Exercise - SpringBoot Project Image

![image-20220622234151785](./../images/image-20220622234151785.png)

Package

![image-20220622234444372](./../images/image-20220622234444372.png)

Saw my jar

![image-20220622234450012](./../images/image-20220622234450012.png)

create Dockerfile & copy out jar from target

![image-20220622234837285](./../images/image-20220622234837285.png)

```
FROM java:8

COPY *.jar /app.jar

CMD ["--server.port=8080"]

EXPOSE 8080

ENTRYPOINT ["java","-jar","/app.jar"]
```

..

build

..

run

..

curl localhost:port/test/hello

![image-20220623000322582](./../images/image-20220623000322582.png)

succeed!

