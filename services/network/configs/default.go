package configs

var (
	DefaultConfig = []byte(`
service_name: "network"
version: "v1.1.1"
grpc_port: 13000
amqp_configs:
  connection: "amqp://rabbitmq:rabbitmq@192.168.100.16"
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "192.168.100.16"
  port: 5432
  name: "network"
  is_ssl_enable: false
`)
)
