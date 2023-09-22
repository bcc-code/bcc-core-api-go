# # Run the openapi-generator Docker image to generate the client SDK
# docker run --rm \
#   -v "./swagger.json:/local/openapi.json" \
#   -v "./models2:/local/out" \
#   openapitools/openapi-generator-cli:v5.3.0 \
#   generate \
#   -i /local/openapi.json \
#   -g go \
#   -o /local/out \
#   --skip-validate-spec \
#   --global-property "models" \
#   --additional-properties="enumClassPrefix=true,isGoSubmodule=true,packageName=models" \

swagger generate model -m models 
sed -i ':a;N;$!ba;s/struct {\n\t\t\([a-zA-Z]*\)\n\t}/\1/g' models/*.go