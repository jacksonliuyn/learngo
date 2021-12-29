package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachine_GetStateName(t *testing.T) {
	m2 := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "LeaderApproveState", m2.GetStateName())
	m2.Approval()
	assert.Equal(t, "FinanceApproveState", m2.GetStateName())
	m2.Reject()
	assert.Equal(t, "LeaderApproveState", m2.GetStateName())
	m2.Approval()
	assert.Equal(t, "FinanceApproveState", m2.GetStateName())
	m2.Approval()
}
