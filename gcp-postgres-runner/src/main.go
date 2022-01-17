package main

import (
	"github.com/tumelohq/docker-images/gcp-postgres-runner/src/cmd"
	_ "gocloud.dev/postgres/gcppostgres"
)

func main() {
	cmd.Execute()
}
