import database
from config.configs import AppConfigs
from database import db
# import uvicorn
# from configparser import ConfigParser

# def config(filename='config/config.ini', section='api'):
#     parser = ConfigParser()
#     parser.read(filename)
#
#     api_config = {}
#     if parser.has_section(section):
#         params = parser.items(section)
#         for param in params:
#             api_config[param[0]] = param[1]
#     else:
#         raise Exception(f'Section {section} not found in the {filename} file')
#
#     return api_config

if __name__ == "__main__":
    conf = AppConfigs('./config/default.ini')

    db = db.connect(conf.database)
#     api_config = config()
#     uvicorn.run("api.server:app", host=api_config['host'], port=int(api_config['port']), reload=True)
