import asyncio
import json
from nats.aio.client import Client as NATS
from app.model.sentiment_model import predict_sentiment
from app.preprocess.cleaner import clean_text
import os

NATS_URL = os.getenv("NATS_URL", "nats://localhost:4222")
INPUT_SUBJECT = os.getenv("NATS_SUB_INPUT", "tweets.raw")
OUTPUT_SUBJECT = os.getenv("NATS_SUB_OUTPUT", "tweets.analyzed")

async def message_handler(msg):
    data = json.loads(msg.data.decode())
    tweet_id = data.get("id")
    text = data.get("content")

    if not text:
        return  # یا log warning

    cleaned = clean_text(text)
    sentiment = predict_sentiment(cleaned)

    result = {
        "id": tweet_id,
        "content": text,
        "sentiment": sentiment,
    }

    await msg._client.publish(OUTPUT_SUBJECT, json.dumps(result).encode())
    print(f"[✔] Tweet analyzed → ID: {tweet_id} | Sentiment: {sentiment}")

async def run_worker():
    nc = NATS()
    await nc.connect(servers=[NATS_URL])

    await nc.subscribe(INPUT_SUBJECT, cb=message_handler)
    print(f"🔄 Listening on NATS subject: {INPUT_SUBJECT}")

    while True:
        await asyncio.sleep(1)
