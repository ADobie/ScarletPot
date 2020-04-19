package help

import (
	"fmt"
	"scarletpot/utils/color"
)

func Info() {
	banner := `

███████╗  ██████╗  █████╗  ██████╗  ██╗     ███████╗ ████████╗   ██████╗  ██████╗ ████████╗
██╔════╝ ██╔════╝ ██╔══██╗ ██╔══██╗ ██║     ██╔════╝ ╚══██╔══╝   ██╔══██╗██╔═══██╗╚══██╔══╝
███████╗ ██║      ███████║ ██████╔╝ ██║     █████╗      ██║  *** ██████╔╝██║   ██║   ██║   
╚════██║ ██║      ██╔══██║ ██╔══██╗ ██║     ██╔══╝      ██║  *** ██╔═══╝ ██║   ██║   ██║   
███████║ ╚██████╗ ██║  ██║ ██║  ██║ ███████╗███████╗    ██║  *** ██║     ╚██████╔╝   ██║   
╚══════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═╝  ╚═╝ ╚══════╝╚══════╝    ╚═╝  *** ╚═╝      ╚═════╝    ╚═╝
`

	println(color.Cyan(banner))
	fmt.Println(color.Cyan("   run"), color.White("	         Start all scarlet service"))
	fmt.Println(color.Cyan("   install"), color.White("	 Start install program"))
	fmt.Println(color.Cyan("   version"), color.White("      Show scarletPot Version"))
	fmt.Println(color.Cyan("   help"), color.White("	 Show help"))
	fmt.Println("")
}
