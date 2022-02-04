package breadcrumb

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Breadcrumb struct {
	Delimiter string
	Root      Crumb
	status    Crumbs
}

func (b Breadcrumb) current() *Crumb {
	if len(b.status) > 0 {
		return &b.status[len(b.status)-1]
	}
	return &b.Root
}

func (b Breadcrumb) debugInfo() {
	log.Debugf("Delimiter: %s\n", b.Delimiter)
	log.Debugf("Root Crumb: %s\n", b.Root.Name)
	log.Debugf("Current Crumb: %s\n", b.current().Name)
}

func (b Breadcrumb) Print() string {
	if len(b.status) == 0 {
		return strings.TrimSpace(b.Delimiter)
	}
	ret := ""
	for _, t := range b.status {
		ret += (t.Display + b.Delimiter)
	}
	return strings.TrimSpace(ret) + " "
}

func (b *Breadcrumb) sub() {
	if len(b.status) > 0 {
		b.status = b.status[:len(b.status)-1]
	}
}

func (b Breadcrumb) Start() {
	r := bufio.NewReader(os.Stdin)
	var s string
	b.status = Crumbs{b.Root}
	for {
		fmt.Fprint(os.Stderr, b.Print())
		s, _ = r.ReadString('\n')
		cmd, _ := sanitizeInput(s)
		switch cmd {
		case "?":
			b.current().options()
		case "help":
			b.current().Help()
		case "clear":
			ClearScreen()
		case "ls":
			b.current().listChildren()
		case "quit":
			return
		case "..":
			b.sub()
		case "~":
			b.status = Crumbs{b.Root}
		case "set":
			err := b.current().set(s)
			if err != nil {
				log.Error(err)
			}
		case "run":
			err := b.current().run()
			if err != nil {
				log.Error(err)
			}
		default:
			c := b.current()
			if crumb, ok := c.HasChild(strings.TrimSpace(s)); ok {
				b.status = append(b.status, *crumb)
			} else {
				log.Error("Invalid selection")
			}
		}
	}
}
