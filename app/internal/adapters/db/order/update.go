package order

import (
	"fmt"
)

func (us *orderStorage) Update(UUID string, data map[string]string) error {

	var params string
	var counter int

	for key, param := range data {
		if param != "" {
			params += key + "='" + param + "' "
		}
		counter += 1
		if counter != len(data) {
			params += ", "
		}
	}

	if _, err := us.Conn.Exec("update orders set "+params+"where uuid=$1", UUID); err != nil {
		e := fmt.Errorf("fatal update the order: %s", err)
		return e
	}
	return nil
}
