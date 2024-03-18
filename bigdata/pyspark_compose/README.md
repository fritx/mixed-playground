```sh
docker compose up

docker compose logs spark-master
# >> Starting Spark master at spark://xxx.xx.x.x:7077
SPARK_MASTER_IP=xxx.xx.x.x

docker compose cp -L sum.py spark-master:/opt/bitnami/spark/sum.py
docker compose exec spark-master spark-submit --master spark://$SPARK_MASTER_IP:7077 sum.py
# >> DAGScheduler: Job 0 finished: sum at /opt/bitnami/spark/sum.py:17, took 6.407227 s
# THE SUM IS HERE: 4950

docker compose cp -L pi.py spark-master:/opt/bitnami/spark/pi.py
docker compose exec spark-master spark-submit --master spark://$SPARK_MASTER_IP:7077 pi.py
# >> DAGScheduler: Job 0 finished: reduce at /opt/bitnami/spark/pi.py:24, took 6.588161 s
# Pi is roughly 3.134560
```

Refs:
- https://stackoverflow.com/questions/36054871/spark-executorenv-doesnt-seem-to-take-any-effect
- https://stackoverflow.com/questions/77622514/cannot-submit-spark-application-to-docker-compose
- https://medium.com/@mehmood9501/using-apache-spark-docker-containers-to-run-pyspark-programs-using-spark-submit-afd6da480e0f
