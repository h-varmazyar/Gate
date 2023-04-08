package configs

var (
	DefaultConfig = []byte(`
service_name: "eagle"
version: "v1.0.0"
grpc_port: 12000
strategies_app:
  service_configs:
    automated:
      redis_address: "localhost:6379"
      core_address: ":10100"
      chipmunk_address: ":11000"
      telegram_bot_address: ":14000"
      broadcast_channel_id: 1001263706636
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5433
  name: "eagle"
  is_ssl_enable: false
`)
)
