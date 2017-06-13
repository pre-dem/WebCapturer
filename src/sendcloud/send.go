package sendcloud

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const sendMailURL = "http://api.sendcloud.net/apiv2/mail/send"

// SendMail send genera mail
func SendMail(conf map[string]interface{}) (string, error) {
	params := make(map[string]string)
	var files map[string]string
	var ok bool
	for key, value := range conf {
		if key == "attachments" {
			if files, ok = value.(map[string]string); !ok {
				files0 := make(map[string]interface{})
				if files0, ok = value.(map[string]interface{}); ok {
					for key, value := range files0 {
						var stringValue string
						if stringValue, ok = value.(string); ok {
							if files == nil {
								files = make(map[string]string)
							}
							files[key] = stringValue
						}
					}
				}
			}
		} else {
			stringValue, ok := value.(string)
			if ok {
				params[key] = stringValue
			}
		}
	}
	return doRequestWithFile(sendMailURL, params, "attachments", files)
}

func doRequestWithFile(url string, params map[string]string, fileField string, files map[string]string) (string, error) {
	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)
	for fileName, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return "", err
		}

		fileWriter, err := writer.CreateFormFile(fileField, fileName)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(fileWriter, file)
		file.Close()
	}
	for key, value := range params {
		writer.WriteField(key, value)
	}
	var err = writer.Close()
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	responseHandler, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer responseHandler.Body.Close()

	bodyByte, err := ioutil.ReadAll(responseHandler.Body)
	if err != nil {
		return string(bodyByte), err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyByte, &result)
	return string(bodyByte), err
}
