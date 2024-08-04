import psycopg2
# from config.configs import AppConfigs.DatabaseConfig

def connect(configs):
    conn = psycopg2.connect(user = configs.user,
    password = configs.password,
    host = configs.host,
    port = configs.port,
    database = configs.name)
    return conn
