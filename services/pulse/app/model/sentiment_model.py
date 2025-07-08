# import joblib
# from app.config import MODEL_PATH
# from app.preprocess.cleaner import preprocess
#
# class SentimentModel:
#     def __init__(self):
#         self.model = joblib.load(MODEL_PATH)
#
#     def predict(self, text):
#         cleaned_text = preprocess(text)
#         return self.model.predict([cleaned_text])[0]


import torch
from transformers import AutoTokenizer, AutoModelForSequenceClassification

# بارگذاری مدل و توکنایزر
model_path = "./../model/parsbert_model"
tokenizer = AutoTokenizer.from_pretrained(model_path)
model = AutoModelForSequenceClassification.from_pretrained(model_path)
model.eval()  # غیرفعال‌سازی dropout

# پیش‌بینی احساسات
def predict_sentiment(text: str) -> int:
    inputs = tokenizer(text, return_tensors="pt", truncation=True, max_length=512, padding="max_length")
    with torch.no_grad():
        outputs = model(**inputs)
        prediction = torch.argmax(outputs.logits, dim=1).item()
    return prediction  # 0 = منفی, 1 = خنثی, 2 = مثبت
