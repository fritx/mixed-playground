# Import the necessary modules
from pyspark.sql import SparkSession
from pyspark.sql.functions import *

# Create a SparkSession
spark = SparkSession.builder \
    .appName("My App")

# fix: Exception in thread "main" java.lang.IllegalStateException:
# # Cannot retrieve files with 'spark' scheme without an active SparkEnv.
# https://stackoverflow.com/questions/77622514/cannot-submit-spark-application-to-docker-compose
# spark = spark.config("spark.executorEnv.PYSPARK_PYTHON", "/opt/bitnami/python/bin/python") \

spark = spark.getOrCreate()
rdd = spark.sparkContext.parallelize(range(1, 100))

print("THE SUM IS HERE:", rdd.sum())

# Stop the SparkSession
spark.stop()
