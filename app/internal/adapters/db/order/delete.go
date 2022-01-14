package order

import "fmt"

func (us *orderStorage) Delete(UUID string) error {
	if _, err := us.Conn.Exec("delete from orders where uuid=$1", UUID); err != nil {
		e := fmt.Errorf("fatal delete the order: %s", err)
		return e
	}
	return nil
}
