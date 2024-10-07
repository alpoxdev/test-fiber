# 빌드 스테이지
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# 의존성 파일 복사 및 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 실행 스테이지
FROM alpine:latest

WORKDIR /root/

# 빌드 스테이지에서 빌드된 실행 파일 복사
COPY --from=builder /app/main .

# 필요한 경우 SSL 인증서 추가
RUN apk --no-cache add ca-certificates

# 애플리케이션 실행
CMD ["./main"]


