package main
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"strconv"
)

func main() {
fmt.Println("                                                                                                  ")
fmt.Println(" █████╗ ███████╗███╗   ███╗    ██████╗ ███████╗████████╗███████╗ ██████╗████████╗ ██████╗ ██████╗ ")
fmt.Println("██╔══██╗██╔════╝████╗ ████║    ██╔══██╗██╔════╝╚══██╔══╝██╔════╝██╔════╝╚══██╔══╝██╔═══██╗██╔══██╗")
fmt.Println("███████║█████╗  ██╔████╔██║    ██║  ██║█████╗     ██║   █████╗  ██║        ██║   ██║   ██║██████╔╝")
fmt.Println("██╔══██║██╔══╝  ██║╚██╔╝██║    ██║  ██║██╔══╝     ██║   ██╔══╝  ██║        ██║   ██║   ██║██╔══██╗")
fmt.Println("██║  ██║███████╗██║ ╚═╝ ██║    ██████╔╝███████╗   ██║   ███████╗╚██████╗   ██║   ╚██████╔╝██║  ██║")
fmt.Println("╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝    ╚═════╝ ╚══════╝   ╚═╝   ╚══════╝ ╚═════╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝")
fmt.Println("                                                                            By @akshaysharma71    ")
	var scanner *bufio.Scanner
	if len(os.Args) < 2 {
		// If no argument is provided, use stdin
		scanner = bufio.NewScanner(os.Stdin)
	}else {
		// If an argument is provided, use the file
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
		}
	// Get number of threads
	threads := 1
	if len(os.Args) >= 2 {
		threads = getThreads(os.Args[2])
	}

	var wg sync.WaitGroup
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			aemScan(line)
		}(line)
		if threads != 1 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	wg.Wait()
	fmt.Println("AEM CMS Detection Finished")
}

func aemScan(line string) {
	// AEM login page
	if checkURL(line+"/libs/granite/core/content/login.html", "Welcome to Adobe Experience Manager") {
		fmt.Println(line)
		return
	}

	// Geometrixx page
	if checkURL(line+"/content/geometrixx/en.html", "Geometrixx has been selling") {
		fmt.Println(line)
		return
	}

	// GetServlet
	if checkURL(line+"/.json", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.childrenlist.json", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.1.json", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.childrenlist.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/...4.2.1...json", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json/t.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json/t.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json/t.png", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json/t.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json;%0qw.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json;%0qw.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json;%0qw.png", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json;%0qw.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json?d.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json?d.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.json?d.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json/j.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json/j.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json/j.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json;%0ee.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json;%0ee.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.ext.json;%0ee.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json/a.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json/a.html", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json/a.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json;%0aa.css", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json;%0aa.ico", "jcr:primaryType|data-coral-columnview-path") ||
		checkURL(line+"/.children.json;%0aa.html", "jcr:primaryType|data-coral-columnview-path") {
		fmt.Println(line)
		return
	}

	// Login_Status_Servlet
	if checkURL(line+"/system/sling/loginstatus.json", "authenticated=") ||
		checkURL(line+"/system/sling/loginstatus.css", "authenticated=") ||
		checkURL(line+"/system/sling/loginstatus.png", "authenticated=") {
		fmt.Println(line)
		return
	}

	// Using_CRX
	if checkURL(line+"/crx/de/index.jsp", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/explorer/browser/index.jsp", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/packmgr/index.jsp", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/de/index.jsp;%0aa.css", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/explorer/browser/index.jsp;%0aa.css", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/packmgr/index.jsp;%0aa.css", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/de/index.jsp;%0aa.ico", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/explorer/browser/index.jsp;%0aa.ico", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/packmgr/index.jsp;%0aa.ico", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/de/index.jsp?a.css", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/explorer/browser/index.jsp?a.html", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/packmgr/index.jsp?a.png", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/de/index.jsp/a.ico", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/explorer/browser/index.jsp/a.css", "CRXDE Lite|Content Explorer|CRX Package Manager") ||
		checkURL(line+"/crx/packmgr/index.jsp/a.png", "CRXDE Lite|Content Explorer|CRX Package Manager") {
		fmt.Println(line)
		return
	}

	// Sling
	if checkURL(line+"/system/sling.js", "ADOBE CONFIDENTIAL|JCR repository") ||
		checkURL(line+"/etc/clientlibs/wcm/foundation/main.css", "ADOBE CONFIDENTIAL|JCR repository") {
		fmt.Println(line)
		return
	}
}

func getThreads(threadsArg string) int {
        threads, err := strconv.Atoi(threadsArg)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        return threads
}

func checkURL(url string, keyword string) bool {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	bodyString := string(bodyBytes)
	if strings.Contains(bodyString, keyword) {
		return true
	}
	return false
}
