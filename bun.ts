Bun.serve({
  fetch(req: Request) {
    return new Response("" + fib(40));
  },
  port: 4000,
});

const fib = (n: number): number => {
  if (n <= 1) return n;
  return fib(n - 1) + fib(n - 2);
};
