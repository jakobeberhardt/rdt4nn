
FROM python:3.11-slim

ENV PYTHONUNBUFFERED=1

RUN apt-get update \
	&& apt-get install -y --no-install-recommends \
		libgomp1 \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY main.py /app/main.py

RUN pip install --no-cache-dir numpy xgboost

CMD ["python", "-u", "/app/main.py"]
