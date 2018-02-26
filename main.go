package main

func main() {
	cm := CommandManager{}
	cm.init()
	var c ICommand
	c = &CommandInit{}
	c.build("init", "Initialize dev enviroments")
	c.addOption("f", "force")
	cm.addCommand(c)
	cm.run()

	//h := Help{}
	//h.run()
	//fmt.Println(h.ProjectName)
}
