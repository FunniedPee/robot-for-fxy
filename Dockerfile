FROM golang:1.22.6 AS Builder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

ARG app=robot-for-fxy

WORKDIR /build

COPY . .
RUN go mod tidy && go build -o robot-for-fxy main.go

FROM scratch
LABEL authors="funniedpee"

COPY --from=Builder /build/robot-for-fxy /robot-for-fxy

CMD ["/robot-for-fxy"]


EXPOSE 8080