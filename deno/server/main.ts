import {Application, Router} from "https://deno.land/x/oak@v11.1.0/mod.ts"
import {config} from "https://deno.land/x/dotenv@v3.2.0/mod.ts"
import {getQuery} from "https://deno.land/x/oak@v11.1.0/helpers.ts"
import {oakCors} from "https://deno.land/x/cors@v1.2.2/mod.ts"

const router = new Router()
const PORT = Deno.env.get("PORT") || 8000

const API_KEY_NAME = config().API_KEY_NAME
const API_KEY_VALUE = config().API_KEY_VALUE
const WEB_URL = config().WEB_URL

router.get("/api/:city", async ctx => {
  try {
    const query = getQuery(ctx, {mergeParams: true})
    ctx.response.type = "application/json"
    const res =
      await fetch(`${WEB_URL}/v1/current.json?${API_KEY_NAME}=${API_KEY_VALUE}&q=${query["city"]}&aqi=yes
      `)

    if (res.ok) {
      ctx.response.status = 200
      const data = await res.json()
      console.log("data", data)
      ctx.response.body = JSON.stringify({status: 200, data})
    } else {
      ctx.response.status = 404
      ctx.response.body = JSON.stringify({status: 404, error: "Not found"})
    }
  } catch (error) {
    console.error("error", error)
    ctx.response.status = 500
    ctx.response.body = JSON.stringify({status: 500, error: error})
  }
})

const app = new Application()

app.use(oakCors()) // Enable CORS for All Routes
app.use(router.routes())
app.use(router.allowedMethods())

console.log(`Server running on port ${PORT} 🚀`)
await app.listen({port: Number(PORT)})
