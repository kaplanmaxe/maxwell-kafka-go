# maxwell-kafka-go

Example implementation of:

- mariadb
- maxwell
- kafka
- zookeeper
- golang producer inserting record every 3 seconds
- golang consumer reading from kafka

## Running

```
docker-composer up -d
docker logs --follow consumer
```