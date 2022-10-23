package main

import (
	"fmt"
	"frickelcloud"
)

func main() {
	for name, host := range frickelcloud.DataCenter {
		free := host.AvailableRessources()
		fmt.Printf("host %s has %d CPU cores, %d MB memory, and %d GB storage available\n",
			name, free.CPU, free.RAM, free.SSD)
	}
}
