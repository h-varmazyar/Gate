package configs

var (
	DefaultConfig = []byte(`
service_name: "network"
version: "v1.1.1"
grpc_port: 13000
db:
  postgres_dsn: "postgresql://postgres:postgres@localhost/network?sslmode=disable"
amqp_configs:
  connection: "amqp://rabbitmq:rabbitmq@localhost"
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5432
  name: "network"
  is_ssl_enable: false
`)
)
