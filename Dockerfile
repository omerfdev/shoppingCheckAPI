# Resmi Golang görüntüsünü kullanın
FROM golang:latest

# Çalışma dizinini belirtin
WORKDIR /app

# Proje dosyalarını kopyalayın
COPY . .

# Bağımlılıkları çözün ve uygulamayı derleyin
RUN go mod download
RUN go build -o main .

# Uygulamayı çalıştırmak için komut
CMD ["./main"]
