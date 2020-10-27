# GoBasha_oauth-api
Oauth For GoBasha
docker run --name cassandra bitnami/cassandra:latest

Ports 7000 92

CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor' : 1};
use oauth ;
create table access_tokens(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint) ;
