import os
from dotenv import load_dotenv
from pathlib import Path

env_path = Path(__file__).resolve().parent.parent / "configs" / ".env"
load_dotenv(dotenv_path=env_path)

MODEL_PATH = os.getenv("MODEL_PATH", "app/model/saved_model.pkl")
LOG_LEVEL = os.getenv("LOG_LEVEL", "info")
PORT = int(os.getenv("PORT", 11000))
