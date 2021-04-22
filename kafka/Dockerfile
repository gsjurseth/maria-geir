FROM bitnami/kafka:2

COPY wait-for-it.sh /wait-for-it.sh
COPY runit.sh /runit.sh

ENTRYPOINT ["/wait-for-it.sh","-h","zookeeper.local","-p","2181","--","/runit.sh"]
