package main

import (
	"fmt"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
)

func main() {
	// Create a capabilities object.
	capabilities := goselenium.Capabilities{}

	// Populate it with the browser you wish to use.
	capabilities.SetBrowser(goselenium.ChromeBrowser())

	// Initialise a new web driver.
	driver, err := goselenium.NewSeleniumWebDriver("http://selenium.apps.okd.codespring.ro/wd/hub", capabilities)
	i := 0

	for err != nil {
		time.Sleep(100 * time.Millisecond)
		driver, err = goselenium.NewSeleniumWebDriver("http://selenium.apps.okd.codespring.ro/wd/hub", capabilities)
		i++
		if i > i && err != nil {
			fmt.Println(err)
			return
		}
	}

	// Create a session.
	_, err = driver.CreateSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Defer the deletion of the session.
	defer driver.DeleteSession()

	// Navigate to Google.
	_, err = driver.Go("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}

	// Hooray, we navigated to Google!
	fmt.Println("Successfully navigated to Google!")
}
