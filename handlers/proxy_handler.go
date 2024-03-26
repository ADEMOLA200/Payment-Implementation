package handlers

import (
	_"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func ProxyHandler(c *fiber.Ctx) error {
	// Target URL (Paystack API)
	targetURL := "https://checkout.paystack.com" + string(c.OriginalURL())

	// Create a new request to the Paystack API
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(targetURL)
	req.Header.SetMethod(c.Method())

	// Copy headers from the original request
	c.Request().Header.VisitAll(func(key, value []byte) {
		req.Header.SetBytesKV(key, value)
	})

	// Perform the request
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Copy response headers to the client response
	resp.Header.VisitAll(func(key, value []byte) {
		c.Response().Header.SetBytesKV(key, value)
	})

	// Copy status code
	c.Status(resp.StatusCode())

	// Copy - response body
	return c.Send(resp.Body())
}
