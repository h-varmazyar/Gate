package configs

var (
	DefaultConfig = []byte(`
service_name: "eagle"
version: "v1.0.0"
grpc_port: 12000
strategies_app:
  service_configs:
    automated:
      redis_address: "localhost:6380"
      core_address: "localhost:10100"
      chipmunk_address: "localhost:11000"
      telegram_bot_address: "localhost:14000"
      broadcast_channel_id: 1001263706636
  automated_worker:
    redis_address: "localhost:6380"
    core_address: "localhost:10100"
    chipmunk_address: "localhost:11000"
    telegram_bot_address: "localhost:14000"
    broadcast_channel_id: -100
  chipmunk_address: "localhost:11000"
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5432
  name: "eagle"
  is_ssl_enable: false
`)
)
