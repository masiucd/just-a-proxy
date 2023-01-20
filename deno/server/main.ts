import {Application, Router} from "https://deno.land/x/oak@v11.1.0/mod.ts"
import {config} from "https://deno.land/x/dotenv@v3.2.0/mod.ts"
// import {getQuery} from "https://deno.land/x/oak@v11.1.0/helpers.ts"

const router = new Router()
const PORT = Deno.env.get("PORT") || 8000

console.log(config().API_KEY)
console.log(config().WEB_URL)

router
  .get("/", (ctx) => {
    ctx.response.body = `<h1> Hello World </h1>`
  })
  // .get("/users", ctx => {
  //   ctx.response.type = "application/json"
  //   ctx.response.body = JSON.stringify([
  //     {name: "John", age: 20},
  //     {name: "Jane", age: 21},
  //   ])
  .get("/users", async (ctx) => {
    const res = await fetch("https://jsonplaceholder.typicode.com/users")
    const data = await res.json()
    ctx.response.type = "application/json"
    ctx.response.body = JSON.stringify(data)
  })
  .get("/users/:id", async (ctx) => {
    const res = await fetch(
      `https://jsonplaceholder.typicode.com/users/${ctx.params.id}`
    )
    const data = await res.json()
    // getQuery(ctx, {mergeParams: true})
    console.log("ctx.params", ctx.params.name)
    ctx.response.type = "application/json"
    ctx.response.body = JSON.stringify({data: data})
  })
  .get("/city/:name", async (ctx) => {
    // TODO: Add a query param to the url
    console.log("ctx.params", ctx.params.name)
    const res = await fetch(
      `${config().WEB_URL}/v1/current.json?key=${config().API_KEY}&q=${
        ctx.params.name
      }&aqi=no`
    )
    // set status code
    if (res.ok) {
      ctx.response.status = 200
      const data = await res.json()
      console.log("data", data)
    }
  })

const app = new Application()
app.use(router.routes())
app.use(router.allowedMethods())

console.log(`Server running on port ${PORT} 🚀`)
await app.listen({port: Number(PORT)})
