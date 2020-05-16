#!/bin/bash

echo "I am going to run the image 'ja_chat_service_v1' for you on port 8080"
echo "Press Enter.... and run the client under client folder/main.go"
read nothing

docker run -it -p 8080:8080 ja_chat_service_v1
