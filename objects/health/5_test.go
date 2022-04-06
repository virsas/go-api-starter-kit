package health

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestList(t *testing.T)   {}
func TestCreate(t *testing.T) {}
func TestGet(t *testing.T)    {}
func TestUpdate(t *testing.T) {}
func TestDelete(t *testing.T) {}
