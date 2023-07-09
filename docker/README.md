# DockerLearning + Docker Compose

This is a learning notes.

learning Docker from [狂神说](https://www.bilibili.com/video/BV1og4y1q7M4?p=1&vd_source=a788bdd4d7cdd9dfe02852346d523cb9)

Total video length : 8h 11m 41s

21.06.2022 - start

22.06.2022  - complete

reinforcement of knowledge + docker compose refer [尚硅谷-杨哥](https://www.bilibili.com/video/BV1gr4y1U7CY?p=95&vd_source=a788bdd4d7cdd9dfe02852346d523cb9)

11.09.2022 - start

11.09.2022 - complete

Environment: Linux

reinforcement of knowledge: The Docker Book by James Turnbull

Recommended to read this although some of the info were outdated.

05.2023 - 07.2023

# What is Docker

Use for DevOps. to avoid situation :

```
person A : installed different environment and do coding

person B : installed different environment

a project works on A's computer but doesn't works on B's computer
```

Docker to solve as it is a image which store alot of layers of required software including the environment.

May including : mysql, redis, elastic search ....

# Important Step for Docker

DockerFile : Generate image from layer by layer of software.

DockerImages ：A product.

DockerContainer : Run an Image to provide services.

# Commands

Remember : systemctl start docker

![image](https://user-images.githubusercontent.com/97860551/175027625-9ab9214f-c859-4f82-979f-c097c856b914.png)

## Help Command

```
docker version
docker info // returns a list of any containers, any images (the building blocks Docker uses to build containers), the execution and storage drivers Docker is using, and its basic configuration.
docker stats                        //status, includes id, name, cpu %, mem usage/ limit, mem% ...
docker networks ls
docker system df
docker help [command] // eg docker help run
```

![image](https://user-images.githubusercontent.com/97860551/174944151-eaa6b429-c7b2-4b00-9fe5-47bc343f1fd7.png)

```
docker [command] --help
```

**Tips:** Docker has a client-server architecture. It has two binaries, the Docker server provided via the dockerd binary and the docker binary, that acts as a client. As a client, the docker binary passes requests to the Docker daemon (e.g., asking it to return information about itself), and then processes those requests when they are returned. (Prior to Docker 1.12 all of this functionality was provided by a single binary: docker.)

## Image Command

```bash
docker [OPTIONS] images             //show downloaded images
       -a                           //show all
       -q                           //only show id
       
docker search imagename             //search image from docker hub. eg. docker search mysql
docker pull imagename[:version]     //download from docker hub. eg.docker pull mysql:5.7. version refer to docker hub website
docker rmi [OPTIONS] imageid        //remove image, eg. docker rmi -f $(docker -aq images) <-- delete all  
           -f                       //force
docker history imageId              //see how the image generated.   
docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]    //Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
docker image prune      //delete <none>name <none>tag 虚悬镜像
docker images [name] // name is optional to filter
docker create [OPTIONS] image [COMMAND] [ARG...] // create a new container but not run
```

## Container Command

```
docker rm                           // docker rm -f `docker ps -a -q` to FORCE remove all
docker run [OPTION] image command   //run image in a container, eg. docker run -it ubuntu /bin/bash
            --name=?                //container name, else give random name. Name are unique, if we try to create 2 containers with same name, the command will fail
            -h                      // hostnamd of the container
            -d                      //run in background
            -it                     //i(keeps STDIN open from the container, even if were not attached to it) + t(assign a pseudo-tty), thus provide an interative shell in the new container
            -p x                    //the port is randomly assigned (use docker ps -l to see port, or docker port containerID portno)
            -p x:y                  //use localhost port x to visit container port y (expose port y), e.g. 8080:80, interface: 127.0.0.1:80:80, or 127.0.0.1::80 assign random port
            -P                      //publish all exposed ports to random ports
            --rm                    //when exit: delete container
            -e                      //[e]nvironment variable eg. -e ES_JAVA_OPTS="-Xms64m -Xmx512m" (elastic search) 
            -v a:b                  //mount localDir:containerDir (Tips: if containerDir doesn't exist Docker will create it.)
            -v a:b:ro or rw         //read only or read write status of the container directory, ro make container directory read-only
            --volumes-from ContainerIdOrName //share volume from ...,  docker run -d -P --volumes-from james_blog jamtur01/apache
            --log-driver            //control logging driver used, docker run --log-driver='syslog' --name daemon_dwayne -d...
            --restart=always        // checks for the container's exit code and makes a decision whether or not to restart it
              always: no matter what exit code 
              on-failure: non zero exit code, it also accepts optional restart cound eg --restart=on-failure:5 ; it is the max attempt to restart
    while using:  exit              // stop the container
                  ctrl + p + q      //run in background
                                    //eg sudo docker run --name daemon_dave -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"

docker ps [OPTIONS]                 //show currently 'up' container
          -a                        //all including 'exit' container
          -n=?                      //last n numbers
          -q                        //quiet - show only container id
          -l                        //last run container

docker port containerID portNum     // show mapped port
```

![image](https://user-images.githubusercontent.com/97860551/174952561-e545b096-38e4-4639-8e03-1fd9a33a03e9.png)

```
docker start containerId // restart stopped container
docker restart containerId
docker stop containerId/name // stop a daemonized container (it sends a SIGTERM signal to container's running process
docker kill containerId/name // sends a SIGKILL signal to container's process
docker kill -s <signal> <container> // send signals to docker 

docker commit [OPTIONS] containerId imageName:version   //commit and create your image in your images (version ignored will become latest)
              -m="?"                                    //description
              -a="?"                                    //author             
```

### Ways to identify containers

1. a short UUID (like f7cbdac22a02)
2. a longer UUID (like f7cbdac22a02e03c9438c729345e54db9d
20cfa2ac1fc3494b6eb60872e74778)
3. a name (like gray_cat)

## Volumes commands

```
docker volume ls
docker volume inspect
```

## Other useful commands

```
docker logs [OPTIONS] containerId   //show logs
            -ft                     //Follow logs output, Timestamp
            --tail number           //show only last [number]

docker top containerId              //show detail of process, shows the user who run it, and entry command
```

![image](https://user-images.githubusercontent.com/97860551/174939595-50749036-57fe-44b3-b9d1-d07057a2baef.png)

```
docker inspect containerId          //show container's detail
       docker inspect --format='{{ .State.Running }}' container // show state. running.
       docker inspect -f='{{ .NetworkSettings.IPAddress }}' container
       docker inspect -f='{{ .NetworkSettings.IPAddress }} {{ .Name }}' container1 container2 // multiple containers and receive input
       Tips: kindly refer to Go template
docker exec -it containerId  command// running a process inside an already running container, eg docker exec -it abc /bin/bash
docker exec -d containerIdOrNamae command // background eg docker exec -d abc touch /etc/new_config_file
docker attach containerId/name           //reattach to started container session
docker cp containerId:path destPath //copy file from container to host ***********IMPORTANT***********
```

# Visualize (rarely use)

### Portainer

- portainer : Docker's visualize management tool

  ![image](https://user-images.githubusercontent.com/97860551/174950970-b82cda18-de1a-4737-9ed9-213e50cf14fb.png)

```
docker run -d -p 8000:8000 -p 9000:9000 --name portainer     --restart=always     -v /var/run/docker.sock:/var/run/docker.sock     -v portainer_data:/data     portainer/portainer
```

**Restart always :**docker重启后portainer也自动重启

visit port 9090

## CIG

> save the stat of docker container, provide health check
>
> - CAdvisor 监控收集 --- 默认存储2分钟数据 --> InfluxDB
> - InfluxDB 存储数据
> - Granfana 展示图表

> docker-compose.yml

```yaml
version: '3.1'
 
volumes:
  grafana_data: {}
 
services:
 influxdb:
  image: tutum/influxdb:0.9
  restart: always
  environment:
    - PRE_CREATE_DB=cadvisor
  ports:
    - "8083:8083"
    - "8086:8086"
  volumes:
    - ./data/influxdb:/data
 
 cadvisor:
  image: google/cadvisor
  links:
    - influxdb:influxsrv
  command: -storage_driver=influxdb -storage_driver_db=cadvisor -storage_driver_host=influxsrv:8086
  restart: always
  ports:
    - "8080:8080"
  volumes:
    - /:/rootfs:ro
    - /var/run:/var/run:rw
    - /sys:/sys:ro
    - /var/lib/docker/:/var/lib/docker:ro
 
 grafana:
  user: "104"
  image: grafana/grafana
  user: "104"
  restart: always
  links:
    - influxdb:influxsrv
  ports:
    - "3000:3000"
  volumes:
    - grafana_data:/var/lib/grafana
  environment:
    - HTTP_USER=admin
    - HTTP_PASS=admin
    - INFLUXDB_HOST=influxsrv
    - INFLUXDB_PORT=8086
    - INFLUXDB_NAME=cadvisor
    - INFLUXDB_USER=root
    - INFLUXDB_PASS=root

```

1. `docker-compose config -q`
2. `docker-compose up -d`
3. testing CAdvisor: visit ip:8080 (first time: might be slow)
4. testing InfluxDB: visit ip:8083
5. testing Grafana: visit ip:3000 (default user/pwd = admin/admin)
   - setting refer [here](https://www.bilibili.com/video/BV1gr4y1U7CY?p=93&spm_id_from=pageDriver&vd_source=a788bdd4d7cdd9dfe02852346d523cb9)

# About [Docker Hub](https://hub.docker.com)

User repo: username/repo (use them at your own risk: they are not validated or verified in any way by Docker Inc.)
Managed by Docker Inc and by seleected vendors who provide curated base images: ubuntu  (top-level repository)

`docker login`
`docker logout`
**Tips:** The credentials will be stored in the $HOME/.dockercfg file. Since Docker 1.7.0 this is now $HOME/.docker/config.json.

# About Image

- What : light-weight software package, which include the software itself and environments and development software used.
- How to get : Hub, copied from friends, selfmake
- using UnionFs --> save file layer by layer --> avoid duplicate layer
- is made up of filesystems layered over each other. At the base is a boot filesystem, bootfs. When a container has booted, it is moved into memory, and the boot filesystem is unmounted to free up the RAM used by the initrd disk image
- next layer is a root filesystem, rootfs. This rootfs can be one or more operating systems (eg. Debian or Ubuntu filesystem). It is a read-only mode in Docker, and Docker takes advantage of a union mount to add more read-only filesystems onthe the root filesystem. Union mount is a mount that allows several filesystems to be mounted at one time but appear to be one filesystem.
- Local images live on our local Docker host in the /var/lib/docker dir. Each image will be inside a directory named for your storage driver. Containers are in /var/lib/docker/containers.

![image](https://user-images.githubusercontent.com/97860551/174950388-df0a6553-48b6-4cd1-b048-c8445bd3437e.png)

![image](https://user-images.githubusercontent.com/97860551/174952680-573cb67e-f566-4fde-96b5-85af77860cb2.png)

**Example of Base Images**

`ubuntu` // hundreds mb

`alpine` // Alpine Linux, 5mb

## 2 ways to build image

### 1. docker commit

```bash
docker run -it ubuntu /bin/bash
apt-get -yqq update
apt-get -y install apache2
exit

docker commit 4aab3ce3cb76 jamtur01/apache2
docker images jamtur01/apache2

docker commit -m "A new custom image" -a "James Turnbull" 4aab3ce3cb76 jamtur/apache2:webserver
```

### 2. Dockerfile

We recommend Dockerfile approach over docker commit because it provides a more repeatable, transparent, and idempotent mechanism for creating images.

```dockerfile
# Version: 0.0.1
FROM ubuntu:16.04
LABEL maintainer="James Turnbull james@example.com"
RUN apt-get update; apt-get install -y nginx
RUN echo 'Hi, I am in your container' \
 >/var/www/html/index.html
EXPOSE 80
```

```bash
docker build -t "jamtur01/static_web" .
docker build -t "jamtur01/static_web:v1" . // tag images 
```

> Each Instruction adds a new layer to the image and then commits the image.
>

This means that if your Dockerfile stops for some reason (for example, if an instruction fails to complete), you will be left with an image you can use. This is highly useful for debugging: you can run a container from this image interactively and then debug why your instruction failed using the last image created.

***Was it removed? I cant really get it.***

# About Volume (Directories synchronize)

![image](https://user-images.githubusercontent.com/97860551/174945930-8337041b-b5b0-45fb-839a-9aab33d5f064.png)

**Everything change in container directory reflect in local directory, vice versa.**

i) prevent the data lost when removing a container.

ii) prevent the needs to exec container to change data.  

iii) Can be shared and reused between containers

iv) A container doesn't have to be running to share its volumes.

v) changes to a volume will not be included when you update an image

vi) volumes persist until no containers use them.

- Method 1: docker run ...

  ```
  -v containerDir           #anonymous mount
  -v name:containerDir      #named mount
  -v localDir:containerDir  #path mount
  
  Q: -v juming-nginx:/etc/nginx:ro nginx <-- what is :ro? (eg rw...)
  read only. This dir cannot be modified using container, however it is allowed when using localDir.
  ```

- Method 2: DockerFile

  ```
  ###filename: dockerfile1
  
  FROM centos 
  
  VOLUME ["volume01","volume02"] # <--- anonymous mount, e.g. VOLUME ["/opt/project"]
  
  CMD echo "----end----"
  CMD /bin/bash
  ```

  to build a image : `docker build -f /..../dockerfile1 -t imageName:tag .`,
  ***there is a '.' at the end of sentence***
  
  Tips: to share these volume
      `docker run -it --name name2 --volumes-from name1 imageName:tag`

  *after remove name1, the shared volume still shared by name2.*

# About DockerFile

- What is DockerFile ? It is a shellscript, to generate an image.
- create mounts everytime when building this dockerfile generated image.
- It uses a basic DSL (Domain Specific Language)

> **Steps:**
>
> 1. write DockerFile
> 2. docker build构建镜像 == >docker build -t centosjava8:1.5 . **[!! do not ignore '.']** // . tells docker to look in the local directory to find the Dockerfile.
> 3. docker run
**Tips:** You can also specify a Git repository as a source for the Dockerfile as we see here:

```bash
docker build -t "repo/name:v1" \
 github.com/turnbullpress/docker-static_web

--no-cache : to stop caching layer
Tips: ensuring the latest packages are available to install by reset the cache when it hits ENV instruction. Thus only change the date when I want to refresh the build
FROM ubuntu:16.04
MAINTAINER James Turnbull "james@example.com"
ENV REFRESHED_AT 2016-07-01
RUN apt-get -qq update

```

**Tips:** since Docker 1.5.0 and later you can also specify a path to a file to use as a build source using the -f flag. For example, docker build -t "jamtur01/static_web" -f /path/to/file.

> If a file named .dockerignore exists in the root of the build context then it is interpreted as a newline-separated list of exclusion patterns. Much like a .gitignore file it excludes the listed files from being treated as part of the build context, and therefore prevents them from being uploaded to the Docker daemon. Globbing can be done using Go’s filepath.

> Below lists the Commands in Dockerfile

```dockerfile
#Commands: ****MUST BE CAPITAL LETTER,
#<-- is comment

FROM centos:7
MAINTAINER jacksoon<abcdefg@gmail.com> // The MAINTAINER instruction is deprecated in Docker 1.13.0.

RUN -- 2 methods, 在DOCKER BUILD时运行
-----------------------------------
RUN <shell command> // by default. Using command wrapper /bin/sh -c, if you are running If you are running the instruction on a platform without a shell or you wish to execute without a shell (for example, to avoid shell string munging), you can specify the instruction in exec format
RUN yum -y install vim
RUN ["可执行文件","参数1"，"参数2”] // exec format
RUN ["./test.php","dev","offline"] # equal RUN ./test.php dev offline

EXPOSE the current container port
---------------------------------
EXPOSE 80 // You can expose ports at run time with the docker run command with the --expose option.

WORKDIR 指定创建容器后默认登陆落脚点
--------------------------------
WORKDIR /usr/bin

USER 指定该镜像以什么样的用户去执行，如果都不指定，默认是root

Tips: you can override the working directory at runtime with the -w flag
docker run -ti -w /var/log ubuntu pwd

ENV 用来在构建镜像过程中设置环境变量
-------------------------------
ENV MY_PATH /usr/mytest
# 这个环境变量可以在后续的任何RUN指令中使用，这就如同在命令前面指定了环境变量前缀一样；也可以在其它指令中直接使用这些环境变量，
# 比如：WORKDIR $MY_PATH

In next RUN eg (RUN gem install unicorn), the command will be : MY_PATH=/usr/mytest gem install unicorn
ENV will be persisted into any containers created from your image.

Multiple env:
ENV RVM_PATH=/home/rvm RVM_ARCHFLAGS="-arch i386"

You can also pass environment variable on the docker run command line using -e flag.
e.g. docker run -it -e "WEB_PORT=8080" ubuntu env

VOLUME 容器数据卷，用于数据保存和持久化工作，-v
----------------------------------------
see 'ABOUT VOLUME'


ADD 将宿主机目录下的文件拷贝进镜像且会自动处理URL和解压tar压缩包,等于COPY+解压
----------------------------------------
e.g. ADD software.lic /opt/application/software.lic

When ADD'ing files Docker uses the ending character of the destination to determine what the source is. If the destination ends in a /, then it considers the source a directory. If it doesn't end in a /, it considers the source a file

The source of the file can also be a URL, e.g.:
ADD http://wordpress.org/latest.zip /root/wordpress.zip

Has SPECIAL MAGIC for taking care of local tar archieve. If a tar archive (gz or bz or xz) is specified as the source file, then Docker will automatically unpack it for you.
ADD latest.tar.gz /var/www/wordpress/
same behaviour with tar -x
if a file or directory with the same name already exists in the destination, it will not be overwritten
files and directory will be created with a mode 0755 and a UID and GID of 0
WARN: this will not work on source is URL
IMPORTANT: If the files or directories added by an ADD instruction change, then this will invalidate the cache for all following instructions in the Dockerfile.



COPY 类似ADD(COPY 不解压)，拷贝文件和目录到镜像中。 将从构建上下文目录中 <源路径> 的文件/目录复制到新的一层的镜像内的 <目标路径> 位置
------------------------------------------------
e.g. COPY conf.d/ /etc/apache2/
This copy files from the conf.d directory to the /etc/apache2/ directory
Source of files must be the path to a file or directory relative to the build context, the local source direcory in which your Dockerfile resides.
The destination should be an absolute path inside the container. If the destination doesnt exist, it is created along with all missing directories in its path, much like how mkdir -p command works

files and directory will be created with UID and GID of 0

 
CMD 指定容器启动后的要干的事情，在DOCKER RUN时运行
--------------------------------------------
# Dockerfile 中可以有多个 CMD 指令，但只有最后一个生效，CMD 会被 docker run 之后的参数替换，eg tomcat 本来dockerFile有CMD ["catalina.sh","run"]， 当你 docker run ... tomcat, 访问localhost:8080是没问题的，但当docker run ... tomcat /bin/bash 时就没有启动tomcat,因为覆盖了CMD ["catalina.sh","run"]
CMD <shell command>
CMD ["可执行文件","参数1"，"参数2”]
when Your Dockerfile has CMD in lastline, you can directly docker run -it repo/name, to run image. But if you do specify a command during docker run, this command will override the one in Dockerfile

ENTRYPOINT 也是用来指定一个容器启动时要运行的命令
------------------------------------------
# 类似于 CMD 指令，但是ENTRYPOINT不会被docker run后面的命令覆盖， 而且这些命令行参数会被当作参数送给 ENTRYPOINT 指令指定的程序
ENTRYPOINT ["<executeable>","<param1>","<param2>"...]
#ENTRYPOINT可以和CMD一起用，一般是变参才会使用 CMD ，这里的 CMD 等于是在给 ENTRYPOINT 传参。
#当指定了ENTRYPOINT后，CMD的含义就发生了变化，不再是直接运行其命令而是将CMD的内容作为参数传递给ENTRYPOINT指令，他两个组合
Tips: you can override the ENTRYPOINT instruction using docker run command with --entrypoint flag
Useful Binding:
ENTRYPOINT ["/usr/sbin/nginx"]
CMD ["-h"]


USER, specifies a user that the image should be run as
------------------------------------------
e.g. USER nginx
This will cause containers created from the image to be run by the nginx user. We can specify a username or a UID and group or GID. Or even a combination thereof, for example:
USER user
USER user:group
USER uid
USER uid:gid
USER user:gid
USER uid:group

You can override this at runtime using -u flag with the docker run command
TIPS: default user = root

LABEL adds metadata to a Docker Image, the metadata is in the form of key/value pairs (Introduced in Docker 1.6)
-------------------------------------------
e.g.
LABEL version="1.0"
LABEL location="New York" type="Data Center" role="Web Server"

we suggested combining all your metadata in a single LABEL instruction to save creating multiple layers with each piece of metadata.
we can inspect the labels on an image using the docker inspect command


STOPSIGNAL sets the system call signal that will be sent to the contaner when you tell it to stop. (Introduced in Docker 1.9)
--------------------------------------------
e.g. 9 (SIGKILL)

ARG defines variables that can be passed at build-time via the docker build command (Introduced in Docker 1.9)
--------------------------------------------
--build-arg flag
e.g.
ARG build
ARG webapp_user=user

the second ARG instruction sets a default, if no value is specified for the argument at build-time then the default is used.
if we run docker build --build-arg build=1234 -t jamtur01/webapp .
the build variable will be set to 1234 and the webapp_user will inherit the default value of user

WARNING: Do not pass secrets like credentials or keys using this. Your credentials will be exposed during build process and in the build history of image
Docker predefined ARG: (To use these predefined variables, pass them using the --build-arg <variable>=<value> to the docker build command)
HTTP_PROXY
http_proxy
HTTPS_PROXY
https_proxy
FTP_PROXY
ftp_proxy
NO_PROXY
no_proxy


SHELL allows the default shell used for the shell form of commands to be overriden.
--------------------------------------------
The default shell on Linux is '["/bin/sh","-c"]' and on Windows is ["cmd", "/S", "/C"]
It is useful on platforms such as Windows where there are multiple shells eg cmd or powershell
It can be used multiple times, each new SHELL overrides previous SHELL instructions, and affects any subsequent instruction


HEALTHCHECK tells Docker how to test a container to check that it is still working correctly. (Introduced in Docker 1.12)
--------------------------------------------
e.g.
HEALTHCHECK --interval=10s --timeout=1m --retries=5 CMD curl http://localhost || exit 1
contains options and then the command you wish to run itself, seperated by a CMD keyword
Only last HEALTHCHECK in Dockerfile will take effect.
You can also disable any health checks specified in any base images you may have inherited with the instruction: HEALTHCHECK NONE

--interval default 30s (run 30s after container launch and subsequently every 30s)
--timeout  default 30s (If the health check takes longer the timeout then it is deemed to have failed)
--retries  default to 3 (The number of failed checks before the container is marked as unhealthy)
The command should exit with 0 to indicate health or 1 to indicate an unhealthy state

Tips: use docker inspect --format '{{.State.Health.Status}}' contanerid,  to see the state of health
Tips: docker inspect --format '{{range .State.Health.Log}} {{.ExitCode}} {{.Output}} {{end}}'  container <<-- output from each health check


ONBUILD adds triggers to images
----------------------------------------------
Is executed when the image is used as the basis of another image (e.g. if you have an image that needs source code added from a specific location that might not yet be available, or if you need to execute a build script that is specific to the environment in which the image is built)
Tips: in docker inspect, we can file "OnBuild": ["ADD . /app/src", "RUN cd /app/src/; make"]

The triggers inserts a new instruction in build process, as it were specified right after FROM instruction.
e.g.
ONBUILD ADD . /app/src
ONBUILD RUN cd /app/src; make

Explaination
1. we build an image A with ONBUILD ADD . /var/www/
2. we build an image B FROM image A
3. we notice ADD . /var/www/ ran on docker build

NOTE: FROM, MAINTENER, ONBUILD cannot be ONBUILD To prevent inception-like recursion in Dockerfile build
Tips: Only inherited one, for example, image C FROM image B, then it will not run ADD . /var/www/ when that image is build


```

## What happens if an instruction fails?

```bash
$ cd static_web
$ sudo docker build -t="jamtur01/static_web" .
Sending build context to Docker daemon 2.56 kB
Sending build context to Docker daemon
Step 1 : FROM ubuntu:16.04
---> 8dbd9e392a96
Step 2 : MAINTAINER James Turnbull "james@example.com"
---> Running in d97e0c1cf6ea
---> 85130977028d
Step 3 : RUN apt-get update
---> Running in 85130977028d
---> 997485f46ec4
Step 4 : RUN apt-get install -y ngin
---> Running in ffca16d58fd8
Reading package lists...
Building dependency tree...
Reading state information...
E: Unable to locate package ngin
2014/06/04 18:41:11 The command [/bin/sh -c apt-get install -y
ngin] returned a non-zero code: 100
```

Let's say I want to debug this failure. I can use docker run command to create a container from the last step that succeeded in my Docker build, in this example, using image ID of 997485f46ec4

*IF named: Dockerfile, there is no need to use -f to mention it. build command will search for it automatically.*

# Push to DockerHub

- register in hub.docker.com
- login in docker using in terminal

```dockerfile
docker login [OPTIONS] [SERVER] //no server = default daemon
              -p passwd       //password [dont use this option, to use encryt in next line]
              -u username     //username 
```

- push to Dockerhub

```
docker push name:tag
docker push static_web // root repository are managed only by docker, Inc., team, thus this attempt is rejected
docker push jamtur01/static_web // ok
```

![image](https://user-images.githubusercontent.com/97860551/175017535-69165373-051f-419f-abdf-0c204ae21492.png)

NOTE: private repositories are a paid-for feature which only available to you or anyone with whom you share it.

**Tips:** If you want to delete image in Docker Hub, you need to sign in and delete it here

# Automated Builds

In Github or Bitbucket repository which contains a Dockerfile to the DockerHub. When we push to this repository, an image build will be triggered and a new image created. This was previously also known as a Trusted Build.
Tips: it also work for private Github and Bitbucket repos too

Steps:

1. Connect GitHub account or BitBucket to your Docker Hub Account:

    - Click Profile -> Create -> Create Automated Build

    - Select 'GitHub'

    - Two options: Public and Private (recommended) and Limited. Choose Public and Private, click 'Allow Access' to complete authorization

    - Select the organization and repository from which you want to construct an Automated Build

        ![image](./images/create_automated_build.png)

    - Click the Create button to add your Automated Build to the Docker Hub.

# Running your own Docker registry

- Make use of private repository on the Docker Hub

- Run your own registry behind the firewall (The team at Docker, Inc., have open-sourced the code they use t run a Docker registry, thus allowing us to build own internal registry)

**Tips:** If you're running Docker behind a proxy or corporate firewall you can also use the HTTPS_PROXY, HTTP_PROXY, NO_PROXY options to control how Docker connects.

## Running a registry from a container

1. Just run the Docker provided container:

    - `docker run -d -p 5000:5000 --name registry registry:2`

        - This will launch a container running version 2.0 of the registry application and bind port 5000 to the local host

2. Test the new registry

    - `docker tag <imageId> docker.example.com:5000/jamtur01/static_web`

    - `docker push docker.example.com:5000/jamtur01/static_web`

    - Now the image is posted in the local registry and available for us to build new containers using `docker run` command

        - `docker run -it docker.example.com:5000/jamtur01/static_web /bin/bash`

Visit [me](https://docs.docker.com/registry/deploying/) to know more on how to manage the registry, backend storage, configuring authentication

**Tips:** There are variety of other services and companies there starting to provide custom Docker registry services: e.g. Quay

# Docker Network

> To connect one container to another container. There are two ways:

- Docker's own internal network.

- From Docker 1.9 and later, using Docker Networking and the `docker network` command.

> The first method (Docker's internal network), is not an overly flexible or powerful solution.

> The Docker Networking is a more realistic method for connecting containers.

- can connect containers to each other across different hosts.

- containers connected via Docker Networking can be stopped, started or restarted without needing to update connection.

- don't need to create a container before you can connect to it. You also don't need to worry about the order in which you run containers and you get **internal container name resolution and discovery inside the network**

## Docker Internal Networking

> Every container is assigned an IP address, provided through an interface created when we installed Docker.

The interface is called **docker0**. It is a virtual Ethernet bridge that connects our **containers and local host network**.

```bash
$ ip a show docker0
4: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc
    noqueue state UP
        link/ether 06:41:69:71:00:ba brd ff:ff:ff:ff:ff:ff
        inet 172.17.42.1/16 scope global docker0
        inet6 fe80::1cb3:6eff:fee2:2df1/64 scope link
        valid_lft forever preferred_lft forever
...
```

- docker0 interface has an RFC1918 private IP address in the 172.16-172.30 range. It will default to 172.17.x.x as a subnet unless that subnet is already in used

- in this case, 172.17.42.1 will be the gateway address for the Docker network and all our Docker containers.

**Tips:** Since Docker 1.5.0 IPv6 addresses are also supported. To enable this run the Docker deamon with the --ipv6 flag.

> A series of interfaces starting with **veth**

Everytime Docker creates a new container, it creates a pair of peer interfaces that are like opposite ends of pipe. It gives one of the peers to the container to become its **eth0** interface and keeps the other peer, with a unique name like **vethec6a**

```
vethec6a Link encap:Ethernet HWaddr 86:e1:95:da:e2:5a
  inet6 addr: fe80::84e1:95ff:feda:e25a/64 Scope:Link
. . .
```

**Tips:** You can think of **veth** interface as one end of a virtual network cable. One end is plugged into the docker0 bridge, and the other end is plugged into the container.

> By binding every veth interface to the docker0 bridge, Docker creates a virtual subnet shared between the host machine and every Docker machine.

You can run your container and `ip a show docker0`, it show the IP address for our container that will be peered with a virtual interface on the host side

You can further run (in container) `traceroute google.com` to observe the first ip, is the docker0 interface gateway IP, on the host network

**Tips:** `docker inspect container` can also find IP address and gateway IP address, furthermore, you can redis-cli -h <IP address> on local Docker host instead of `redis-cli -h 127.0.0.1 -p <port>`

You can run (not in container) `sudo iptables -t nat -L -n` to observer a DNAT (destination NAT), rule that routes traffic from out container (if any such as redis) to port xxxxx on the docker host

> Two big rough edges:

- we need to hard-code the IP address of our Redis Container into our application

- restarting container: Container's IP Address change.

## Docker Networking (Introduced in 1.9 release)

> To use Docker Networks, we first need to create a network and then launch a container inside that network.

`docker network create app`

It creates a bridge network called app. A network ID is returned for the network.

**Tips:** in addition to *bridge* network, which exist on a single host, we can also create overlay network, which allow us to span **multiple hosts**.

> Then, specify the network with --net= when `docker run`

`docker run -d --net=app --name db jamtur01/redis`

**Note:** we can see a lot more information after running this command in `docker network inspect app`

> A Docker network will also add the `app` network as a domain suffix for the network, any host in the `app` network can be resolved by `hostname.app`, here `redis.app`

```bash
# in one of the container:
root:/# ping <containerName>.<netName>
```

Thus, in your redis code, you may write host: <containerName> and port 6379 e.g. `redis = Redis.new(:host => 'db', :port => '6379')

> Extra testing...

```bash
$ docker run -p 4567 --net=app --name network_test -it jamtur01/sinatra /bin/bash

# we need the dnsutils and iputils-ping packages to get the nslookup and ping binaries respectively
root: /# apt-get install -y dnsutils iputils-ping
root: /# nslookup redis # to resolve the redis container

```

![nslookup](images/nslookup.png)

A Docker network will also add the `app` network as a domain suffix for the network, any host in the `app` network can be resolved by `hostname.app`, here `redis.app`

> Connecting existing containers to the network

```bash
docker run -d --name db2 jamtur01/redis
docker network connect app db2
docker network inspect app
```

> Disconnecting a host from a network

```bash
docker network disconnect app db2
```

> Other commands

`docker network inspect app`
`docker network ls`
`docker network rm`

# Continuous integration

> Demo here is using Jenkins CI. Build a Jenkins server in Docker that runs other Docker containers.

- Refer example/ci/README.md

# Docker Orchestration

## Single container orchestration -> Docker Compose

> 可以管理多个 Docker 容器组成一个应用。你需要定义一个 YAML 格式的配置文件docker-compose.yml，写好多个容器之间的调用关系。然后，只要一个命令，就能同时启动/关闭这些容器

> Official website documentation

<https://docs.docker.com/compose/compose-file/compose-file-v3/>

### Installing Docker Compose on Linux

**Replace 1.8.0 with release number of current docker compose**

```bash
sudo curl -L https://github.com/docker/compose/releases/download/1.8.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose
```

### Uninstall Docker Compose

uninstall `rm /usr/local/bin/docker-compose`

### docker-compose.yml

```yaml
image: jamtur01/composeapp

command: python app.py

ports:
  - "5000:5000"

volumes:
  - .:/composeapp

links:
  - redis

# build instruction and provide the path to a Dockerfile 
# to have Compose build an image and then create services from it
build: /home/james/compose app 
```

### command

```bash
docker-compose -h                           # 查看帮助
docker-compose up                           # 启动所有docker-compose服务
docker-compose up -d                        # 启动所有docker-compose服务并后台运行
docker-compose down                         # 停止并删除容器、网络、卷、镜像。
docker-compose exec  yml里面的服务id                 # 进入容器实例内部  docker-compose exec docker-compose.yml文件中写的服务id /bin/bash
docker-compose ps                      # 展示当前docker-compose编排过的运行的所有容器
docker-compose top                     # 展示当前docker-compose编排过的容器进程
 
docker-compose logs  yml里面的服务id     # 查看容器输出日志
dokcer-compose config     # 检查配置
dokcer-compose config -q  # 检查配置，有问题才有输出
docker-compose restart   # 重启服务
docker-compose start     # 启动服务
docker-compose rm        # remove services
docker-compose stop      # 停止服务
docker-compose kill # if docker-compose stop did not work...

```

### comparison

mysql > redis > jar

#### without docker compose

```bash
docker exec -it mysql57 /bin/bash
mysql -uroot -p
create database db2021;
use db2021;
CREATE TABLE `t_user` (...),

docker run  -p 6379:6379 --name redis608 --privileged=true -v /app/redis/redis.conf:/etc/redis/redis.conf -v /app/redis/data:/data -d redis:6.0.8 redis-server /etc/redis/redis.conf

docker run -d -p 6001:6001 myjarservice1:1.6

```

**Be noted that in myjarservice1, the application.properties host is fixed ip**

> application.properties

```properties
server.port=6001
# ========================alibaba.druid相关配置=====================
spring.datasource.type=com.alibaba.druid.pool.DruidDataSource
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
spring.datasource.url=jdbc:mysql://192.168.111.169:3306/db2021?useUnicode=true&characterEncoding=utf-8&useSSL=false
spring.datasource.username=root
spring.datasource.password=123456
spring.datasource.druid.test-while-idle=false
# ========================redis相关配置=====================
spring.redis.database=0
spring.redis.host=192.168.111.169
spring.redis.port=6379
spring.redis.password=
spring.redis.lettuce.pool.max-active=8
spring.redis.lettuce.pool.max-wait=-1ms
spring.redis.lettuce.pool.max-idle=8
spring.redis.lettuce.pool.min-idle=0
# ========================mybatis相关配置===================
mybatis.mapper-locations=classpath:mapper/*.xml
mybatis.type-aliases-package=com.atguigu.docker.entities
# ========================swagger=====================
spring.swagger2.enabled=true

```

#### using docker compose

> application.properties

```properties
server.port=6001

# ========================alibaba.druid相关配置=====================
spring.datasource.type=com.alibaba.druid.pool.DruidDataSource
spring.datasource.driver-class-name=com.mysql.jdbc.Driver
#spring.datasource.url=jdbc:mysql://192.168.111.169:3306/db2021?useUnicode=true&characterEncoding=utf-8&useSSL=false
spring.datasource.url=jdbc:mysql://mysql:3306/db2021?useUnicode=true&characterEncoding=utf-8&useSSL=false
spring.datasource.username=root
spring.datasource.password=123456
spring.datasource.druid.test-while-idle=false

# ========================redis相关配置=====================
spring.redis.database=0
#spring.redis.host=192.168.111.169
spring.redis.host=redis
spring.redis.port=6379
spring.redis.password=
spring.redis.lettuce.pool.max-active=8
spring.redis.lettuce.pool.max-wait=-1ms
spring.redis.lettuce.pool.max-idle=8
spring.redis.lettuce.pool.min-idle=0

# ========================mybatis相关配置===================
mybatis.mapper-locations=classpath:mapper/*.xml
mybatis.type-aliases-package=com.atguigu.docker.entities

# ========================swagger=====================
spring.swagger2.enabled=true

```

> docker-compose.yml

```yaml
version: "3"
 
services:
  microService:
    image: zzyy_docker:1.6
    container_name: ms01 ## --name ms01
    ports: # -v 6001:6001
      - "6001:6001" 
    volumes:
      - /app/microService:/data
    networks:  # --network atguigu_net
      - atguigu_net 
    depends_on: 
      - redis
      - mysql
 
  redis: # default name: mydocker_redis
    image: redis:6.0.8
    ports:
      - "6379:6379"
    volumes:
      - /app/redis/redis.conf:/etc/redis/redis.conf
      - /app/redis/data:/data
    networks: 
      - atguigu_net
    command: redis-server /etc/redis/redis.conf
 
  mysql: # default name: mydocker_mysql
    image: mysql:5.7
    environment:
      MYSQL_ROOT_PASSWORD: '123456'
      MYSQL_ALLOW_EMPTY_PASSWORD: 'no'
      MYSQL_DATABASE: 'db2021'
      MYSQL_USER: 'zzyy'
      MYSQL_PASSWORD: 'zzyy123'
    ports:
       - "3306:3306"
    volumes:
       - /app/mysql/db:/var/lib/mysql
       - /app/mysql/conf/my.cnf:/etc/my.cnf
       - /app/mysql/init:/docker-entrypoint-initdb.d
    networks:
      - atguigu_net
    command: --default-authentication-plugin=mysql_native_password #解决外部无法访问
 
networks: 
   atguigu_net: # 会创建bridge # default name: mydocker_atguigu_net

```

`docker-compose config -q` # 检查配置，有问题才有输出

`docker-compose up -d`

`docker-compose stop`

**Tips:** 10多20个的时候选择k8s


## Distributed service discovery -> Consul

> Consul is a specialized datastore that uses consensus algorithms.
>
> Services can register them selves with Consul and share that registration information
in a highly available and distributed manner

> It is also interesting because it provides:

- A service catelog with an API instead of the traditional key=value store
of most service discovery tools

- Both a DNS-based query interface through an inbuilt DNS server and a HTTP based
REST API to query the information. The choice of interfaces, especially the DNS-based interface,
allows you to easily drop Consul into your existing environment.

- Service monitoring AKA health checks. Consul has powerful service monitoring built into the tool.

## Orchestration and clustering of Docker -> Swarm

> Skipped for now

# Using the Docker API

## Three Specific APIs in the Docker ecosystem.

- The Registry API - provides integration with the Docker registry, which stores our images.

- The Docker Hub API - provides integration with the Docker Hub.

- The Docker Remote API - provides integration with the Docker daemon.

### The Docker Remote API

> Available client libraries. e.g Go, Ruby ...

![clientlibrary](https://docs.docker.com/engine/api/sdk/examples/)

> It is provided by the Docker Daemon

By default, docker daemons binds to a socket: unix:///var/run/docker.sock, on the host on which it is running

The daemon runs with `root` privileges so as to have access needed to manage the appropriate resources. 
If a group named `docker` exists on your system, Docker will apply ownership of the socket to that group.
Hence, any user that belongs to the `docker` group can run Docker without needing `root` privileges.

> If we want remote access to the API, we need to bind the Docker daemon to a network interface.

This is done by passing or adjusting the -H flag to the Docker daemon.

Tips: Querying the Docker API locally

```bash
$ curl --unix-socket /var/run/docker.sock http:/info
{"ID":"PH4R:BT7H:44F6:GQGP:FS2O:7OZO:HQ2P:NSVF:MK27:NBGZ:N3VP:
  K2O5","Containers":3,"ContainersRunning":3,"ContainersPaused"
  :0,"ContainersStopped":0,"Images":3,"
  . . .
}
```

> On most distributions, we can bind the Docker daemon to a network interface by editing the daemon's startup configuration files

- Ubuntu / Debian
    - Older: /etc/default/docker

    - Release with Upstart: /etc/init/docker.conf

    - systemd release: /lib/systemd/system/docker.service

- Red Hat, Fedora, and related distro: /etc/sysconfig/docker

    - systemd release: /usr/lib/systemd/system/docker.service

```
ExecStart=/usr/bin/dockerd --selinux-enabled

to

ExecStart=/usr/bin/dockerd --selinux-enabled -H tcp://0.0.0.0:2375

```

Reload and restart the daemon: `sudo systemctl --system daemon-reload`

> Connect to a remote Docker daemon

`sudo docker -H docker.example.com:2375 info`

- -H to specify host

**Tips:** docker will also honor the `DOCKER_HOST` environment variable rather than requiring the continued use of the -H flag

```bash
$ export DOCKER_HOST="tcp://docker.example.com:2375"
```

> Using the info API endpoint

`$ curl http://docker.example.com:2375/info | python3 -mjson.tool`

```bash
# list images
$ curl http://docker.example.com:2375/images/json | python 3 -mjson.tool

# inspect image
$ curl http://docker.example.com:2375/images/b608dbb10e2564f5bd0eef045bf297e56b1149edc70bece54fef4b217261a473/json | python 3 -mjson.tool

# search the images in docker hub
$ curl "http://docker.example.com:2375/images/search?term=jamtur01" | python3 -mjson.tool

# list running container (docker ps)
$ curl "http://docker.example.com:2375/containers/json" | python3 -mjson.tool
# list running and stopped container (docker ps)
$ curl "http://docker.example.com:2375/containers/json?all=1" | python3 -mjson.tool
# show container
$ curl http://docker.example.com:2375/containers/591ba02d8d149e5ae5ec2ea30ffe85ed47558b9a40b7405e3b71553d9e59bed3/json | python3 -mjson.tool

# create container
$ curl -X POST -H "Content-Type: application/json" \
http://docker.example.com:2375/containers/create \
-d '{
    "Image": "jamtur01/jekyll"
}'
$ curl -X POST -H "Content-Type: application/json" \
http://docker.example.com:2375/containers/create?name=jekyll \
-d '{
    "Image": "jamtur01/jekyll",
    "Hostname": "jekyll"
}'

# start the container, it same as docker run jamtur01/jekyll
$ curl -X POST -H "Content-Type: application/json" \
http://docker.example.com:2375/containers/591ba02d8d149e5ae5ec2ea30ffe85ed47558b9a40b7405e3b71553d9e59bed3/start \
-d '{
    "PublishAllPorts":true
}'
```

#### Authenticating the Docker Remote API

> The authentication uses TLS/SSL certificates to secure your connection to the API.

There are couple ways too, including using a full PKI infrastructure, either creating our own Certificate Authority (CA) or using an existing CA.

**Example of creating our own Certificate Authority(CA)**

##### Creating a CA

```bash
$ which openssl

# Create a CA directory
$ sudo mkdir /etc/docker
$ cd /etc/docker

# generating a private key
$ echo 01 | sudo tee ca.srl
# we'll specify a passphrase for CA key, we'll need it to create and sign certificates with our new CA.
$ sudo openssl genrsa -des3 -out ca-key.pem

# creating a CA certificate
# we'll need this later to verify our secure connection.
$ sudo openssl req -new -x509 -days 365 -key ca-key.pem -out ca.pem
# Enter pass phrase

```

##### Creating a Server Certicate

```bash
# Creating a server certificate signing request and key
# We can use our new CA to sign and validate a certificate signing request or CSR
# and key for our Docker server.

# create key
$ sudo openssl genrsa -des3 -out server-key.pem
# Enter passphrase..

# create server certificate signing request (CSR).
# The most important option here is `Common Name` or CN. 
# This should either be the FQDN (full qualified domain name) of the Docker server
# (i.e., what is resolved to in DNS; e.g. docker.example.com) or *. Which allow us to use the server certificate on any server.
$ sudo openssl req -new -key server-key.pem -out server.csr
# folks connect to our host via IP address, so we need to configure for that (replace x.x.x.x with the IP address of Docker daemon)
$ echo subjectAltName = IP:x.x.x.x,IP:127.0.0.1 > extfile.cnf
# sign our CSR and generate our server certificate.
$ sudo openssl x509 -req -days 365 -in server.csr -CA ca.pem \ 
-CAkey ca-key.pem -out server-cert.pem -extfile extfile.cnf

# let's strip out the passphrase from our server key. We can't enter one when the Docker daemon starts, so we need to remove it
# (FROM ENCRYPTED PRIVATE KEY to PRIVATE KEY)
$ sudo openssl rsa -in server-key.pem -out server-key.pem

# let's add some tighter permissions to the files to better protect them.
$ sudo chmod 0600 /etc/docker/server-key.pem /etc/docker/server-cert.pem \
/etc/docker/ca-key.pem /etc/docker/ca.pem

# configuring the docker daemon
# --tlsverify enable TLS
# you can use --tls flag to enable TLS, but not client-side authentication
# 2375 should be only used for unauthenticated connections. 2376 is the default TLS/SSL port for Docker.
$ vi ...
ExecStart=/usr/bin/docker -d -H tcp://0.0.0.0:2376 --tlsverify --tlscacert=/etc/docker/ca.pem --tlscert=/etc/docker/server-cert.pem --tlskey=/etc/docker/server-key.pem

$ sudo systemctl --system daemon-reload
```

##### Creating a client certificate and key

```bash
$ sudo openssl genrsa -des3 -out client-key.pem

# create a client CSR
# The most important option here is `Common Name` or CN. 
# This should either be the FQDN (full qualified domain name) of the Docker server
# (i.e., what is resolved to in DNS; e.g. docker.example.com) or *. Which allow us to use the server certificate on any server.
$ sudo openssl req -new -key client-key.pem -out client.csr

# enable client authentication for our key by adding some extended SSL attributes.
$ echo extendedKeyUsage = clientAuth > extfile.cnf

# sign CSR
$ sudo openssl x509 -req -days 365 -in client.csr -CA ca.pem \
-CAkey ca-key.pem -out client-cert.pem -extfile extfile.cnf

# let's strip out the passphrase from our client key. To allow us to use it with the Docker client
# (FROM ENCRYPTED PRIVATE KEY to PRIVATE KEY)
$ sudo openssl rsa -in client-key.pem -out client-key.pem

# let's install them into .docker directory. This is the default location where Docker will look for
# certificates and keys. Docker will specifically look for
# key.pem, cert.pem, and ca.pem
$ mkdir -p ~/.docker/
$ cp ca.pem ~/.docker/ca.pem
$ cp client-key.pem ~/.docker/key.pem
$ cp client-cert.pem ~/.docker/cert.pem
$ chmod 0600 ~/.docker/key.pem ~/.docker/cert.pem

# test the connection to the Docker daemon from the client
# we don't need to specify any certificate or key files, because Docker 
# has automatically looked these up in our /.docker/
# if we did need to specify these files, we could with the --tlscacert, --tlscert, and --tlskey
$ sudo docker -H=docker.example.com:2376 --tlsverify info

```

# Getting Help and Extending Docker

## Getting help

### The Docker forums

![click me](https://forums.docker.com)

### Docker on IRC

![#docker and #docker-dev, on the Freenode IRC network](http://freenode.net)

The #docker channel is generally for user help and general Docker issues, whereas
#docker-dev is where contributors to Docker’s source code gather.

### Docker on Github

Docker (and most of its components and ecosystem) is hosted on GitHub. The
principal repository for Docker itself is https://github.com/docker/docker/.

### Reporting Issues

When you log a ticket, please remember to include the following background information:

- The output of docker info and docker version.

- The output of uname -a.

- Your operating system and version (e.g., Ubuntu 16.04).

Then provide a detailed explanation of your problem and the steps others can take
to reproduce it.

If you’re logging a feature request, carefully explain what you want and how you
propose it might work. Think carefully about generic use cases: is your feature
something that will make life easier for just you or for everyone?

Please take a moment to check that an issue doesn’t already exist documenting
your bug report or feature request. If it does, you can add a quick ”+1” or ”I have
this problem too”, or if you feel your input expands on the proposed implementation
or bug fix, then add a more substantive update.

### Submitting a pull request

If you’re happy with your documentation update, bug fix, or new feature, you
can submit a pull request for it on GitHub. To do so, you should fork the Docker
repository and make changes on your fork in a feature branch:

- If it is a bug fix branch, name it XXXX-something, where XXXX is the number
of the issue.

- If it is a feature branch, create a feature issue to announce your intentions,
and name it XXXX-something, where XXXX is the number of the issue.

Any pull request with a feature in it should include updates to the documentation.
You should use the process above to test your documentation changes before you
submit your pull request.

# Extra knowledge

> 亿级缓存算法

- 哈希取余分区, hash(key) % N个机器台数，计算出哈希值，用来决定数据映射到哪一个节点上

- 一致性哈希环, Object 顺时钟找到node，就算一个node死了会顺时钟找到活着的node。所以缺点就是数据倾斜。可是完成了扩展性和容错性

  ![image-20220911101054665](images/image-20220911101054665.png)

- **hash slot**

- Backup volume

  - ```bash
    docker run --rm --volumes-from james_blog \
    -v $(pwd):/backup ubuntu \
    tar cvf /backup/james_blog_backup.tar /var/www/html
    ```

# Problem faced

> no ifconfig

```bash
$ apt-get update
$ apt-get install net-tools
```

> cannot open directory. : permission denied

```bash
--privileged=true
```

> Docker - mysql cannot display chinese

```bash
Why? Show variables like 'character%' ==> latin1 instead of utf8
```

# Other Knowledge Earned

> Curl command

refer [here](https://curl.se/docs/manual.html)

```bash
Download file
-------------
-o --output-file e.g. curl -o thatpage.html http://www.example.com/
-O --remote-name e.g. curl -O http://www.example.com/index.html
   download multiple files  curl -O www.haxx.se/index.html -O curl.se/download.html

Request
-------
-- Password
--- Default Basic auth
curl -u name:passwd http://machine.domain/full/path/to/file 
or
curl http://name:passwd@machine.domain/full/path/to/file

-- POST
-d data, must be urlencoded or use --data-urlencode
   e.g. curl -d "name=Rafael%20Sagula&phone=3320780" http://www.where.com/guest.cgi
--data-urlencode
   e.g. curl --data-urlencode "name=Rafael Sagula&phone=3320780" http://www.where.com/guest.cgi 
-F form, if contents read from files, use @. Below example, the field name coolfiles is used to send three files
   e.g. curl -F "coolfiles=@fil1.gif;type=image/gif,fil2.txt ,fil3.html" http://www.post.com/postit.cgi
   tips: If the content-type is not specified, curl will try to guess from the file extension (it only knows a few), or use the previously specified type (from an earlier file if several files are specified in a list) or else it will use the default type application/octet-stream.
   e.g. curl -F "file=@cooltext.txt" -F "yourname=Daniel" -F "filedescription=Cool text file with cool text inside" http://www.post.com/postit.cgi

-X --request <method> e.g. curl -X DELETE https://...
-e referer e.g. curl -e www.coolsite.com http://www.showme.com/
-A user-agent e.g. curl -A 'Mozilla/3.0 (Win95; I)' http://www.nationsbank.com/
-b <data|filename> e.g. curl -b "name=Daniel" www.sillypage.com
--dump-header use previously received cookies in following sessions. If you get cookies from a server and store them in a file in a manner similar to: 
   step1: curl --dump-header headers www.example.com
   step2: curl -b headers.txt www.example.com
While saving headers to a file is a working way to store cookies, it is however error-prone and not the preferred way to do this. Instead, make curl save the incoming cookies using the well-known Netscape cookie format like this:
  curl -c cookies.txt www.example.com
  curl -b cookies.txt -c cookies.txt www.example.com // this to read and write from Netscape cookie file
-H extra headers
    e.g. curl -H "X-you-and-me: yes" www.love.com
-i response with headers
-E certificate
    e.g. curl -E /path/to/cert.pem:password https://secure.site.com/
--tlv1.0 Many older HTTPS servers have problems with specific SSL or TLS versions, which newer versions of OpenSSL etc use, therefore it is sometimes useful to specify what TLS version curl should use. Otherwise it use a sensible TLS default version.
    e.g. curl --tlv1.0 https://secure.site.com/
-C --continue-at <offset> 
    e.g. continue download curl -C - -o file ftp://ftp.server.com/path/file
    e.g. continue upload curl -C - -T file ftp://ftp.server.com/path/file
    e.g. continue download from web curl -C - -o file http://www.server.com/
```

> Wget

```bash
wget -q --show-progress https://code.jquery.com/jquery-3.6.0.min.js

--show-progress
-q --quiet
-O --output-document
```

> Apt-get for Dockerfile

- iproute2 # ip command

- traceroute # traceroute command `traceroute google.com`

- dnsutils # nslookup

- iputils-ping # ping


> awk

- print second column with line have 'inet '

    `| awk '/inet / {print $2}'`

- print second column with line have 'inet ', columns seperated by ':'

    `| awk -F ':' '/inet / {print $2}'`
