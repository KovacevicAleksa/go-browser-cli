package IO

import (
	"fmt"
	"log"
	"os"
)

func DeleteFile(name string) {

	e := os.Remove(name)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("Succes removed file name: %s\n", name)
}
