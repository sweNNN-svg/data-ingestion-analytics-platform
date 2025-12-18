# 1. Base image
FROM python:3.12-slim

# 2. Çalışma dizini
WORKDIR /app

# 3. Bağımlılıkların kopyalanması (Cache avantajı için önce requirements.txt)
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# 4. Kalan tüm kodun kopyalanması
COPY . .

# 5. API'nin çalışacağı portun açılması
EXPOSE 8000

# 6. Uygulamanın başlatılması
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]