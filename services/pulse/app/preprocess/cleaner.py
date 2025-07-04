import re
import emoji
from hazm import Normalizer, word_tokenize, stopwords_list, Stemmer

normalizer = Normalizer()
stemmer = Stemmer()
stopwords = set(stopwords_list())

def clean_text(text):
    text = emoji.replace_emoji(text, replace='')
    text = re.sub(r"http\S+|www\S+|https\S+", '', text)
    text = re.sub(r"@[A-Za-z0-9_]+", '', text)
    text = re.sub(r"#", '', text)
    text = re.sub(r"[^\w\s\u0600-\u06FF]", '', text)
    text = re.sub(r"\d+", '', text)
    text = re.sub(r"\s+", ' ', text).strip()
    return text

def preprocess(text):
    text = clean_text(text)
    text = normalizer.normalize(text)
    tokens = word_tokenize(text)
    tokens = [stemmer.stem(token) for token in tokens if token not in stopwords]
    return " ".join(tokens)
