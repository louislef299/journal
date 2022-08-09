package journal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	entries := getAllEntries()

	// Call the render function with the name of the template to render
	Render(c, gin.H{
		"title":   "Home Page",
		"payload": entries}, "index.html")
}

func GetEntry(c *gin.Context) {
	// Check if the article ID is valid
	if date, err := strconv.Atoi(c.Param("date")); err == nil {
		// Check if the article exists
		if entry, err := getEntryByDate(date); err == nil {
			Render(c, gin.H{
				"title":   entry.Title,
				"payload": entry,
			}, "entry.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func AddEntry(c *gin.Context) {
	newEntry := entry{}
	if err := c.Bind(&newEntry); err == nil {
		/*today := time.Now()
		newEntry.Date, err = strconv.Atoi(fmt.Sprintf("%v%v%v", today.Day(), today.Month(), today.Year()))
		if err != nil {
			panic(err)
		}*/
		newEntry.Date = len(entryList) + 1
		entryList = append(entryList, newEntry)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}

}
