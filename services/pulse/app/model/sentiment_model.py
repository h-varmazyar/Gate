import joblib
from app.config import MODEL_PATH
from app.preprocess.cleaner import preprocess

class SentimentModel:
    def __init__(self):
        self.model = joblib.load(MODEL_PATH)

    def predict(self, text):
        cleaned_text = preprocess(text)
        return self.model.predict([cleaned_text])[0]
