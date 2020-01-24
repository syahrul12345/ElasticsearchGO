FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 5555
ENV db_user="syahrul"
ENV db_pass="password"
ENV db_name="acepointerproject"
ENV db_host="localhost"
ENV token_password="123"
CMD ["./main"]