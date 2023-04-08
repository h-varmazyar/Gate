package configs

var (
	DefaultConfig = []byte(`
service_name: "core"
version: "v1.0.0"
grpc_port: 10100
amqp_configs:
  connection: "amqp://guest:guest@localhost"
coinex_configs:
  coinex_callback_queue: "coinex_callback"
  chipmunk_ohlc_queue: "chipmunk_ohlc"
brokerages_app:
  service_configs:
    chipmunk_grpc_address: ":11000"
    eagle_grpc_address: ":12000"
platforms_app:
  service_configs:
    chipmunk_grpc_address: ":11000"
functions_app:
  service_configs:
    network_grpc_address: ":13000"
    coinex:
      coinex_callback_queue: "coinex_callback"
      coinex_public_rate_limiter_id: "0dee8939-0160-4609-8fc9-de75d884086c"
      coinex_spot_api_rate_limiter_id: "523f39ce-c433-4aac-815f-0cf511a475d8"
      chipmunk_ohlc_queue: "chipmunk_ohlc"
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5433
  name: "core"
  is_ssl_enable: false
`)
)
