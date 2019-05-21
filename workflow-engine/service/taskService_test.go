package service_test

import (
	"testing"

	"github.com/mumushuiding/go-workflow/workflow-engine/service"
)

func TestWithdrawTest(t *testing.T) {
	service.WithDrawTask(1515, 832, "11025")
}
