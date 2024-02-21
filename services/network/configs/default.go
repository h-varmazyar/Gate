package configs

var (
	DefaultConfig = []byte(`
service_name: "network"
version: "v1.1.1"
grpc_port: 13000
amqp_configs:
  connection: "amqp://rabbitmq:rabbitmq@localhost:5672"
db:
  type: "postgreSQL"
  username: "postgres"
  password: "e1M9nEw33FfNs8U"
  host: "localhost"
  port: 5433
  name: "network"
  is_ssl_enable: false
`)
)
