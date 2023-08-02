FROM golang:1.20.6-alpine3.17
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 3000
CMD [ "/app/main" ]