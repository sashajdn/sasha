package handler

import (
	"testing"

	"github.com/sashajdn/sasha/service.trade-account/dao"
)

func TestMain(m *testing.M) {
	dao.WithMock()
}
