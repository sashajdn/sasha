CREATE KEYSPACE IF NOT EXISTS openai WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'sasha-infra-cassandra-dc1': '3'
} AND durable_writes = true;
