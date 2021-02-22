package logger

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type netHttpClientLogger struct {
	http.Client
	LOGFLAG bool
}

func (bcl *netHttpClientLogger) Get(url string) (resp *http.Response, err error) {
	// capture the response or error
	getResp, getErr := bcl.Client.Get(url)

	ioWriter := os.Stdout

	// create or open a file to log to
	if bcl.LOGFLAG {
		ioWriter, err = os.OpenFile("./get.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	// create a log.Logger to direct output of logging
	logger := log.New(ioWriter, "", log.LstdFlags)

	// logging conditions
	if err != nil {
		logger.Println("FAILURE: GET Request")
		req, _ := httputil.DumpRequest(getResp.Request, true)
		logger.Println("URL: " + string(req))
		logger.Println("ERROR: " + err.Error())
		logger.Println("STATUS: " + getResp.Status)
		//logger.Println("STATUS CODE: " + getResp.StatusCode)
	} else {
		bodyBytes, _ := ioutil.ReadAll(getResp.Body)
		logger.Println("SUCCESS: GET Request")
		logger.Println("URL: " + url)
		logger.Println("RESPONSE: " + string(bodyBytes))
	}

	return getResp, getErr
}

func (bcl *netHttpClientLogger) Post(url string) (resp *http.Response, err error) {
	// capture the response or error
	postResp, postErr := bcl.Client.Get(url)

	ioWriter := os.Stdout

	// create or open a file to log to
	if bcl.LOGFLAG {
		ioWriter, err = os.OpenFile("./get.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	// create a log.Logger to direct output of logging
	logger := log.New(ioWriter, "", log.LstdFlags)

	// logging conditions
	if err != nil {
		logger.Println("FAILURE: POST Request")
		req, _ := httputil.DumpRequest(postResp.Request, true)
		logger.Println("URL: " + string(req))
		logger.Println("ERROR: " + err.Error())
		logger.Println("STATUS: " + postResp.Status)
		//logger.Println("STATUS CODE: " + postResp.StatusCode)
	} else {
		bodyBytes, _ := ioutil.ReadAll(postResp.Body)
		logger.Println("SUCCESS: POST Request")
		logger.Println("URL: " + url)
		logger.Println("RESPONSE: " + string(bodyBytes))
	}

	return postResp, getErr
}

func (bcl *netHttpClientLogger) Delete(url string) (resp *http.Response, err error) {
	// capture the response or error
	deleteResp, getErr := bcl.Client.Get(url)

	ioWriter := os.Stdout

	// create or open a file to log to
	if bcl.LOGFLAG {
		ioWriter, err = os.OpenFile("./get.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	// create a log.Logger to direct output of logging
	logger := log.New(ioWriter, "", log.LstdFlags)

	// logging conditions
	if err != nil {
		logger.Println("FAILURE: GET Request")
		req, _ := httputil.DumpRequest(deleteResp.Request, true)
		logger.Println("URL: " + string(req))
		logger.Println("ERROR: " + err.Error())
		logger.Println("STATUS: " + deleteResp.Status)
		//logger.Println("STATUS CODE: " + deleteResp.StatusCode)
	} else {
		bodyBytes, _ := ioutil.ReadAll(deleteResp.Body)
		logger.Println("SUCCESS: GET Request")
		logger.Println("URL: " + url)
		logger.Println("RESPONSE: " + string(bodyBytes))
	}

	return deleteResp, getErr
}

func newLogger() netHttpClientLogger {
	return netHttpClientLogger{
		LOGFLAG: true,
	}
}

func (bcl *netHttpClientLogger) SetLogFlag(flag bool) {
	bcl.LOGFLAG = flag
}
