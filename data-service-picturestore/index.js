var express = require('express');
var app = express();
var cors = require('cors');
app.use(cors({origin: 'http://localhost:8080'}));
var multer = require('multer')().single();
app.use(multer);
var request = require('request');

var variable = {}

app.get('/', function (req, res) {

  variable.name = false;
  variable.email = false;
  
  request({
    headers: {
      'x-access-token': req.headers['x-access-token']
    },
    //uri: 'http://auth-service-picturestore:3000/api/auth/verify-token',
    uri: 'http://localhost:3000/api/auth/verify-token',
    method: 'GET'
  }, function (err, response, body) {
	  if (!err && response.statusCode == 200) {

		  var jsonData = JSON.parse(body);
		  variable.name = jsonData.name;
	    variable.email = jsonData.email;	

		  res.send(variable);
	  }
	  else{
      // Failed to verify the token
		  res.send(variable);
	  }		
  });
});

app.listen(process.env.PORT || 4000, function () {
  console.log('Example app listening on port 4000!');
});
