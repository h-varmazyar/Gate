package configs

var (
	DefaultConfig = []byte(`
service_name: "raven"
version: "v1.1.0"
http_port: 8585
host: "localhost"
api_external_Address: "localhost:8585"
chipmunk_router:
  chipmunk_address: "localhost:11100"
core_router:
  core_address: "localhost:10100"
eagle_router:
  eagle_address: "localhost:12000"
telegram_bot_router:
  telegram_bot_address: "localhost:14000"
network_router:
  network_address: "localhost:13000"
`)
)
