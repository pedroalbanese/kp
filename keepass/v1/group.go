package keepassv1

import (
	"fmt"
	k "github.com/mostfunkyduck/kp/keepass"
	"zombiezen.com/go/sandpass/pkg/keepass"
)

type Group struct {
	group *keepass.Group
}

func NewGroup(group *keepass.Group) k.Group {
	return &Group{
		group: group,
	}
}

func (g *Group) Name() string {
	return g.group.Name
}

func (g *Group) SetName(name string) {
	g.group.Name = name
}

func (g *Group) Parent() k.Group {
	return NewGroup(g.group.Parent())
}

func (g *Group) SetParent(parent k.Group) error {
	if err := g.group.SetParent(parent.Raw().(*keepass.Group)); err != nil {
		return fmt.Errorf("could not change group parent: %s", err)
	}
	return nil
}

func (g *Group) Entries() (rv []k.Entry) {
	for _, each := range g.group.Entries() {
		rv = append(rv, NewEntry(each))
	}
	return rv
}

func (g *Group) Groups() (rv []k.Group) {
	for _, each := range g.group.Groups() {
		rv = append(rv, NewGroup(each))
	}
	return rv
}

func (g *Group) IsRoot() bool {
	return g.Parent() == nil
}

func (g *Group) NewSubgroup(name string) k.Group {
	newGroup := g.group.NewSubgroup()
	newGroup.Name = name
	return &Group{
		group: newGroup,
	}
}

func (g *Group) RemoveSubgroup(subgroup k.Group) error {
	return g.group.RemoveSubgroup(subgroup.Raw().(*keepass.Group))
}

func (g *Group) Pwd() (fullPath string) {
	group := g.group
	for ; group != nil; group = group.Parent() {
		if group.IsRoot() {
			fullPath = "/" + fullPath
			break
		}
		fullPath = group.Name + "/" + fullPath
	}
	return fullPath
}

func (g *Group) Raw() interface{} {
	return g.group
}

func (g *Group) NewEntry() (k.Entry, error) {
	entry, err := g.group.NewEntry()
	if err != nil {
		return nil, err
	}
	return &Entry{
		entry: entry,
	}, nil
}

func (g *Group) RemoveEntry(e k.Entry) error {
	return g.group.RemoveEntry(e.Raw().(*keepass.Entry))
}
