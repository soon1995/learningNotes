```bash
$ docker build -t jamtur01/composeapp .

# run interactively
$ docker-compose up

# run daemonized
$ docker-compose up -d
```

Compose has created two new services: `composeapp_redis_1` and `composeapp_web_1`

