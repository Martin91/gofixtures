package txdrivers

import "fmt"

// TxDriverName returns transactional database driver name
func TxDriverName(driverName string) string {
	return fmt.Sprintf("tx%s", driverName)
}
