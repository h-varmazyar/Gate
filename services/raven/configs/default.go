package configs

var (
	DefaultConfig = []byte(`
service_name: "raven"
version: "v1.1.0"
http_port: 8585
host: "localhost"
api_external_Address: "localhost:8585"
chipmunk_router:
  chipmunk_address: ":11000"
core_router:
  core_address: ":10100"
eagle_router:
  eagle_address: ":12000"
telegram_bot_router:
  telegram_bot_address: ":14000"
network_router:
  network_address: ":13000"
`)
)
