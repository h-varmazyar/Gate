import os
import google.generativeai as genai
from .provider import Provider

# A token is equivalent to about 4 characters for Gemini models
class GeminiProvider(Provider):
    def send_request(self, posts):
        genai.configure(api_key=provider.token)
        # Choose a model that's appropriate for your use case.
        model = genai.GenerativeModel('gemini-1.5-flash')

        prompt = "Write a story about a magic backpack."

        response = model.generate_content(prompt)

        print(response.text)
