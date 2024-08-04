from configparser import ConfigParser

class AppConfigs:
    def __init__(self, filename='./config/default.ini'):
        config = ConfigParser()
        config.read(filename)

        self.database = self.DatabaseConfig(config['database'])
        self.api = self.APIConfig(config['api'])
#         self.logging = self.LoggingConfig(config['logging'])

    class DatabaseConfig:
        def __init__(self, db_config):
            self.host = db_config['host']
            self.port = db_config['port']
            self.user = db_config['user']
            self.password = db_config['password']
            self.name = db_config['name']

    class APIConfig:
        def __init__(self, api_config):
            self.host = api_config['host']
            self.port = api_config['port']

    class LoggingConfig:
        def __init__(self, logging_config):
            self.level = logging_config['level']
            self.file = logging_config['file']
