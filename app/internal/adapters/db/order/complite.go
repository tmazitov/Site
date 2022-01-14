package order

import "fmt"

func (us *orderStorage) Complite(UUID string) error {
	if _, err := us.Conn.Exec("update orders set status=$1 where uuid=$2", "done", UUID); err != nil {
		e := fmt.Errorf("fatal comlite the order: %s", err)
		return e
	}
	return nil
}
