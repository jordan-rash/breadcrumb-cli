package breadcrumb

import (
	"errors"
	"fmt"
)

type Crumb struct {
	Name     string
	Display  string
	Parent   *Crumb
	Children Crumbs
	Help     func() string
	Action   func(map[string]interface{}) error
	Options  []string
	State    map[string]interface{}
}

type Crumbs []Crumb

func (c Crumb) HasChild(input string) (*Crumb, bool) {
	for _, tC := range c.Children {
		if tC.Display == input {
			return &tC, true
		}
	}
	return nil, false
}

func (c *Crumb) set(input string) error {
	_, args := sanitizeInput(input)
	if len(args) == 2 {
		for _, o := range c.Options {
			if args[0] == o {
				c.State[args[0]] = args[1]
				return nil
			}
		}
		return errors.New(fmt.Sprintf("[%s] was not found to be a valid setting for [%s]", args[0], c.Display))
	} else {
		return errors.New("set takes exactly 2 inputs. [setting] [value]")
	}
}

func (c Crumb) listChildren() {
	for _, tC := range c.Children {
		fmt.Printf(tC.Name + "\t" + tC.Display + "\n")
	}
}

func (c Crumb) options() {
	for i, o := range c.Options {
		fmt.Printf("%d. %s\t%v\n", i+1, o, c.State[o])
	}
}

func (c Crumb) run() error {
	return c.Action(c.State)
}
