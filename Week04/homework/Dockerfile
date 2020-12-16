# FROM registry.cn-hangzhou.aliyuncs.com/artifact/golang:1.11-alpine as builder

# WORKDIR /go/src/artifact/common

# # COPY . .

# # RUN go get github.com/Masterminds/glide \
# # 	# && glide install \
# # 	&& CGO_ENABLED=0 make build
# RUN CGO_ENABLED=0 make build

FROM registry.cn-hangzhou.aliyuncs.com/artifact/alpine:3.7

WORKDIR /wwwroot

# COPY --from=builder /go/src/artifact/common/build/app ./build/app
# COPY --from=builder /go/src/artifact/common/config ./config

COPY ./build/app ./build/app
COPY ./config ./config

RUN chown -R app:app ./
RUN ls -aslh /wwwroot

EXPOSE 8000

# USER app

ENTRYPOINT ["/wwwroot/build/app"]
