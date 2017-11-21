docker network create -d bridge roachnet

docker run -d \
--name=roach1 \
--hostname=roach1 \
--net=roachnet \
-p 26257:26257 -p 8080:8080  \
-v "$/tmp/cockroach-data/roach1:/cockroach/cockroach-data"  \
cockroachdb/cockroach:v1.1.2 start --insecure

cockroach user set maxroach --insecure
cockroach sql --insecure -e 'CREATE DATABASE urlrepo'
cockroach sql --insecure -e 'GRANT ALL ON DATABASE bank TO maxroach'

