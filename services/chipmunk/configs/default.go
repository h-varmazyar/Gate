package configs

var (
	DefaultConfig = []byte(`
service_name: "chipmunk"
version: "v1.0.0"
grpc_port: 11000
amqp_configs:
  connection: "amqp://rabbitmq:rabbitmq@localhost"
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
  host: "localhost"
  port: 5432
  name: "chipmunk"
  is_ssl_enable: false
`)

	DefaultEnv = []byte(`
SERVICE_NAME="chipmunk"
VERSION="v2.0.0"
GRPC_PORT="11000"
AMQP_CONFIGS_CONNECTION="amqp://rabbitmq:rabbitmq@localhost:5672"
BUFFER_CONFIGS_CANDLE_BUFFER_LENGTH="400"
MARKETS_APP_SERVICE_CONFIGS_NETWORK_ADDRESS="network.gate.svc:13000"
MARKETS_APP_SERVICE_CONFIGS_EAGLE_ADDRESS="eagle.gate.svc:12000"
MARKETS_APP_SERVICE_CONFIGS_CORE_ADDRESS="core.gate.svc:10100"
MARKETS_APP_WORKER_CONFIGS_CORE_ADDRESS="core.gate.svc:10100"
MARKETS_APP_SERVICE_CONFIGS_MARKET_STATISTICS_WORKER_INTERVAL="5ms"
CANDLES_APP_SERVICE_CONFIGS_CORE_ADDRESS="core.gate.svc:10100"
CANDLES_APP_SERVICE_CONFIGS_EAGLE_ADDRESS="eagle.gate.svc:12000"
CANDLES_APP_WORKER_CONFIGS_CORE_ADDRESS="core.gate.svc:10100"
CANDLES_APP_WORKER_CONFIGS_PRIMARY_DATA_QUEUE="chipmunk_ohlc"
CANDLES_APP_WORKER_CONFIGS_CONSUMER_COUNT="10"
CANDLES_APP_WORKER_CONFIGS_MISSED_CANDLES_INTERVAL="10m"
CANDLES_APP_WORKER_CONFIGS_LAST_CANDLES_INTERVAL="3s"
CANDLES_APP_WORKER_CONFIGS_REDUNDANT_REMOVER_INTERVAL="15m"
WALLETS_APP_SERVICE_CONFIGS_CORE_ADDRESS="core.gate.svc:10100"
WALLETS_APP_WORKER_CONFIGS_CORE_ADDRESS="core.gate.svc:10100"
WALLETS_APP_WORKER_CONFIGS_WALLET_WORKER_INTERVAL="10s"
DB.TYPE="postgreSQL"
DB.USERNAME="postgres"
DB.PASSWORD="NAc2PdbLVH15nDdC3zXL7HyY1Ozbnzxb"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="chipmunk"
DB_IS_SSL_ENABLE="false"
`)
)
