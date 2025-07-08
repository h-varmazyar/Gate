from transformers import AutoTokenizer, AutoModelForSequenceClassification
import torch
import os

# مسیر مدل ذخیره‌شده
MODEL_DIR = os.path.join(os.path.dirname(__file__), "parsbert_model")

# بارگذاری tokenizer و مدل
tokenizer = AutoTokenizer.from_pretrained(MODEL_DIR)
model = AutoModelForSequenceClassification.from_pretrained(MODEL_DIR)
model.eval()  # غیرفعال کردن dropout

# تعیین دستگاه اجرایی (GPU یا CPU)
device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
model = model.to(device)

# تابع پیش‌بینی احساس
def predict_sentiment(text: str) -> int:
    inputs = tokenizer(text, return_tensors="pt", truncation=True, max_length=512, padding="max_length")
    inputs = {k: v.to(device) for k, v in inputs.items()}

    with torch.no_grad():
        outputs = model(**inputs)
        prediction = torch.argmax(outputs.logits, dim=1).item()
    return prediction  # مثلاً: 0 = منفی، 1 = خنثی، 2 = مثبت
