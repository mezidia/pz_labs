const http = require('http');

const server = http.createServer().listen('8080', () => console.log('listening on port 8080'));

server.on('request', (req,res) => {
  if (req.url === '/forums') {
    const data = JSON.stringify('forums returned from server');
    console.log(req.url, data);
    res.write(data);
    res.end();
  } else if (req.url === '/registrateUser') {
    const data = JSON.stringify('user created');
    console.log(req.url, data);
    res.write(data);
    res.end();
  }
})