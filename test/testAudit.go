package test

import (
	"go-api-starter-kit/utils/audit"
)

type TestAudit struct {
	audit.AuditHandler
}

func NewTestAudit() (*TestAudit, error) {
	audit := &TestAudit{}
	return audit, nil
}

func (l *TestAudit) Write(msg string) error {
	return nil
}
