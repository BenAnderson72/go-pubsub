FROM golang:1.17-alpine

ARG SERVICE

WORKDIR /app

COPY . .

RUN rm -Rf app/google-cloud-sdk
RUN rm -Rf app/google-cloud-cli-454.0.0-linux-x86_64.tar.gz
RUN rm -Rf app/notes.txt

RUN go build -o main cmd/${SERVICE}/main.go

CMD ["./main"]