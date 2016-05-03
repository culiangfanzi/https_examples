package main


import(
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)


func restartHandle(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("this is restart!"))
}


func offlineHandle(w http.ResponseWriter ,r *http.Request){
	w.Write([]byte("this is offline!"))
}


func main(){
	pool := x509.NewCertPool()
	caCertPath :="ca.crt"
	
	caCrt ,err := ioutil.ReadFile(caCertPath)
	if err !=nil{
		fmt.Println("read err:",err)
		return	
}
	pool.AppendCertsFromPEM(caCrt)


	//mux := http.NewServeMux()
	http.HandleFunc("/admin/restart",restartHandle)
	http.HandleFunc("/admin/offline",offlineHandle)


	s:=&http.Server{
		Addr: ":8081",
		Handler: nil,
		TLSConfig :&tls.Config{
			ClientCAs: pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			//ClientAuth :tls.RequestClientCert,
			//InsecureSkipVerify :true,
		},
	}

	

	err = s.ListenAndServeTLS("server.crt","server.key")
	if err!=nil{
	fmt.Println("listen error:",err)	
}


}
