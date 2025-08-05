package secure

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func GetMachineKey() ([]byte, error) {
	var identifiers []string

	hostname, err := os.Hostname()
	if err == nil {
		identifiers = append(identifiers, hostname)
	}

	currentUser, err := user.Current()
	if err == nil {
		identifiers = append(identifiers, currentUser.Uid, currentUser.Username)
	}

	identifiers = append(identifiers, runtime.GOOS, runtime.GOARCH)

	if len(identifiers) == 0 {
		return nil, fmt.Errorf("failed to gather machine identifiers")
	}

	combined := strings.Join(identifiers, "|")
	hash := md5.Sum([]byte(combined))
	return hash[:], nil
}
