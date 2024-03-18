```sh
docker pull jupyter/pyspark-notebook
docker run -it --rm -p 8888:8888 -v $PWD:/home/jovyan jupyter/pyspark-notebook
```

Refs:
- https://medium.com/@bee811101/%E4%BD%BF%E7%94%A8-docker-%E5%BF%AB%E9%80%9F%E5%BB%BA%E7%BD%AE-pyspark-%E7%92%B0%E5%A2%83-657f9d8bff3a
