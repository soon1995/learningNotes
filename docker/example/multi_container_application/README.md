# A Multi-container Application Stack

> Dockerizing a Node.js application that make use of Express framework with a Redis back end

> Multi-containers:

- A Node container to serve our Node application

    - nodejs/nodeapp/

        - `wget https://raw.githubusercontent.com/jamtur01/dockerbook-code/master/code/6/node/nodejs/nodeapp/package.json`

        - `wget https://raw.githubusercontent.com/jamtur01/dockerbook-code/master/code/6/node/nodejs/nodeapp/server.js`

    - nodejs => docker build -t jamtur01/nodejs .

- A Redis primary container to hold and cluster our state

    - redis_base => docker build -t jamtur01/redis .

    - redis_primary => docker build -t jamtur01/redis_primary .

- Two Redis replica containers to cluster our state.

    - redis_replica => docker build -t jamtur01/redis_replica .

- A logging container to capture our application logs.


> Creating Redis back-end cluster

```bash
docker network create express

docker run -d -h redis_primary --net express --name redis_primary jamtur01/redis_primary

docker logs redis_primary # nothing output, our redis server is logging to a file rather than to stdout

docker run -it --rm --volumes-from redis_primary ubuntu cat /var/log/redis/redis-server.log

# run replica
docker run -d -h redis_replica1 --name redis_replica1 --net express jamtur01/redis_replica

docker run -it --rm --volumes-from redis_replica1 ubuntu cat /var/log/redis/redis-replica.log

# run replica2
docker run -d -h redis_replica2 --name redis_replica2 --net express jamtur01/redis_replica

docker run -it --rm --volumes-from redis_replica2 ubuntu cat /var/log/redis/redis-replica.log
```

> Create Node Container

```bash
docker run -d --name nodeapp -p 3000:3000 --net express jamtur01/nodejs
```

> LogStash Server

**logstash.conf**

Monitors two files. Logstash will watch these files and send any new data inside of them
into Logstash.

`output` takes any events Logstash receives and outputs them to standard out. In a real word
Logstash configuration we would output to an Elasticsearch cluster or other destination.

```
input {
  file {
    type => "syslog"
    path => ["/var/log/nodeapp/nodeapp.log", "/var/log/redis/redis-server.log"]
  }
}

output {
  stdout {
    codec => rubydebug
  }
}
```

```bash
docker build -t jamtur01/logstash .

docker run -d --name logstash \
--volumes-from redis_primary \
--volumes-from nodeapp \
jamtur01/logstash
```
