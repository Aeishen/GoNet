#一个php客户端程序，通过socket连接调用jsonrpc实现的go服务端RPC方法。

<?php

class JsonRPC {

    private $conn;

    function __construct($host, $port) {
        $this->conn = fsockopen($host, $port, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }

    public function Call($method, $params) {
        if (!$this->conn) {
            return false;
        }
        $err = fwrite($this->conn, json_encode(array(
                'method' => $method,
                'params' => array($params),
                'id'     => 0,
            ))."\n");
        if ($err === false) {
            return false;
        }
        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }
        return json_decode($line,true);
    }
}

$client = new JsonRPC("localhost", 1234);
$args = array('N'=>9, 'M'=>2);
$r = $client->Call("Args.Multiply", $args);
printf("%d * %d = %d\n", $args['N'], $args['M'], $r['result']);
