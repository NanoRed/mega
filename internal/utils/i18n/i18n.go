package i18n

import (
	"fmt"

	"github.com/RedAFD/mega/internal/config"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	switch config.AppLang {
	case "en":
		printer := message.NewPrinter(language.English)
		Sprintf = func(key string, a ...interface{}) string {
			return printer.Sprintf(key, a...)
		}
	default:
		Sprintf = func(key string, a ...interface{}) string {
			if len(a) == 0 {
				return key
			}
			return fmt.Sprintf(key, a...)
		}
	}
}

var Sprintf func(key string, a ...interface{}) string
