require("http")
  .createServer((req, res) => res.end("" + fib(40)))
  .listen(4000);

const fib = (n) => {
  if (n <= 1) return n;
  return fib(n - 1) + fib(n - 2);
};
// require("http")
//   .createServer((req, res) => {
//     require("fs").createReadStream("index.html").pipe(res);
//   })
//   .listen(3000);
