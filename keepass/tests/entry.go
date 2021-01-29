package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	k "github.com/mostfunkyduck/kp/keepass"
)

func RunTestNoParent(t *testing.T, r Resources) {
	name := "shmoo"
	e := r.Entry
	if !e.Set(k.Value{Name: "Title", Value: []byte(name)}) {
		t.Fatalf("could not set title")
	}
	output, err := e.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}
	// this guy has no parent, shouldn't even have the root "/" in the path
	if output != name {
		t.Fatalf("[%s] !+ [%s]", output, name)
	}

	if parent := e.Parent(); parent != nil {
		t.Fatalf("%v", parent)
	}
}

func RunTestRegularPath(t *testing.T, r Resources) {
	name := "asldkfjalskdfjasldkfjasfd"
	e, err := r.Group.NewEntry(name)
	if err != nil {
		t.Fatalf(err.Error())
	}

	path, err := e.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}
	expected, err := r.Group.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}
	expected += name
	if path != expected {
		t.Fatalf("[%s] != [%s]", path, expected)
	}

	parent := r.Entry.Parent()
	if parent == nil {
		t.Fatalf("%v", r)
	}

	parentPath, err := parent.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}

	groupPath, err := r.Group.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if parentPath != groupPath {
		t.Fatalf("[%s] != [%s]", parentPath, groupPath)
	}

	newEntry := r.BlankEntry
	if err := newEntry.SetParent(r.Group); err != nil {
		t.Fatalf(err.Error())
	}

	entryPath, err := newEntry.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}

	groupPath, err = r.Group.Path()
	if err != nil {
		t.Fatalf(err.Error())
	}

	expected = groupPath + newEntry.Title()
	if entryPath != expected {
		t.Fatalf("[%s] != [%s]", entryPath, expected)
	}
}

// kpv1 only supports a limited set of fields, so we have to let the caller
// specify what value to set

func RunTestEntryTimeFuncs(t *testing.T, r Resources) {
	newTime := time.Now().Add(time.Duration(1) * time.Hour)
	r.Entry.SetCreationTime(newTime)
	if !r.Entry.CreationTime().Equal(newTime) {
		t.Fatalf("%v, %v", newTime, r.Entry.CreationTime())
	}

	newTime = newTime.Add(time.Duration(1) * time.Hour)
	r.Entry.SetLastModificationTime(newTime)
	if !r.Entry.LastModificationTime().Equal(newTime) {
		t.Fatalf("%v, %v", newTime, r.Entry.LastModificationTime())
	}

	newTime = newTime.Add(time.Duration(1) * time.Hour)
	r.Entry.SetLastAccessTime(newTime)
	if !r.Entry.LastAccessTime().Equal(newTime) {
		t.Fatalf("%v, %v", newTime, r.Entry.LastAccessTime())
	}
}
func RunTestEntryPasswordTitleFuncs(t *testing.T, r Resources) {
	password := "swordfish"
	r.Entry.SetPassword(password)
	if r.Entry.Password() != password {
		t.Fatalf("[%s] != [%s]", r.Entry.Password(), password)
	}

	title := "blobulence"
	r.Entry.SetTitle(title)
	if r.Entry.Title() != title {
		t.Fatalf("[%s] != [%s]", r.Entry.Title(), title)
	}
}

func RunTestSearchInNestedSubgroup(t *testing.T, r Resources) {
	sg, err := r.Group.NewSubgroup("RunTestSearchInNestedSubgroup")
	if err != nil {
		t.Fatalf(err.Error())
	}

	e, err := sg.NewEntry("askdfhjaskjfhasf")
	if err != nil {
		t.Fatalf(err.Error())
	}

	paths := r.Db.Root().Search(regexp.MustCompile(e.Title()))

	expected := "/" + r.Group.Name() + "/" + sg.Name() + "/" + e.Title()
	if paths[0] != expected {
		t.Fatalf("[%s] != [%s]", paths[0], expected)
	}
}

// string printed for protected fields
var redactedString = "[redacted]"

func testOutput(e k.Entry, full bool) (output string, failures string) {
	output = e.Output(full)
	for _, value := range e.Values() {
		expected := string(value.Name) + ":\t" + string(value.Value)
		if value.Protected && !full {
			expected = redactedString
		}
		if value.Type == k.BINARY {
			expected = fmt.Sprintf("binary: %d bytes", len(value.Value))
		}
		if !strings.Contains(output, expected) {
			failures = fmt.Sprintf("%svalue [%s] should have been in output\n", failures, expected)
		}
	}
	return
}

func RunTestOutput(t *testing.T, e k.Entry) {
	// layer on a bunch of test values on top of the defaults, using the entry generated by the caller
	e.Set(k.Value{
		Name:  "test binary",
		Value: []byte("this is a test binary"),
		Type:  k.BINARY,
	})

	e.Set(k.Value{
		Name:      "test protected longstring",
		Value:     []byte("this is a test longstring\nlong string\nlongstring"),
		Protected: true,
		Type:      k.LONGSTRING,
	})

	if output, failures := testOutput(e, true); failures != "" {
		t.Fatalf("testing full output failed: \n output:\n%s\nfailures:\n%s", output, failures)
	}

	if output, failures := testOutput(e, false); failures != "" {
		t.Fatalf("testing full output failed: \n output:\n%s\nfailures:\n%s", output, failures)
	}

	if failures := RunTestFormatTime(e); failures != "" {
		t.Fatalf("testing time formatting failed: %s", failures)
	}
}

// This tests the FormatTime utility function via the creation time field output of an entry
func RunTestFormatTime(e k.Entry) (failures string) {
	tests := map[*regexp.Regexp]time.Time{
		// test exactly one day
		regexp.MustCompile(`Creation Time:\t.*\(1 day\(s\) ago\)`): time.Now().Add(time.Duration(-1) * time.Hour * 24),
		// test a full 2 days
		regexp.MustCompile(`Creation Time:\t.*\(2 day\(s\) ago\)`): time.Now().Add(time.Duration(-1) * time.Hour * 24 * 2),
		// test a month
		regexp.MustCompile(`Creation Time:\t.*\(about 1 month\(s\) ago\)`): time.Now().Add(time.Duration(-1) * time.Hour * 24 * 32),
		// test more than 2 months
		regexp.MustCompile(`Creation Time:\t.*\(about 2 month\(s\) ago\)`): time.Now().Add(time.Duration(-1) * time.Hour * 24 * 65),
		// test a year
		regexp.MustCompile(`Creation Time:\t.*\(about 1 year\(s\) ago\)`): time.Now().Add(time.Duration(-1) * time.Hour * 24 * 365),
		// test 2 years
		regexp.MustCompile(`Creation Time:\t.*\(about 2 year\(s\) ago\)`): time.Now().Add(time.Duration(-1) * time.Hour * 24 * 365 * 2),
	}

	for expression, timestamp := range tests {
		e.SetCreationTime(timestamp)
		str := e.Output(true)
		if !expression.Match([]byte(str)) {
			failures = fmt.Sprintf("%soutput [%s] doesn't contain creation time [%s]\n", failures, str, expression)
		}
	}
	return failures
}
