# etl_worker/Dockerfile
FROM python:3.12-slim

# Bağımlılıklar
RUN apt-get update && apt-get install -y libpq-dev gcc && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Bağımlılıkları kur
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Kodları kopyala
COPY . .

# Uygulamayı başlat
CMD ["python", "main.py"]