package scan_test

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pqppq/pScan/scan"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name   string
		host   string
		expLen int
		expErr error
	}{
		{"AddNew", "host2", 2, nil},
		{"AddExisting", "host1", 1, scan.ErrExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostList{}
			// initialize list
			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}
			err := hl.Add(tc.host)

			if tc.expErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expErr) {
					t.Fatalf("Expected error %q, got %q instead\n", tc.expErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}
			if len(hl.Hosts) != tc.expLen {
				t.Errorf("Expected list lejgth %d, got %d instead\n", tc.expLen, len(hl.Hosts))
			}
			if hl.Hosts[1] != tc.host {
				t.Errorf("Expected host name %q as index 1, got %q instead\n", tc.host, hl.Hosts[1])
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name   string
		host   string
		expLen int
		expErr error
	}{
		{"RemoveExisting", "host1", 1, nil},
		{"RemoveNotFound", "host3", 1, scan.ErrNotExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostList{}

			// initialize list
			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}

			err := hl.Remove(tc.host)
			if tc.expErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expErr) {
					t.Fatalf("Expected error %q, got %q instead\n", tc.expErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}
			if len(hl.Hosts) != tc.expLen {
				t.Errorf("Expected list lejgth %d, got %d instead\n", tc.expLen, len(hl.Hosts))
			}
			if hl.Hosts[0] == tc.host {
				t.Errorf("Host name %q should not be in the list\n", tc.host)
			}
		})
	}
}

func TestSave(t *testing.T) {
	hl1 := scan.HostList{}
	hl2 := scan.HostList{}

	hostName := "host1"
	hl1.Add(hostName)

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s\n", err)
	}
	defer os.Remove(tf.Name())

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s\n", err)
	}
	if err := hl2.Load(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s\n", err)
	}

	if hl1.Hosts[0] != hl2.Hosts[0] {
		t.Fatalf("Host %q should match %q host\n", hl1.Hosts[0], hl2.Hosts[0])
	}
}

func TestLoadNoFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s\n", err)
	}
	if err := os.Remove(tf.Name()); err != nil {
		t.Fatalf("Error deleting temp file: %s\n", err)
	}
	hl := &scan.HostList{}
	if err := hl.Load(tf.Name()); err != nil {
		t.Fatalf("Expected no error, got %s instead\n", err)
	}
}
