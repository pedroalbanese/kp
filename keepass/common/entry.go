package common

// Entry is a wrapper around an entry driver, holding functions
// common to both kp1 and kp2
import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	k "github.com/mostfunkyduck/kp/keepass"
)

type Entry struct {
	db     k.Database
	driver k.Entry
}

// findPathToEntry returns all the groups in the path leading to an entry *but not the entry itself*
// the path returned will also not include the source group
func findPathToEntry(source k.Group, target k.Entry) (rv []k.Group, err error) {
	// this library doesn't appear to support child->parent links, so we have to find the needful ourselves
	for _, entry := range source.Entries() {
		equal, err := CompareUUIDs(target, entry)
		if err != nil {
			return []k.Group{}, err
		}
		if equal {
			return []k.Group{source}, nil
		}
	}

	groups := source.Groups()
	for _, group := range groups {
		newGroups, err := findPathToEntry(group, target)
		if err != nil {
			// not putting the path in this error message because it might trigger an infinite loop
			// since this is part of the path traversal algo
			return []k.Group{}, fmt.Errorf("error finding path to '%s' from '%s': %s", target.Title(), group.Name(), err)
		}
		if len(newGroups) > 0 {
			return append([]k.Group{source}, newGroups...), nil
		}
	}
	return []k.Group{}, nil
}

// Path returns the fully qualified path to the entry, if there's no parent, only the name is returned
func (e *Entry) Path() (path string, err error) {
	pathGroups, err := findPathToEntry(e.DB().Root(), e.driver)
	if err != nil {
		return path, fmt.Errorf("could not find path from root to %s: %s", e.driver.Title(), err)
	}

	for _, each := range pathGroups {
		path = path + each.Name() + "/"
	}
	path = path + e.driver.Title()
	return
}

func (e *Entry) Parent() k.Group {
	pathGroups, err := findPathToEntry(e.DB().Root(), e.driver)
	if err != nil {
		return nil
	}
	if len(pathGroups) == 0 {
		return nil
	}

	return pathGroups[len(pathGroups)-1]
}

func (e *Entry) SetParent(g k.Group) error {
	pathGroups, err := FindPathToGroup(e.DB().Root(), g)
	if len(pathGroups) == 0 || err != nil {
		errorString := fmt.Sprintf("could not find a path from the db root to '%s', is this a valid group?", g.Name())

		if err != nil {
			errorString = errorString + fmt.Sprintf(" (error occurred: %s)", err)
		}
		return fmt.Errorf(errorString)
	}

	// this constitutes a move, so remove the entry from its old parent and put it in the new one
	if parent := e.Parent(); parent != nil {
		if err := parent.RemoveEntry(e.driver); err != nil {
			return fmt.Errorf("could not remove entry from old parent: %s", err)
		}
	}

	// add the now-orphaned entry to the new parent
	if err := g.AddEntry(e.driver); err != nil {
		return fmt.Errorf("cannot add entry to group: %s", err)
	}
	return nil
}

func (e *Entry) Output(full bool) (val string) {
	var b strings.Builder
	val = "\n"
	fmt.Fprintf(&b, "\n")
	// Output all the metadata first
	uuidString, err := e.driver.UUIDString()
	if err != nil {
		uuidString = fmt.Sprintf("<could not render UUID string: %s>", err)
	}

	// b64 the UUID string since it sometimes contains garbage characters, esp in v2
	fmt.Fprintf(&b, "UUID:\t%s\n", base64.StdEncoding.EncodeToString([]byte(uuidString)))
	fmt.Fprintf(&b, "Creation Time:\t%s\n", FormatTime(e.driver.CreationTime()))
	fmt.Fprintf(&b, "Last Modified:\t%s\n", FormatTime(e.driver.LastModificationTime()))
	fmt.Fprintf(&b, "Last Accessed:\t%s\n", FormatTime(e.driver.LastAccessTime()))
	expiredTime := e.driver.ExpiredTime()
	// do we want to highlight this as expired?
	highlightExpiry := false
	if expiredTime != (time.Time{}) && expiredTime.Before(time.Now()) {
		highlightExpiry = true
		fmt.Fprintf(&b, "\033[31m")
	}
	fmt.Fprintf(&b, "Expiration Date:\t%s\n", FormatTime(e.driver.ExpiredTime()))
	if highlightExpiry {
		fmt.Fprintf(&b, "\033[0m")
	}

	values, err := e.driver.Values()
	if err != nil {
		val = "error while reading values: " + err.Error()
		return
	}
	for _, val := range values {

		title := strings.Title(val.Name())

		fmt.Fprintf(&b, "%s:\t%s\n", title, val.FormattedValue(full))
	}
	return b.String()
}

// TODO test various fields to make sure they are searchable, consider adding searchability toggle
func (e *Entry) Search(term *regexp.Regexp) (paths []string, err error) {
	values, err := e.driver.Values()
	if err != nil {
		return []string{}, fmt.Errorf("error reading values from entry: %s", err)
	}
	for _, val := range values {
		if !val.Searchable() {
			continue
		}
		content := string(val.FormattedValue(true))
		if term.FindString(content) != "" {
			// something in this entry matched, let's return it
			path, _ := e.Path()
			paths = append(paths, path)
			break
		}
	}

	return
}

func (e *Entry) DB() k.Database {
	return e.db
}

func (e *Entry) SetDB(db k.Database) {
	e.db = db
}

// SetEntry sets the internal entry driver for this wrapper
func (e *Entry) SetDriver(entry k.Entry) {
	e.driver = entry
}