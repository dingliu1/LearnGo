package network

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example in bi-directional TLS connection!\n")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!")
}

func handlerTls(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example tls server.\n"))
}

func TestHttpHander(t *testing.T) {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":7111", nil)
}

func TestHttpsServer(t *testing.T) {

	http.HandleFunc("/", handlerTls)
	err := http.ListenAndServeTLS(":7111", "server.crt", "server.key", nil)
	log.Fatal(err)
}

func TestHttpsServerVerifyClientCert(t *testing.T) {

	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":7111",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}

func TestHttpsClient(t *testing.T) {
	resp, err := http.Get("https://localhost:7111")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestHttpsClientWithoutVerifyServerCertificate(t *testing.T) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:7111")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// https://tonybai.com/2015/04/30/go-and-https/

func TestHttpsClientVerifyServerCertificate(t *testing.T) {

	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:7111")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestHttpsClientVerifyServerCertificateAndProvideClientCert(t *testing.T) {

	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("restdaclient.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:7111")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}


func TestHttpsClientVerifyServerCertificateAndProvideClientCertPracticeRestda(t *testing.T) {

	pool := x509.NewCertPool()
	caCertPath := "netactrootca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("cert.crt", "mydoman.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}


	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://clab381lbwas.netact.nsn-rdnet.net:9997", nil)
	if err != nil {
		log.Fatal(err)
	}
	auth := "wassrvid:wassrvid"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)
	req.Header.Add("Accept", "application/vnd.nokia-system-credential-response-v1+json;charset=utf-8")

	q := req.URL.Query()
	q.Add("username", "omc")
	q.Add("type", "appserv")
	q.Add("instance", "appserv")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	//resp, err := client.Get("https://clab381lbwas.netact.nsn-rdnet.net:9997")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
