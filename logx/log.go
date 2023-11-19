package logx

import (
	"fmt"
	"log"
	"os"
)

func Debug(format string, args ...interface{}) {
	if _, set := os.LookupEnv("DEBUG"); !set {
		return
	}

	log.Println(fmt.Sprintf("[DEBUG] "+format, args))
}

func Info() {

}
