#!/usr/bin/env bash
echo "Preparing directories"
INPUT_DIR=./input
OUTPUT_DIR=./output
rm -rf "${OUTPUT_DIR}"
mkdir -p "${INPUT_DIR}"
mkdir -p "${OUTPUT_DIR}"
echo "Generating API V1 client"
curl https://docs.jumpcloud.com/api/1.0/index.yaml -o "${INPUT_DIR}/swagger_v1.yaml"
docker-compose run --rm swagger-codegen generate -i /swagger-api/yaml/swagger_v1.yaml -l go -c /config/config_v1.json -o /swagger-api/out/v1

echo "Generating API V2 client"
curl https://docs.jumpcloud.com/api/2.0/index.yaml -o "${INPUT_DIR}/swagger_v2.yaml"
docker-compose run --rm swagger-codegen generate -i /swagger-api/yaml/swagger_v2.yaml -l go -c /config/config_v2.json -o /swagger-api/out/v2

echo "Client code has been generated in ${OUTPUT_DIR}"
