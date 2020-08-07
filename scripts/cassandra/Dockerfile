FROM cassandra

ENV CASSANDRA_KEYSPACE hydromon
COPY entrypoint-wrap.sh /entrypoint-wrap.sh

ENTRYPOINT ["/entrypoint-wrap.sh"]
CMD ["cassandra", "-f"]
