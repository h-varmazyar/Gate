#!/bin/bash

source venv/bin/activate

echo "✅ Updating requirements.txt from current environment..."
pip freeze > requirements.txt
echo "✅ requirements.txt updated successfully!"
