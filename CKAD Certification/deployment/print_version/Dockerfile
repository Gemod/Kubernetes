# Usa un'immagine Python leggera
FROM python:3.11-slim


RUN apt update && apt install -y procps && rm -rf /var/lib/apt/lists/*
# Imposta la working directory
WORKDIR /app

# Copia i file di progetto
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY main.py .

# Espone la porta di FastAPI
EXPOSE 8080

# Comando di avvio
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080"]
