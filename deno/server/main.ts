import {Application, Router} from "https://deno.land/x/oak@v11.1.0/mod.ts"

const router = new Router()
const PORT = Deno.env.get("PORT") || 8000

router
  .get("/", ctx => {
    ctx.response.body = `<h1> Hello World </h1>`
  })
  .get("/users", ctx => {
    ctx.response.type = "application/json"
    ctx.response.body = JSON.stringify([
      {name: "John", age: 20},
      {name: "Jane", age: 21},
    ])
  })

const app = new Application()
app.use(router.routes())
app.use(router.allowedMethods())

console.log(`Server running on port ${PORT} 🚀`)
await app.listen({port: Number(PORT)})
