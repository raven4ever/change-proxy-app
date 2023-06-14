FROM python:3.11 as builder

WORKDIR /app

COPY src/requirements.txt .

RUN pip install -r requirements.txt
RUN pip install pyinstaller

COPY src/ ./

RUN pyinstaller --onefile --clean --name feynman main.py