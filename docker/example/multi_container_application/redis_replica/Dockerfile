FROM jamtur01/redis
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-01

# tells any containers built from this image that they are a replica of the redis_primary
# host and should attempt replication on port 6379
ENTRYPOINT [ "redis-server", "--protected-mode no", "--logfile /var/log/redis/redis-replica.log", "--slaveof redis_primary 6379" ]
