const http = require("http");

const request = http.request({
  host : "localhost",
  port : process.env.PORT || '3000',
  timeout : 2000
}, (res) => process.exit(res.statusCode !== 200));

request.on('error', () => process.exit(1));

request.end();
