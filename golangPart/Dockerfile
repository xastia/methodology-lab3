
# Збираємо виконуваний файл за допомогою образу golang
FROM golang:1.17 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/fizzbuzz


# Використовуємо образ distroless для виконання додатка
FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/fizzbuzz /app/fizzbuzz

CMD ["/app/fizzbuzz"]

