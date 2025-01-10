package configs

var (
	DefaultConfig = []byte(`
service_name: "chipmunk"
version: "v1.0.0"
grpc_port: 11100
amqp_configs:
  connection: "amqp://rabbitmq:rabbitmq@localhost:5672"
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
    data_warmup_mood: false
    data_correction_mode: true
    normal_data_gathering: false
posts_app:
  workers_configs:
    running: false
    network_address: "localhost:10101"
    sahamyab_post_collector_url: "https://www.sahamyab.com/guest/twiter/list?v=0.1"
    max_sentiment_detector_token_length:
    sentiment_detector_token:
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
  host: "localhost"
  port: 5432
  name: "chipmunk"
  is_ssl_enable: false
`)
)
