package common

import "fmt"

/*
попытка использовать подход из tkd
есть набор миксинов, которые добавляются там где надо
но не получается сделать сцепку, потому что данная модель является вложенной в виджет и не видит других миксинов
*/
type Anchor struct {
}

func (a *Anchor) SetTextAnchor(position string) *Anchor {
	command := "-anchor"
	// this._tk.eval("%s configure " ~ anchorCommand ~ " %s", this.id, position);
	script := fmt.Sprintf("%s configure %s %s", "TODO: id", command, position)
	fmt.Println(script)
	return a
}

func (a *Anchor) ConfigureAnchor(position string) *Anchor {
	command := "-anchor"
	// this._tk.eval("%s configure " ~ anchorCommand ~ " %s", this.id, position);
	script := fmt.Sprintf("%s configure %s %s", "TODO: id", command, position)
	fmt.Println(script)
	return a
}
