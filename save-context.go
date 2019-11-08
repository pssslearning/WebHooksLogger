package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func saveContextToFile(c *gin.Context) error {
	fileName, err := getFileName(c.Request.Method)
	if err != nil {
		return err
	}

	fh, err := os.Create("./log/" + fileName + ".txt")
	if err != nil {
		return err
	}
	defer fh.Close()

	w := bufio.NewWriter(fh)

	_, err = w.WriteString("\n---- URL & Query Params (begin) ------")
	if err != nil {
		return err
	}

	_, err = w.WriteString("\n" + c.Request.RequestURI)
	if err != nil {
		return err
	}

	queryParams := c.Request.URL.Query()
	for k, v := range queryParams {
		_, err = w.WriteString("\n" + k + ": " + strings.Join(v, ", "))
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString("\n---- URL & Query Params (end) --------")
	if err != nil {
		return err
	}

	w.Flush()
	if err != nil {
		return err
	}

	theHeadersMap := c.Request.Header
	_, err = w.WriteString("\n---- HEADERS (begin) -----------------")
	if err != nil {
		return err
	}

	for k, v := range theHeadersMap {
		_, err = w.WriteString("\n" + k + ": " + strings.Join(v, ", "))
		if err != nil {
			return err
		}
	}

	_, err = w.WriteString("\n---- HEADERS (end) -------------------")
	if err != nil {
		return err
	}

	w.Flush()
	if err != nil {
		return err
	}

	_, err = w.WriteString("\n---- BODY    (begin) -----------------\n")
	if err != nil {
		return err
	}
	br := c.Request.Body
	defer br.Close()
	body, err := ioutil.ReadAll(br)
	if err != nil {
		return err
	}
	_, err = w.Write(body)
	if err != nil {
		return err
	}

	_, err = w.WriteString("\n---- BODY    (end) -------------------")
	if err != nil {
		return err
	}

	w.Flush()
	if err != nil {
		return err
	}

	return nil
}
