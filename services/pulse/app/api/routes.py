from fastapi import APIRouter, HTTPException
from pydantic import BaseModel
from app.model.sentiment_model import SentimentModel

router = APIRouter()
model = SentimentModel()

class TextIn(BaseModel):
    text: str

@router.post("/predict")
def predict_sentiment(data: TextIn):
    try:
        prediction = model.predict(data.text)
        return {"sentiment": prediction}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
