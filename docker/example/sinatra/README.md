```bash
$ wget --cut-dirs=3 -nH -r --reject Dockerfile,index.html --noparent http://dockerbook.com/code/5/sinatra/webapp/
docker build -t jamtur01/sinatra .
docker run -d -p 4567 --name webapp -v $PWD/webapp:/opt/webapp jamtur01/sinatra

# Creating a redis database
docker build -t jamtur01/redis . # run this in redis dir
docker run -d -p 6379 --name redis jamtur01/redis
# To connect this instance, we need to install redis client, by redis-tools package on Ubuntu
sudo apt-get -y install redis-tools
# or On Red Hat related distributions.
sudo yum install -y -q redis
# then
redis-cli -h 127.0.0.1 -p <port>
# use quit to quit client

docker rm -f redis
docker rm -f webapp

docker network create app
docker run -d --net=app --name redis jamtur01/redis
docker run -d --net=app --name webapp_redis -v $PWD/webapp:/opt/webapp jamtur01/sinatra
curl -i -H 'Accept: application/json' -d 'name=Foo&status=Bar' http://localhost:49162/json
curl -i http://localhost:49162/json

```
