package file

import (
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/vault/audit"
	"github.com/hashicorp/vault/helper/salt"
)

func TestAuditFile_fileModeNew(t *testing.T) {
	salter, _ := salt.NewSalt(nil, nil)
	
	modeStr := "0777"
	mode, err := strconv.ParseUint(modeStr, 8, 32)
	
	file := "auditTest.txt"
	
	config := map[string]string{
		"path": file,
		"mode": modeStr,
	}

	_, err = Factory(&audit.BackendConfig{
		Salt:   salter,
		Config: config,
	})
	
	if err != nil {
		t.Fatalf("%v", err)
	}

    info,_ := os.Stat(file)
    createdMode := info.Mode()
	if createdMode != os.FileMode(mode) {
		t.Fatalf("File Mode does not match.")
	}
	
	err = os.Remove(file)
}

func TestAuditFile_fileModeExisting(t *testing.T) {
	salter, _ := salt.NewSalt(nil, nil)
	
	modeStr := "0600"
	m, err := strconv.ParseUint(modeStr, 8, 32)
	mode := os.FileMode(m)
	
	file := "auditTest.txt"
	
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	err = f.Close()
	if err != nil {
		t.Fatalf("Failure to close the file.")
	}
	
	config := map[string]string{
		"path": file,
	}

	_, err = Factory(&audit.BackendConfig{
		Salt:   salter,
		Config: config,
	})
	
	if err != nil {
		t.Fatalf("%v", err)
	}

    info,_ := os.Stat(file)
    createdMode := info.Mode()
	if createdMode != mode {
		t.Fatalf("File Mode does not match.")
	}
	
	err = os.Remove(file)
}