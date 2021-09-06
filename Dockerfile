FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD app /app/app

CMD ["./app"]