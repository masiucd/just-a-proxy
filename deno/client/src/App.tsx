import {Component, createEffect, createSignal, Show} from "solid-js"
import {z} from "zod"
import "./index.css"
import {cn} from "./lib/utils"

const Location = z.object({
  name: z.string(),
  region: z.string(),
  country: z.string(),
  lat: z.number(),
  lon: z.number(),
  tz_id: z.string(),
  localtime_epoch: z.number(),
  localtime: z.string(),
})

const Current = z.object({
  last_updated_epoch: z.number(),
  last_updated: z.string(),
  temp_c: z.number(),
  temp_f: z.number(),
  is_day: z.number(),
  condition: z.object({
    text: z.string(),
    icon: z.string(),
    code: z.number(),
  }),
})

const Weather = z.object({
  location: Location,
  current: Current,
})

// type LocationType = z.infer<typeof Location>
// type CurrentType = z.infer<typeof Current>
type WeatherType = z.infer<typeof Weather>

const App: Component = () => {
  const [search, setSearch] = createSignal("")
  const [weather, setWeather] = createSignal<WeatherType | null>(null)

  createEffect(() => {
    console.log(weather())
    console.log(search())
  })

  http: return (
    <main class="min-h-screen flex">
      <div class="max-w-2xl m-auto w-full">
        <form
          class="flex w-full items-end p-2"
          onSubmit={async e => {
            e.preventDefault()
            console.log("ENTER")
            const response = await fetch(
              `http://api.weatherapi.com/v1/current.json?key=0c883d3faadf4bc0b1d100828230801&q=${search()}&aqi=no`
            )
            const data = await response.json()
            console.log("data", data)
            setWeather(Weather.parse(data))
          }}
        >
          <div class="flex-1">
            <label for="city" class="pl-1">
              Search for your city
            </label>
            <input
              id="city"
              class="border border-slate-900 w-full min-h-[3rem] rounded-tl-md shadow outline-none pl-2"
              type="text"
              value={search()}
              onInput={e => {
                const value = e.currentTarget.value.toLowerCase()
                setSearch(value)
              }}
            />
          </div>
          <button
            class="border border-slate-900 h-[3rem] px-4 rounded-tr-md bg-blue-600 text-white"
            type="submit"
          >
            Search
          </button>
        </form>
        <div class="p-2">
          <h2 class="text-4xl">Suggested cities</h2>
          <ul class="py-2 flex flex-col gap-2 animate-in slide-in-from-top-96">
            <li>
              <button
                onClick={() => {
                  setSearch("Warszawa")
                }}
                class={cn(
                  "text-sm",
                  search() === "Warszawa"
                    ? " border-b border-blue-600 text-blue-400 "
                    : ""
                )}
              >
                Warszawa
              </button>
            </li>
            <li>
              <button
                onClick={() => {
                  setSearch("London")
                }}
                class={cn(
                  "text-sm",
                  search() === "London"
                    ? " border-b border-blue-600 text-blue-400 "
                    : ""
                )}
              >
                London
              </button>
            </li>
            <li>
              <button
                onClick={() => {
                  setSearch("Paris")
                }}
                class={cn(
                  "text-sm",
                  search() === "Paris"
                    ? " border-b border-blue-600 text-blue-400 "
                    : ""
                )}
              >
                Paris
              </button>
            </li>
          </ul>
        </div>
        <div class="p-2">
          <Show
            when={weather() !== null}
            fallback={<p class="text-3xl">No city has been selected!</p>}
          >
            <div>{weather()?.location.name}</div>
            <div>{weather()?.location.country}</div>
            <div>{weather()?.location.localtime}</div>
            <div>{weather()?.location.lat}</div>
          </Show>
        </div>
      </div>
    </main>
  )
}

export default App
