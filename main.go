//go:generate swagger generate model --accept-definitions-only -f api/swagger.yml -t ./generated
package main

import "github.com/jpeden/smartthings-nest/cmd"

func main() {
	cmd.Execute()
}
