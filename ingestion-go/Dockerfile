# Multi-stage Dockerfile for Go application
# Multi-stage build, build aşamasındaki gereksiz dosyaları final image'a dahil etmez
# Bu sayede çok daha küçük ve güvenli bir production image elde ederiz

# ============================================
# Stage 1: Builder Stage
# ============================================
# Bu stage'de Go kodunu derleyeceğiz
FROM golang:1.23-alpine AS builder

# Build sırasında gereken araçları yükle
RUN apk add --no-cache git

# Çalışma dizinini ayarla
WORKDIR /app

# go.mod ve go.sum dosyalarını kopyala
# Bu dosyalar önce kopyalanır çünkü bağımlılıklar değişmediği sürece
# Docker cache'i kullanılabilir (layer caching optimizasyonu)
COPY go.mod go.sum ./

# Bağımlılıkları indir (cache'lenir)
RUN go mod download

# Tüm kaynak kodunu kopyala
COPY . .

# Go kodunu derle
# -o server: Çıktı dosyasının adı "server" olacak
# . : Mevcut dizindeki tüm Go kodunu derle
# CGO_ENABLED=0: CGO'yu devre dışı bırak (static binary, daha küçük)
# GOOS=linux: Linux için derle
# GOARCH=amd64: AMD64 mimarisi için derle
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# ============================================
# Stage 2: Final Stage (Runtime)
# ============================================
# Bu stage'de sadece derlenmiş binary'yi alacağız
# Alpine Linux çok küçük bir base image'dir (yaklaşık 5MB)
FROM alpine:latest

# Güvenlik için non-root user oluştur
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Çalışma dizinini ayarla
WORKDIR /app

# COPY --from=builder açıklaması:
# Bu komut, önceki stage'den (builder) dosya kopyalar.
# --from=builder: builder stage'inden kopyalama yapılacağını belirtir
# /app/server: builder stage'indeki derlenmiş binary dosyası
# /app/server: final stage'deki hedef konum
# Bu sayede sadece derlenmiş binary final image'a eklenir,
# Go toolchain, kaynak kodlar ve build araçları eklenmez.
# Sonuç: Çok daha küçük bir image (alpine ~5MB + binary ~10-20MB = ~15-25MB toplam)
COPY --from=builder /app/server /app/server

# Non-root user'a dosya sahipliğini ver
RUN chown -R appuser:appuser /app

# Non-root user olarak çalış
USER appuser

# Uygulamanın çalışacağı portu belirt
EXPOSE 3000

# Health check endpoint'i (opsiyonel ama önerilir)
# Alpine'de wget yok, bu yüzden curl kullanıyoruz veya basit bir check yapıyoruz
# Eğer curl kullanmak isterseniz: RUN apk add --no-cache curl
# HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
#   CMD wget --no-verbose --tries=1 --spider http://localhost:3000/health || exit 1

# Uygulamayı başlat
# Alpine'de shell yok, bu yüzden exec form kullanıyoruz
CMD ["./server"]

