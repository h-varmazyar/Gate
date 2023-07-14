package configs

var (
	DefaultConfig = []byte(`
service_name: "chipmunk"
version: "v1.0.0"
grpc_port: 11000
amqp_configs:
  connection: "amqp://rabbitmq:rabbitmq@185.110.191.66"
buffer_configs:
  candle_buffer_length: 400
markets_app:
  service_configs:
    network_address: ":13000"
    eagle_address: ":12000"
    core_address: ":10100"
  worker_configs:
    core_address: ":10100"
    market_statistics_worker_interval: "5ms"
candles_app:
  service_configs:
    core_address: ":10100"
    eagle_address: ":12000"
  worker_configs:
    core_address: ":10100"
    primary_data_queue: "chipmunk_ohlc"
    consumer_count: 10
    missed_candles_interval: "10m"
    last_candles_interval: "3s"
    redundant_remover_interval: "15m"
indicators_app:
  service_configs:
resolutions_app:
  service_configs:
wallets_app:
  service_configs:
    core_address: ":10100"
  buffer_configs:
  worker_configs:
    core_address: ":10100"
    wallet_worker_interval: "10s"
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "185.110.191.66"
  port: 5432
  name: "chipmunk"
  is_ssl_enable: false
`)
)
