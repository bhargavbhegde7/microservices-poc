var express = require('express');
var router = express.Router();
var User = require('../user/User');

var jwt = require('jsonwebtoken');
var bcrypt = require('bcryptjs');
var config = require('../config');

var cors = require('cors');
router.use(cors({origin: ['http://localhost:8080']}));

var multer = require('multer')().single();
router.use(multer);

var VerifyToken = require('./VerifyToken');

router.post('/register', function(req, res) {
  if(req.body.password1 === req.body.password2){
    var hashedPassword = bcrypt.hashSync(req.body.password1, 8);
    User.create({
      name : req.body.name,
      email : req.body.email,
      password : hashedPassword
    },
    function (err, user) {
      if (err) return res.status(500).send("There was a problem registering the user.")
      // create a token
      var token = jwt.sign({ id: user._id }, config.secret, {
        expiresIn: 86400 // expires in 24 hours
      });
      res.status(200).send({ auth: true, token: token });
    });
  }
  else{
    res.status(401).send({ auth: false, token: null });
  }
});

router.post('/login', function(req, res) {
  User.findOne({ email: req.body.email }, function (err, user) {
    if (err) return res.status(500).send('Error on the server.');
    if (!user) return res.status(404).send('No user found.');
    var passwordIsValid = bcrypt.compareSync(req.body.password, user.password);
    if (!passwordIsValid) return res.status(401).send({ auth: false, token: null });
    var token = jwt.sign({ id: user._id }, config.secret, {
      expiresIn: 86400 // expires in 24 hours
    });
    res.status(200).send({ auth: true, token: token });
  });
});

/*router.get('/me', function(req, res) {
  var token = req.headers['x-access-token'];
  if (!token) return res.status(401).send({ auth: false, message: 'No token provided.' });
  
  jwt.verify(token, config.secret, function(err, decoded) {
    if (err) return res.status(500).send({ auth: false, message: 'Failed to authenticate token.' });
    
    User.findById(decoded.id, 
	  { password: 0 }, // projection
	  function (err, user) {
	    if (err) return res.status(500).send("There was a problem finding the user.");
	    if (!user) return res.status(404).send("No user found.");
	    res.status(200).send(user);
	});
  });
});*/

router.get('/verify-token', VerifyToken, function(req, res, next) {
  User.findById(req.userId, { password: 0 }, function (err, user) {
    if (err) return res.status(500).send("There was a problem finding the user.");
    if (!user) return res.status(404).send("No user found.");
    
    res.status(200).send(user);
  });
});

module.exports = router;