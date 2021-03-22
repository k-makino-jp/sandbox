const fs = require('fs');
var https = require('https');

// https server
const options = {
  key: fs.readFileSync('server_key.pem'),
  cert: fs.readFileSync('server_crt.pem')
};

https.createServer(options, function (req, res) {
  res.writeHead(200);
  res.end("hello world\n");
}).listen(8443);

