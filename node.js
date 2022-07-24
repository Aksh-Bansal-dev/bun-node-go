const express = require("express");

const app = express();

// app.use(express.static("."));

app.get("/", (req, res) => {
  res.end("" + fib(40));
});

app.listen(4002);
const fib = (n) => {
  if (n <= 1) return n;
  return fib(n - 1) + fib(n - 2);
};
