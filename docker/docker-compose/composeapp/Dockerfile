# Compose Sample appication image
FROM python:2.7
LABEL maintainer="James Turnbull james@example.com"
ENV REFRESHED_AT 2016-06-01

ADD . /composeapp

WORKDIR /composeapp

RUN pip install -r requirements.txt
