package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		reqDump := dumpRequest(c.Request(), reqBody)
		respDump := dumpResponse(c.Response().Status, c.Response().Header(), resBody)
		fmt.Println("Request:\n", reqDump)
		fmt.Println("Response:\n", respDump)
		c.String(http.StatusOK, fmt.Sprintf(`
Request received by server:
%s
===
Response sent by server:	
%s
`, reqDump, respDump))
	}))

	if os.Getenv("ALLOW_ORIGINS") != "" {
		// Configure CORS
		config := middleware.CORSConfig{
			AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
			AllowHeaders: strings.Split(os.Getenv("ALLOW_HEADERS"), ","),
			AllowMethods: strings.Split(os.Getenv("ALLOW_METHODS"), ","),
		}
		e.Use(middleware.CORSWithConfig(config))
	}

	e.Static("/static", "assets")
	e.Any("/*", func(_ echo.Context) error { return nil })

	var addr string
	if addr = os.Getenv("LISTEN_ADDR"); addr == "" {
		addr = ":8989"
	}
	e.Logger.Fatal(e.Start(addr))
}

func dumpRequest(r *http.Request, body []byte) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Method: %s\nURL: %s\nHeaders:\n", r.Method, r.URL))
	for k, v := range r.Header {
		sb.WriteString(fmt.Sprintf("\t%s: %s\n", k, v))
	}
	sb.WriteString("Body:")
	sb.WriteString(fmt.Sprintf("\t%s\n", string(body)))
	return sb.String()
}

func dumpResponse(status int, header http.Header, body []byte) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Status: %d\nHeaders:\n", status))
	for k, v := range header {
		sb.WriteString(fmt.Sprintf("\t%s: %s\n", k, v))
	}
	sb.WriteString("Body")
	sb.WriteString(fmt.Sprintf("\t%s\n", string(body)))
	return sb.String()
}
