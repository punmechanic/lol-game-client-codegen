all: api/

openapi.json:
	wget --ca-certificate riotgames.pem https://127.0.0.1:2999/swagger/v3/openapi.json

openapi-generator-cli.jar:
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.0/openapi-generator-cli-4.3.0.jar -O openapi-generator-cli.jar

api/: openapi-generator-cli.jar openapi.json
	java -jar openapi-generator-cli.jar generate -i openapi.json -g go --skip-validate-spec -o api
