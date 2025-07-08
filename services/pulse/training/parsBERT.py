import pandas as pd
from sklearn.model_selection import train_test_split
from datasets import Dataset
from transformers import AutoTokenizer, AutoModelForSequenceClassification, Trainer, TrainingArguments
from transformers import DataCollatorWithPadding
import torch
import gc

# 1. خواندن داده
df = pd.read_csv("./assets/data/sahamyab_posts.csv")

df["sentiment"] = pd.to_numeric(df["sentiment"], errors="coerce")
df.dropna(subset=['sentiment'], inplace=True)

# 2. برچسب‌ها را نرمال کنیم: مقادیر بین -1 تا 1 → به 3 کلاس (منفی، خنثی، مثبت)
def labelize(value):

    if value < -0.2:
        return 0  # منفی
    elif value > 0.2:
        return 2  # مثبت
    else:
        return 1  # خنثی

df["label"] = df["sentiment"].apply(labelize)

# 3. تقسیم داده
train_df, test_df = train_test_split(df, test_size=0.2, random_state=42)

# 4. تبدیل به Dataset
train_ds = Dataset.from_pandas(train_df[["content", "label"]])
test_ds = Dataset.from_pandas(test_df[["content", "label"]])

# 5. بارگذاری tokenizer
tokenizer = AutoTokenizer.from_pretrained("HooshvareLab/bert-base-parsbert-uncased")

def tokenize_function(example):
    return tokenizer(
    example["content"],
     truncation=True,
     padding='max_length',
     max_length=512
     )

train_ds = train_ds.map(tokenize_function, batched=True)
test_ds = test_ds.map(tokenize_function, batched=True)

# 6. بارگذاری مدل
model = AutoModelForSequenceClassification.from_pretrained(
    "HooshvareLab/bert-base-parsbert-uncased",
    num_labels=3
)

# 7. تنظیمات آموز
gc.collect()


training_args = TrainingArguments(
    output_dir="./parsbert-checkpoints",
    eval_strategy="epoch",
    save_strategy="epoch",
    per_device_train_batch_size=4,
    per_device_eval_batch_size=4,
    gradient_accumulation_steps=4,
    num_train_epochs=3,
    logging_dir="./logs",
    logging_steps=100,
    load_best_model_at_end=True,
    metric_for_best_model="accuracy",
    save_total_limit=1,
    fp16=False,
)


# 8. collator برای padding خودکار
data_collator = DataCollatorWithPadding(tokenizer=tokenizer)

# 9. ارزیابی
from sklearn.metrics import accuracy_score, f1_score

def compute_metrics(eval_pred):
    logits, labels = eval_pred
    preds = logits.argmax(axis=-1)
    return {
        "accuracy": accuracy_score(labels, preds),
        "f1": f1_score(labels, preds, average="macro")
    }

# 10. آموزش
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=train_ds,
    eval_dataset=test_ds,
    tokenizer=tokenizer,
    data_collator=data_collator,
    compute_metrics=compute_metrics
)

trainer.train()

# 11. ذخیره مدل
model.save_pretrained("app/model/parsbert_model")
tokenizer.save_pretrained("app/model/parsbert_model")
