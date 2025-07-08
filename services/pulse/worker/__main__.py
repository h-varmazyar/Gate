import asyncio
from worker.sentiment_worker import run_worker

if __name__ == "__main__":
    try:
        asyncio.run(run_worker())
    except KeyboardInterrupt:
        print("✋ Worker stopped by user.")