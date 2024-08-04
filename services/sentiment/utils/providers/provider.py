import time

class Provider:
    def __init__(self, name, token, tokens_per_minute, requests_per_day, tokens_per_day):
        self.name = name
        self.token = token
        self.requests_per_minute = requests_per_minute
        self.tokens_per_minutes = tokens_per_minutes
        self.requests_per_day = requests_per_day
        self.tokens_per_day = tokens_per_day

        self.last_reset_time = time.time()

    def reset_quota(self):
        if time.time() - self.last_reset_time >= self.quota_reset_interval:
            self.quota = self.max_quota
            self.last_reset_time = time.time()

    def can_handle(self):
        self.reset_quota()
        return self.quota > 0

    def handle_request(self, posts):
        if not self.can_handle():
            return False

        response = self.send_request(posts)
        if response:
            self.quota -= 1
            return response
        return False

    def send_request(self, posts):
        raise NotImplementedError("Subclasses should implement this method")

def create_providers(configs):
    from .gemini import GeminiProvider

    providers = []
    providers.append(GeminiProvider(name='gemini-1.5-pro', request_per_minute=2, tokens_per_minute=32000,requests_per_day=50))
    providers.append(GeminiProvider(name='gemini-1.5-flash', request_per_minute=15, tokens_per_minute=1000000,requests_per_day=1500))
    providers.append(GeminiProvider(name='gemini-1.0-pro', request_per_minute=15, tokens_per_minute=32000,requests_per_day=1500))

    return providers
