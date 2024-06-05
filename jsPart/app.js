const express = require('express');
const app = express();
const port = 8080;

app.get('/', (req, res) => {
  res.send('Hello, This is my last part of lab3!!');
});

app.listen(port, () => {
  console.log(`App listening at http://localhost:${port}`);
});
