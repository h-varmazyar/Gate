from fastapi import FastAPI
from app.api.schemas import TextInput
from app.model.parsbert_model import predict_sentiment

app = FastAPI()

@app.get("/")
def root():
    return {"message": "Sentiment API is running!"}

@app.post("/predict")
def predict(input: TextInput):
    print(input.text)
    sentiment = predict_sentiment(input.text)
    return {"sentiment": sentiment}
