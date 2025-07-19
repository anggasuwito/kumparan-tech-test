package utils

import "fmt"

func GetVerifiedPINKey(accountID string, pinType string) string {
	return fmt.Sprintf("verified_account_pin:%s_%s", accountID, pinType)
}
