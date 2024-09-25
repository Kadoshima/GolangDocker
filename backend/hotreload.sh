#!/bin/bash

# コンテナ名を変数にする
CONTAINER_NAME="golangdocker-backend-1"

# コンテナの状態をチェックして、立ち上がっていれば再起動する
if [ $(docker ps -q -f name=$CONTAINER_NAME) ]; then
    echo "Restarting container: $CONTAINER_NAME"
    docker restart $CONTAINER_NAME
    echo "Container restarted successfully."
else
    echo "Container $CONTAINER_NAME is not running."
fi
