package configs

var (
	DefaultConfig = []byte(`
service_name: "core"
version: "v1.0.0"
grpc_port: 10101
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5432
  name: "indicator"
  is_ssl_enable: false
nats_url: "nats://localhost:4222"
`)
)
