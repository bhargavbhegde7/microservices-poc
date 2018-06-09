var bcrypt = require('bcryptjs');
var hashed = bcrypt.hashSync("bhargav");
console.log(hashed);
