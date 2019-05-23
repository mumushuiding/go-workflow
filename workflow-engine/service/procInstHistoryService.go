package service

import (
	"github.com/mumushuiding/go-workflow/workflow-engine/model"
)

// DelProcInstHistoryByID DelProcInstHistoryByID
func DelProcInstHistoryByID(id int) error {
	return model.DelProcInstHistoryByID(id)
}
