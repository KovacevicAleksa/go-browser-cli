package commands

import (
	"fmt"
	"time"

	AI "go-browser/AI"
	IO "go-browser/IO"
	searchbrowser "go-browser/search-browser"
	"go-browser/site"
	utils "go-browser/utils"
)

func HandleHelp() {
	page := utils.UserWriteNum("Enter page number for help (e.g., 1, 2, 3):")
	utils.DisplayHelp(page)
}

func HandleExit() {
	fmt.Println("Exiting program.")
}

func HandleCreate() {
	name := utils.UserWriteString("Enter file name:")
	text := utils.UserWriteString("Enter file content:")
	IO.CreateFile(name, text)
}

func HandleRead() {
	name := utils.UserWriteString("Enter file name for reading:")
	IO.ReadFile(name)
}

func HandleDelete() {
	name := utils.UserWriteString("Enter file name for deletion:")
	IO.DeleteFile(name)
}

func HandleUpdate() {
	name := utils.UserWriteString("Enter file name for update:")
	text := utils.UserWriteString("Enter new content:")
	IO.UpdateFile(name, text)
}

func HandleRename() {
	name := utils.UserWriteString("Enter file name to rename:")
	newName := utils.UserWriteString("Enter new file name:")
	IO.RenameFile(name, newName)
}

func HandleAbout() {
	fmt.Println("Version 0.0.0 - Go Browser Tool")
}

func HandleList() {
	IO.ListFile(".")
}

func HandleAIChat() {
	text := utils.UserWriteString("Enter text:")
	response := AI.ChatGPT(text)
	fmt.Println(response)
}

func HandleGoogle() {
	search := utils.UserWriteString("Enter text for search:")
	result, err := searchbrowser.SearchGoogle(search)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func HandleSitePerformance() {
	url := utils.UserWriteString("Enter site URL to test performance:")
	live := utils.UserWriteBool("Enable live monitoring? (true/false):")
	timeout := 10 * time.Second
	err := site.MeasureSitePerformance(url, timeout, live)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func HandleSiteContent() {
	url := utils.UserWriteString("Enter site URL:")
	element := utils.UserWriteString("Specify the target element (or leave empty for all):")
	includeAttributes := utils.UserWriteBool("(true/false) Include attributes in HTML elements?")
	filter := utils.UserWriteBool("(true/false) Filter unnecessary elements like scripts?")

	content, err := site.SiteContent(url, element, includeAttributes, filter)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Extracted Content:")
		fmt.Println(content)

		save := utils.UserWriteBool("(true/false) Save content?")
		if save {
			name := utils.UserWriteString("Enter file name to save content:")
			IO.CreateFile(name, content)
		}
	}
}
