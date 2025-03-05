package commands

import (
	"fmt"
	"go-browser/ai"
	"go-browser/io"
	searchbrowser "go-browser/search-browser"
	"go-browser/site"
	"go-browser/types"
	utils "go-browser/utils"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"
)

// HandleHelp displays the help page based on the user's input.
func HandleHelp() {
	page := utils.UserWriteNum("Enter page number for help (e.g., 1, 2, 3):")
	utils.DisplayHelp(page)
}

// HandleExit prints a message and exits the program.
func HandleExit() {
	log.Println("INFO: Exiting program.")
	os.Exit(0)
}

// HandleCreate prompts the user for a file name and content, then creates the file.
func HandleCreate() {
	name := utils.UserWriteString("Enter file name:")
	text := utils.UserWriteString("Enter file content:")
	path := utils.UserWriteString("Enter path: (press enter for main folder)")

	if err := io.CreateFile(name, text, path, true); err != nil {
		log.Printf("ERROR: Failed to create file: %v", err)
	}
}

// HandleRead prompts the user for a file name and reads the content of the file.
func HandleRead() {
	name := utils.UserWriteString("Enter file name for reading:")
	path := utils.UserWriteString("Enter path: (press enter for main folder)")
	io.ReadFile(name, path)
}

// HandleDelete prompts the user for a file name and deletes the file.
func HandleDelete() {
	name := utils.UserWriteString("Enter file name for deletion:")
	io.DeleteFile(name)
}

// HandleUpdate prompts the user for a file name and new content, then updates the file.
func HandleUpdate() {
	name := utils.UserWriteString("Enter file name for update:")
	text := utils.UserWriteString("Enter new content:")
	io.UpdateFile(name, text, false)
}

// HandleRename prompts the user for a file name and a new name, then renames the file.
func HandleRename() {
	name := utils.UserWriteString("Enter file name to rename:")
	newName := utils.UserWriteString("Enter new file name:")
	path := utils.UserWriteString("Enter path: (press enter for main folder)")
	io.RenameFile(name, newName, path)
}

// HandleAbout displays information about the program version.
func HandleAbout() {
	fmt.Println("INFO: Version 0.0.0 - Go Browser Tool")
}

// HandleList lists all files in the current directory.
func HandleList() {
	path := utils.UserWriteString("Enter Path: (enter for main folder)")
	io.ListFile(path)
}

func HandleSystemLogs() {
	path := "user_files/logs"
	size, err := utils.FolderSize(path)
	FormatedSize := utils.FormatFileSize(size)
	fmt.Println(FormatedSize)
	log.Println(err)
}

// HandleAIChat prompts the user for text and sends it to the AI model for a response.
func HandleAIChat() {
	text := utils.UserWriteString("Enter text:")
	response := ai.ChatAI(text)
	fmt.Println("AI Response:", response)
}

// HandleGoogle prompts the user for search text and performs a Google search.
func HandleGoogle() {
	search := utils.UserWriteString("Enter text for search:")

	result, err := searchbrowser.SearchGoogle(search)
	if err != nil {
		log.Println("ERROR: Error:", err)
		return
	}

	// Display the search results
	for _, res := range result {
		fmt.Println("Search Result:", res)
	}

	historytext := "Google search: " + "search " + search + "\n"
	io.UpdateFile("../user_files/history/history.txt", historytext, false)
}

// HandleSitePerformance prompts the user for a URL and performs a performance test on the site.
func HandleSitePerformance() {
	url := utils.UserWriteString("Enter site URL to test performance:")
	live := utils.UserWriteBool("Enable live monitoring? (true/false):")
	interval := utils.UserWriteNum("How often send request in secounds")
	timeout := 10 * time.Second
	err := site.MeasureSitePerformance(url, timeout, live, interval)
	if err != nil {
		log.Println("ERROR: Error:", err)
	}
}

