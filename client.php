    public function actionSimplehttps(){
         $url="https://127.0.0.1:8081/admin/offline";
        $curl=curl_init();
        curl_setopt($curl, CURLOPT_URL, $url);
        //curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 0);
        curl_setopt($curl,CURLOPT_SSL_VERIFYPEER,true);
        curl_setopt($curl,CURLOPT_CAINFO,  dirname(__DIR__)."/config/ca.crt");
        curl_setopt($curl,CURLOPT_SSLCERT,dirname(__DIR__)."/config/client.crt");
        curl_setopt($curl,CURLOPT_SSLKEY,dirname(__DIR__)."/config/client.key");
        curl_setopt($curl, CURLOPT_SSLKEYTYPE, "PEM");
        $res2=curl_exec($curl);
        var_dump($res2);
    }
