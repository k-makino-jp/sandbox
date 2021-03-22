const fs = require('fs');
// var http = require('http');
var https = require('https');

// http server
// var server = http.createServer('request', function(req, res) {
//     res.writeHead(200, {'Content-Type' : 'text/plain'});
//     res.write('hello world');
//     res.end();
// }).listen(8080);

// https server
const options = {
  key: fs.readFileSync('server_key.pem'),
  cert: fs.readFileSync('server_crt.pem')
};

https.createServer(options, function (req, res) {
  res.writeHead(200);
  res.end("hello world\n");
}).listen(8443);

