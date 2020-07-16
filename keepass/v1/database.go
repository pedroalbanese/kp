package keepassv1

// wraps a v1 database with utility functions that allow it to be integrated
// into the shell.

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	k "github.com/mostfunkyduck/kp/keepass"
	"zombiezen.com/go/sandpass/pkg/keepass"
)

type Database struct {
	currentLocation k.Group
	db              *keepass.Database
	savePath        string
}

var backupExtension = ".kpbackup"

func NewDatabase(db *keepass.Database, savePath string) k.Database {
	rv := &Database{
		currentLocation: NewGroup(db.Root()),
		db:              db,
		savePath:        savePath,
	}
	return rv
}

// traversePath will, given a starting location and a UNIX-style path, will walk the path and return the final location or an error
// if the path points to an entry, the parent group is returned as well as the entry.
// If the path points to a group, the entry will be nil
func (d *Database) TraversePath(startingLocation k.Group, fullPath string) (finalLocation k.Group, finalEntry k.Entry, err error) {
	currentLocation := startingLocation
	root := d.Root()
	if fullPath == "/" {
		// short circuit now
		return root, nil, nil
	}

	if strings.HasPrefix(fullPath, "/") {
		// the user entered a fully qualified path, so start at the top
		currentLocation = root
	}

	// break the path up into components, remove terminal slashes since they don't actually do anything
	path := strings.Split(strings.TrimSuffix(fullPath, "/"), "/")
	// tracks whether or not the traversal encountered an entry
loop:
	for i, part := range path {
		if part == "." || part == "" {
			continue
		}

		if part == ".." {
			// if we're not at the root, go up a level
			if currentLocation.Parent() != nil {
				currentLocation = currentLocation.Parent()
				continue
			}
			// we're at the root, the user wanted to go higher, that's no bueno
			return nil, nil, fmt.Errorf("tried to go to parent directory of '/'")
		}

		// regular traversal
		for _, group := range currentLocation.Groups() {
			// was the entity being searched for this group?
			if group.Name() == part {
				currentLocation = group
				continue loop
			}
		}

		for j, entry := range currentLocation.Entries() {
			// is the entity we're looking for this index or this entry?
			if entry.Get("title").Value.(string) == part || strconv.Itoa(j) == part {
				if i != len(path)-1 {
					// we encountered an entry before the end of the path, entries have no subgroups,
					// so this path is invalid
					return nil, nil, fmt.Errorf("invalid path '%s': '%s' is an entry, not a group", entry.Pwd(), fullPath)
				}
				// this is the end of the path, return the parent group and the entry
				return currentLocation, entry, nil
			}
		}
		// getting here means that we found neither a group nor an entry that matched 'part'
		// both of the loops looking for those short circuit when they find what they need
		return nil, nil, fmt.Errorf("could not find a group or entry named '%s'", part)
	}
	// we went all the way through the path and it points to currentLocation,
	// if it pointed to an entry, it would have returned above
	return currentLocation, nil, nil
}

// Root returns the DB root
func (d *Database) Root() k.Group {
	return NewGroup(d.db.Root())
}

// Backup will create a backup of the current database to a temporary location
// in case saving the database causes some kind of corruption
func (d *Database) Backup() error {
	backupPath := d.SavePath() + backupExtension
	w, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("could not open file '%s': %s", backupPath, err)
	}

	if err := d.db.Write(w); err != nil {
		return fmt.Errorf("could not write to file '%s': %s", backupPath, err)
	}
	return nil
}

// RemoveBackup removes a temporary backup file
func (d *Database) RemoveBackup() error {
	backupPath := d.SavePath() + backupExtension
	if err := os.Remove(backupPath); err != nil {
		return fmt.Errorf("could not remove backup file '%s': %s", backupPath, err)
	}
	return nil
}

// SavePath returns the current save location for the DB
func (d *Database) SavePath() string {
	return d.savePath
}

func (d *Database) SetSavePath(newPath string) {
	d.savePath = newPath
}

// Save will backup the DB, save it, then remove the backup is the save was successful
func (d *Database) Save() error {
	savePath := d.SavePath()

	if err := d.Backup(); err != nil {
		return fmt.Errorf("could not back up database: %s", err)
	}

	w, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("could not open/create db save location [%s]: %s", savePath, err)
	}

	if err = d.db.Write(w); err != nil {
		return fmt.Errorf("error writing database to [%s]: %s", savePath, err)
	}

	if err := d.RemoveBackup(); err != nil {
		return fmt.Errorf("could not remove backup after saving: %s", err)
	}
	return nil
}

func (d *Database) SetOptions(o k.Options) error {
	opts := &keepass.Options{
		Password: o.Password,
		KeyFile:  o.KeyReader,
	}
	if err := d.db.SetOpts(opts); err != nil {
		return fmt.Errorf("could not set DB options: %s", err)
	}
	return nil
}

func (d *Database) CurrentLocation() k.Group {
	return d.currentLocation
}

func (d *Database) SetCurrentLocation(g k.Group) {
	d.currentLocation = g
}

func (d *Database) Raw() interface{} {
	return d.db
}

// Pwd will walk up the group hierarchy to determine the path to the current location
func (d *Database) Pwd() (fullPath string) {
	group := d.CurrentLocation()
	return group.Pwd()
}
