package order

import (
	"fmt"
	"strings"
)

func (us *orderStorage) Update(UUID string, data map[string]string) error {

	var params string
	for key, param := range data {
		if param != "" {
			params += strings.ToLower(key + "='" + param + "' ")
		}
	}

	fmt.Println(params)

	if _, err := us.Conn.Exec("update orders set "+params+"where uuid=$1", UUID); err != nil {
		e := fmt.Errorf("fatal update the order: %s", err)
		return e
	}
	return nil
}
