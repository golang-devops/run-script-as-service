package main

import (
	"fmt"
	"strings"
)

func getStringFromRecovery(r interface{}) string {
	switch t := r.(type) {
	case error:
		return t.Error()
	}
	return fmt.Sprintf("%#v", r)
}

func cleanOutput(output string) string {
	s := output
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\r", "", -1)
	return s
}
