package getResources

import (
	"github.com/kubeless/kubeless/pkg/functions"
	"github.com/sirupsen/logrus"
)

func GetResources(event functions.Event, context functions.Context) (string, error) {
	logrus.Info(event.Data)
	return "", nil
}
func main() {
	logrus.Info("qweqwe")
}
