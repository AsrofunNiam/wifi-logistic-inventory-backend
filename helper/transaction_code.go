package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func GenerateTransactionCode(db *gorm.DB, model interface{}, prefix string, date time.Time) (string, error) {
	datePart := date.Format("20060102")
	codePrefix := fmt.Sprintf("%s-%s-", prefix, datePart)

	var latestCode string
	err := db.Model(model).
		Where("code LIKE ?", codePrefix+"%").
		Order("code DESC").
		Limit(1).
		Pluck("code", &latestCode).Error
	if err != nil {
		return "", err
	}

	sequence := 1
	if latestCode != "" {
		parts := strings.Split(latestCode, "-")
		if len(parts) > 0 {
			lastSequence, parseErr := strconv.Atoi(parts[len(parts)-1])
			if parseErr != nil {
				return "", parseErr
			}
			sequence = lastSequence + 1
		}
	}

	return fmt.Sprintf("%s%03d", codePrefix, sequence), nil
}
