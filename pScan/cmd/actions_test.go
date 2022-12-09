package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/pqppq/pScan/scan"
)

func setup(t *testing.T, hosts []string, initList bool) (string, func()) {
	// create temp file
	tf, err := ioutil.TempFile("", "pScan")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	// initialize list if needed
	if initList {
		hl := &scan.HostList{}
		for _, h := range hosts {
			hl.Add(h)
		}
		if err := hl.Save(tf.Name()); err != nil {
			t.Fatal(err)
		}
	}
	// return temp file name and cleanup function
	return tf.Name(), func() { os.Remove(tf.Name()) }
}

func TestHostActions(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	// test cases for actions tet
	testCases := []struct {
		name       string
		args       []string
		expOut     string
		initList   bool
		actionFunc func(io.Writer, string, []string) error
	}{
		{
			name:       "AddAction",
			args:       hosts,
			expOut:     "Added host: host1\nAdded host: host2\nAdded host: host3\n",
			initList:   false,
			actionFunc: addAction,
		},
		{
			name:       "ListAction",
			expOut:     "host1\nhost2\nhost3\n",
			initList:   true,
			actionFunc: listAction,
		},
		{
			name:       "DeleteAction",
			args:       []string{"host1", "host2"},
			expOut:     "Deleted host: host1\nDeleted host: host2\n",
			initList:   true,
			actionFunc: deleteAction,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup action test
			tf, cleanup := setup(t, hosts, tc.initList)
			defer cleanup()

			// define var to captuer output
			var out bytes.Buffer

			// execute action and captuer action output
			if err := tc.actionFunc(&out, tf, tc.args); err != nil {
				t.Fatalf("Expected no error, got %q\n", err)
			}

			// test action output
			if out.String() != tc.expOut {
				t.Errorf("Expected output %q, but got %q instead", tc.expOut, out.String())
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	// setup integration test
	tf, cleanup := setup(t, hosts, false)
	defer cleanup()

	delHost := "host2"
	hostsEnd := []string{
		"host1", "host3",
	}

	// define var to captuer out
	var out bytes.Buffer

	// difine expected out for all actions
	expectedOut := ""
	for _, v := range hosts {
		expectedOut += fmt.Sprintf("Added host: %s\n", v)
	}

	expectedOut += strings.Join(hosts, "\n")
	expectedOut += fmt.Sprintln()
	expectedOut += fmt.Sprintf("Deleted host: %s\n", delHost)
	expectedOut += strings.Join(hostsEnd, "\n")
	expectedOut += fmt.Sprintln()

	// add hosts to the list
	if err := addAction(&out, tf, hosts); err != nil {
		t.Fatalf("Expected no error, got %q instead\n", err)
	}
	// lists hosts
	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q instead\n", err)
	}
	// delete host2
	if err := deleteAction(&out, tf, []string{delHost}); err != nil {
		t.Fatalf("Expected no error, got %q instead\n", err)
	}
	// lists hosts after delete
	if err := listAction(&out, tf, nil); err != nil {
		t.Fatalf("Expected no error, got %q instead\n", err)
	}

	// testing integration output
	if out.String() != expectedOut {
		t.Errorf("Expected output %q, got %q\n", expectedOut, out.String())
	}
}
