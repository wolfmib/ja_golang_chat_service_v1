## Setup yourown Go

---

- install go (google by yourself... 😅)
- go into your project folder
- type 'go mod init'
- install proto:
  - go get -u github.com/golang/protobuf/protoc-gen-go
- source ja_create_golang_env.sh
  - it adds your GOPATH into PATH
- write your own proto code under proto folder 📂

  - add the following inside the proto file
  - [📜]: option go_package = "github.com/wolfmib/ja_golang_chat_service_v1/proto";
    - your github or gitlab repository/project_name/proto

- use the following command:
  - protoc -I proto proto/chat.proto --go_out=plugins=grpc:proto

---

## Docker build 🐟

---

- under root_folder
- Build images: ja_chat_service_v1 ✅

  - docker build --tag=ja_chat_service_v1 .

-
