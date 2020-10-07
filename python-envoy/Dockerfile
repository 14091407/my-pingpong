FROM grpc/python:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ONBUILD COPY requirements.txt /usr/src/app/
ONBUILD RUN pip3 install --no-cache-dir -r requirements.txt

COPY . .
RUN python3 server.py