// HandleSiteContent prompts the user for input to fetch and display content from a specified site.
func HandleSiteContent() {
	url := utils.UserWriteString("Enter site URL:")
	element := utils.UserWriteString("Specify the target element (or leave empty for all):")
	includeAttributes := utils.UserWriteBool("(true/false) Include attributes in HTML elements?")
	filter := utils.UserWriteBool("(true/false) Filter unnecessary elements like scripts?")
	langBool := utils.UserWriteBool("(true/false) Default language en-US")
	lang := "en-US"
	if !langBool {
		lang = utils.UserWriteString("Language: (en-US, es-ES, zh-CN..)")
	}

	// Create an instance of SiteOptions with user input
	options := types.SiteOptions{
		URL:               url,
		Element:           element,
		Language:          lang,
		IncludeAttributes: includeAttributes,
		Filter:            filter,
	}

	// Call SiteContent with the SiteOptions instance
	content, err := site.SiteContent(options)
	if err != nil {
		log.Println("ERROR: Error:", err)
		return
	}

	// Display the extracted content
	fmt.Println("Extracted Content:", content)

	// Ask the user if they want to save the content
	save := utils.UserWriteBool("(true/false) Save content?")
	if save {
		name := utils.UserWriteString("Enter file name to save content: (.text, .html...)")
		path := utils.UserWriteString("Enter path: (press enter for main folder)")
		if err := io.CreateFile(name, content, path, false); err != nil {
			log.Printf("ERROR: Failed to save content to file: %v\n", err)
		} else {
			log.Println("INFO: Content saved successfully.")
		}
	}
	historytext := "SiteContent" + "url " + url + "element: " + element + "includeAttributes " + strconv.FormatBool(includeAttributes) + "langBool " + strconv.FormatBool(langBool) + "lang " + lang + "\n"
	io.UpdateFile("../user_files/history/history.txt", historytext, false)
}

// HandleHttpRequest handles an HTTP request by taking user input for the URL, method, and optional body content.
// It then performs the request and logs the response and status code.
func HandleHttpRequest() {
	// Get the URL input from the user
	url := utils.UserWriteString("Enter URL for Fetch:")

	// Example: https://jsonplaceholder.typicode.com/posts
	// POST
	// {"id":101,"title":"foo","body":"bar","userId":1}
	method := utils.UserWriteString("Enter method:")

	var bodyContent string

	// Get JSON body input if needed (for POST method)
	if method == "POST" {
		bodyContent = utils.UserWriteJson("Enter Body: {..}")
	}

	// If the method is POST, format the body as JSON
	var body []byte
	if method == "POST" && bodyContent != "" {
		body = []byte(bodyContent)
	} else if method == "POST" && bodyContent == "" {
		// Body cannot be empty for POST method
		log.Println("ERROR: For POST method, the body cannot be empty.")
		return
	} else {
		// For GET method, the body is empty
		body = nil
	}

	// Call the HTTP request function with the given parameters
	response, statusCode, err := site.HttpRequest(url, method, body)
	if err != nil {
		// Handle errors if the HTTP request fails
		log.Println("ERROR: Error:", err)
	} else {
		slog.Info("Response: ", "Status Code", statusCode)
		// Print the response and status code
		fmt.Printf("Response: %s\nStatus Code: %d\n", response, statusCode)
	}
	historytext := "HTTPRequest: " + " url " + url + "method: " + method + "statusCode " + strconv.Itoa(statusCode) + "\n"
	io.UpdateFile("../user_files/history/history.txt", historytext, false)
}

// HandleHistory manages the history of executed HTTP requests.
// It allows the user to delete the history file or view its contents.
func HandleHistory() {
	fmt.Println("History commands (delete, view)")
	command := utils.UserWriteString("Enter command:")

	if !io.PathExists("./user_files/history/history.txt") {
		fmt.Println("History file does not exist.")
		return
	}

	switch command {
	case "delete":
		//empty file
		io.UpdateFile("../user_files/history/history.txt", "", true)
		fmt.Println("History file deleted.")
	case "view":
		io.ReadFile("history.txt", "../user_files/history")
	default:
		fmt.Println("Unknown command")
	}
}
