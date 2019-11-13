const Koa = require("koa");
// init koa
const app = new Koa();

//  all incoming requests intercepted and response is Hello World
app.use(async ctx => (ctx.body = "Hello World"));
//  init port
const PORT = process.env.PORT || 3000;
//  start listening on port
app.listen(PORT, () => console.log(`Server listening on port: ${PORT}`));
