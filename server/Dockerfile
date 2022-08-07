FROM python:3.9.13-slim-buster

WORKDIR /usr/src/app

COPY . .
RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "./server.py" ]