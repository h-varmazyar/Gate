package configs

var (
	DefaultConfig = []byte(`
service_name: telegram bot
version: v1.1.0
grpc_port: 14000
service_configs:
  bot_configs:
    debug_mode: true
    bot_token: 1327918034:AAEDnz0mz89uxv6TznlM4CMXe-IAyjpsUg0
  handler_configs:
    brokerage_address: :10000
    chipmunk_address: :11000
db:
  type: "postgreSQL"
  username: "postgres"
  password: "postgres"
  host: "localhost"
  port: 5432
  name: "telegramBot"
  is_ssl_enable: false
`)
)
