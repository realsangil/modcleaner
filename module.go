package main

import (
	"fmt"
	"strings"
)

type Module struct {
	Package string
	Version string
}

func (m Module) String() string {
	if m.Version == "" {
		return m.Package
	}
	return fmt.Sprintf("%s@%s", m.Package, m.Version)
}

func parseModule(str string) Module {
	m := strings.Split(str, "@")
	ver := ""
	if len(m) == 2 {
		ver = m[1]
	}
	pkg := m[0]

	return Module{
		Package: pkg,
		Version: ver,
	}
}
