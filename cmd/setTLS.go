/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>
limitations under the License.
*/
package cmd

import (
	"crypto/tls"
	"net/http"
)

func SetPolicy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
