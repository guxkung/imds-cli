# IMDS query to json (cURL with golang)
build with go for http capabilities support

## Support IMDSv2 with `getToken` and `query`commands using Cobra-CLI
## For IMDSv1 `query` command can be utilized without `getToken`
## Test-case for both commands with ec2-metadata-mock to perform SITL
## Containerized the CLI with Dockerfile and build instruction

### Setup development environment
install cobra-cli
`cobra-cli init` # to create main and cmd package stub

`cobra-cli add $command` # to generate the command stub

`go run main.go $command` to run the command

### Examples
`go run main.go query localhost:1338/latest/meta-data` # for list of array outputs 

`go run main.go query localhost:1338/latest/meta-data/ami-id` # for a single value output

### Build/Install
`go build -o cli`

move the executable cli to where the PATH environment variable can reach

`cli $command`

### test with ec2-metadata-mock
start ec2-metadata-mock with default options

run `go test -v` in cmd directory

### build/install with docker (and github actions as CI)
```
docker build -t $REPO/$image:$tag .
docker login
docker push $REPO/$image:$tag
```
on the client side can then `docker pull $REPO/$image:$tag` to start using the container
