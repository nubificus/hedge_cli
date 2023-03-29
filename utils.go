package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"

	hedge "github.com/nubificus/hedge_cli/hedge_api"
)

func prettyPrint(vms []hedge.VM) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "NAME", "MODID", "MODNAME"})
	for _, v := range vms {
		t.AppendRow([]interface{}{v.ID, v.Name, v.ModID, v.ModName})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}
