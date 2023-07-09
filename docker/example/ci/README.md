# CI

## Single

```bash
sudo mkdir -p /var/jenkins_home
cd /var/jenkins_home
# UID = 1000
sudo chown -R 1000 /var/jenkins_home

docker build -t jamtur01/jenkins .
# /var/jenkins_home contain Jenkin's configuration data and allow us to perpetuate its state across container launches.
# /var/run/docker.sock, is the socket for Docker's daemon into Docker container, This will allow us to run Docker containers from inside our Jenkins container
# !!! This is a security risk. By binding the Docker socket inside the Jenkins container, you give the container access to the underlying Docker Host.
# I recommend you only do this if you are comfortable that the Jenkins container, any other containers on that Docker host are at a comparable security level
docker run -d -p 8080:8080 -p 50000:50000 \
-v /var/jenkins_home:/var/jenkins_home \
-v /var/run/docker.sock:/var/run/docker.sock \
--name jenkins \
jamtur01/jenkins

# stream log until meet: Jenkins is fully up and running
# you may find your initial admin password here too (or in /var/jenkins_home/secrets/initialAdminPassword)
docker logs jenkins -f

connect to localhost:8080

```

After that, we start a new job in localhost:8080:

1. Under advanced, use custom workspace : /var/jenkins_home/jobs/${JOB_NAME}/workspace 

2. Under Source control, add a repo : https://github.com/turnbullpress/docker-jenkins-sample

3. Under Build Step, exec a script: 

```bash
# Build the image to be used for this job.
IMAGE=$(sudo docker build . | tail -1 | awk '{ print $NF }')
# Build the directory to be mounted into Docker.
# WORKSPACE is a variable created by Jenkins, holding the workspace dir we defined in
# our job earlier
MNT="$WORKSPACE/.."
# Execute the build inside Docker.
# rake spec runs RSpec tests
CONTAINER=$(sudo docker run -d -v $MNT:/opt/project/ $IMAGE /bin/
bash -c 'cd /opt/project/workspace; rake spec')
# Attach to the container so that we can see the output.
sudo docker attach $CONTAINER
# Get its exit code as soon as the container stops.
# block until the command the container is executing finishes
RC=$(sudo docker wait $CONTAINER)
# Delete the container we've just used.
sudo docker rm $CONTAINER
# Exit with the same value as that with which the process exited.
exit $RC
```
What does this script do?

It create a new Docker image using Dockerfile contained in Git repository we've just specified

Below is Dockerfile in Git repo
```Dockerfile
FROM ubuntu:16.04
MAINTAINER James Turnbull "james@example.com"
ENV REFRESHED_AT 2016-06-01
RUN apt-get update
RUN apt-get -y install ruby rake
RUN gem install --no-rdoc --no-ri rspec ci_reporter_rspec
```
Here we’re building an Ubuntu host, installing Ruby and RubyGems, and then
installing two gems: rspec and ci_reporter_rspec. This will build an image that
we can test using a typical Ruby-based application that relies on the RSpec test
framework. The ci_reporter_rspec gem allows RSpec output to be converted
to JUnit-formatted XML that Jenkins can consume. 


4. Add post-build action -> publish JUnit test result.

    - In `Test report XMLs`, specify: `spec/reports/*.xml`, this is the location of the ci_reporter gem's XML output.

5. Save

*Jenkins Tips:*

- Build now to build now

- Build History show history of builds

- The first time the test run is slow because Docker is building new image. The next time would be faster as Docker have the required image prepared.

- Click Console Output to see the commands that have been executed

- we can also check the `Test Result link` for the uploaded JUnit test results if required.

- we can automate Jenkins job further by enabling SCM polling, which triggers automatic builds when new commits are made to the repository. Similar automation can be achieved with a post-commit hook or via a GitHub or Bitbucket repository hook

- [parameterized builds](https://wiki.jenkins-ci.org/display/JENKINS/Parameterized+Build) can be used to make shell script step more generic to suit multiple frameworks and languages


## Multi-configuration Jenkins

> If we wanted to test our application on multiple platforms e.g. Ubuntu, Debian, and CentOS

1. New Item

2. Select `Multi-configuration project`

3. git: https://github.com/turnbullpress/docker-jenkins-sample.git.

4. Add axis -> user-defined axis -> Name: OS -> Values:

```
centos
debian
ubuntu
```
5. Build Environment -> Delete workspace before build starts. It cleans up our build environment by deleting the checked-out repository prior to initiating a new set of jobs.

6. Shell:

```bash
# Build the image to be used for this run.
cd $OS; IMAGE=$(sudo docker build . | tail -1 | awk '{ print $NF }')
# Build the directory to be mounted into Docker.
MNT="$WORKSPACE/.."
# Execute the build inside Docker.
CONTAINER=$(sudo docker run -d -v "$MNT:/opt/project" $IMAGE /bin
/bash -c "cd /opt/project/$OS; rake spec")
# Attach to the container's streams so that we can see the output
.
sudo docker attach $CONTAINER
# As soon as the process exits, get its return value.
RC=$(sudo docker wait $CONTAINER)
# Delete the container we've just used.
sudo docker rm $CONTAINER
# Exit with the same value as that with which the process exited.
exit $RC
```

We change OS every time we're executing a job. Each different OS has different Dockerfile
e.g. CentOS use yum to install

7. Publish JUnit test result report -> spec/reports/*.xml

8. Save

9. Build Now

