import time
import requests
from providers.provider import create_providers
from config.configs import provider_configs

class Tagger:
    def __init__(self, configs):
        self.chipmunkAddress = configs['chipmunkAddress']
        self.providers = create_providers(provider_configs)

    def get_posts(self):
        response = requests.get('{self.chipmunkAddress}/v1/posts/non-polarity')
        if response.status_code == 200:
            return response.json()
        return []

    def send_polarity(self, posts, response):
        data = {'posts': posts, 'response': response}
        requests.post('{self.chipmunkAddress}/v1/posts/polarity', json=data)

    def start():
        while True:
            posts = get_posts()
            if not posts:
                time.sleep(5)
                continue

            for provider in providers:
                if provider.can_handle():
                    response = provider.handle_request(posts)
                    if response:
                        send_polarity(posts, response)
                        break
            else:
                print("No available provider, waiting...")
                time.sleep(10)