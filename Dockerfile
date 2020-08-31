# Dockerfile
FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN rm .env
RUN mv .env.deploy .env
RUN go build -o task-api .
CMD ["/app/task-api"]