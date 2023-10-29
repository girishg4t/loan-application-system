FROM golang:1.21-bullseye as build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN mkdir target
RUN go build -o ./target/loan-application-backend main.go

FROM ubuntu
RUN apt-get update
RUN apt-get install -y ca-certificates && \
    apt-get clean
WORKDIR /app
COPY --from=build /app/target/* ./ /app/.env ./

EXPOSE 8080

ENTRYPOINT ["./loan-application-backend"]
