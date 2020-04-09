# docker build -t projecta:v1.0 -f cmd/projectA/Dockerfile .
# Đứng ở ngoài cùng chạy lệnh build image ở trên

# build image dựa trên golang:1.14.1 tên build_base
FROM golang:1.14.1 AS build_base
# chạy dòng lệnh: mkdir -p /gobuild/bins tức là tạo 1 thư mục /gobuild/bins trong image build_base
# mkdir là tạo thư mục, -p là tạo folder nếu nó chưa tồn tại thì tạo luôn
RUN mkdir -p /gobuild/bins
# WORKDIR /gobuild là chỉ định thư mục làm việc là thư mục /gobuild mới tạo ở trên
WORKDIR /gobuild
# ENV là môi trường
ENV GO111MODULE=on
# Copy file go.mod vào thư mục làm việc, là thư mục /gobuild ở cái WORKDIR
COPY go.mod .
# download các lib cho go module
RUN go mod download

# build image từ build_base, tên là server_builder
FROM build_base AS server_builder
# Tương tự, tạo thư mục /gobuild/bins trong image server_builder
RUN mkdir -p /gobuild/bins
# Chỉ định thư mục làm việc là /gobuild vừa mới tạo ở trên
WORKDIR /gobuild
# Copy tất cả source vào /gobuild
COPY ./ /gobuild/
# Build cái file main.go, lúc này đường dẫn sẽ là cmd/projectA/main.go
# go build [-o output] [-i] [build flags] [packages]
# Tức là build file cmd/projectA/main.go (-i là input) thành binary vào /bins/binarypath (-o là output)
RUN go build -o bins/binarypath -i cmd/projectA/main.go

# Từ image ubuntu:18.04
FROM ubuntu:18.04
# Install/update một số thứ cần thiết
RUN apt update && apt-get install -y ca-certificates
# Copy từ image server_builder ở trên, thư mục /gobuild/bins/binarypath vào thư mục /binarypath
COPY --from=server_builder /gobuild/bins/binarypath /binarypath
# Tạo điểm vào là /binarypath
ENTRYPOINT [ "/binarypath" ]