package output

import (
	"encoding/json"
	"github.com/valyala/bytebufferpool"
	"io"
	"net/url"
	"path"
	"strings"
)

type JSONResult struct {
	Url string `json:"url"`
}

func WriteURLs(writer io.Writer, results <-chan string, blacklistMap map[string]struct{}, RemoveParameters bool) error {
	lastURL := make(map[string]struct{})
	for result := range results {
		buf := bytebufferpool.Get()
		if len(blacklistMap) != 0 {
			u, err := url.Parse(result)
			if err != nil {
				continue
			}
			base := strings.Split(path.Base(u.Path), ".")
			ext := base[len(base)-1]
			if ext != "" {
				_, ok := blacklistMap[strings.ToLower(ext)]
				if ok {
					continue
				}
			}
		}
		if RemoveParameters {
			u, err := url.Parse(result)
			if err != nil {
				continue
			}
			if _, ok := lastURL[u.Host+u.Path]; ok {
				continue
			} else {
				lastURL[u.Host+u.Path] = struct{}{} ;
			}

		}

		buf.B = append(buf.B, []byte(result)...)
		buf.B = append(buf.B, "\n"...)
		_, err := writer.Write(buf.B)
		if err != nil {
			return err
		}
		bytebufferpool.Put(buf)
	}
	return nil
}

func WriteURLsJSON(writer io.Writer, results <-chan string, blacklistMap map[string]struct{}, RemoveParameters bool) {
	var jr JSONResult
	enc := json.NewEncoder(writer)
	for result := range results {
		if len(blacklistMap) != 0 {
			u, err := url.Parse(result)
			if err != nil {
				continue
			}
			base := strings.Split(path.Base(u.Path), ".")
			ext := base[len(base)-1]
			if ext != "" {
				_, ok := blacklistMap[strings.ToLower(ext)]
				if ok {
					continue
				}
			}
		}
		jr.Url = result
		if err := enc.Encode(jr); err != nil {
			// todo: handle this error
			continue
		}
	}
}
