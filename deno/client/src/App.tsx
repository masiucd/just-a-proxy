import {Component, createEffect, createSignal, For, onMount} from "solid-js"
import {z} from "zod"
import "./index.css"

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

type LocationType = z.infer<typeof Location>
type CurrentType = z.infer<typeof Current>
type WeatherType = z.infer<typeof Weather>

const App: Component = () => {
  const [weather, setWeather] = createSignal<WeatherType | null>(null)
  onMount(async () => {
    const response = await fetch(
      "http://api.weatherapi.com/v1/current.json?key=0c883d3faadf4bc0b1d100828230801&q=London&aqi=no"
    )
    const data = await response.json()
    setWeather(Weather.parse(data))
    console.log(data)
  })

  return (
    <main class="min-h-screen flex">
      <div class="max-w-2xl m-auto w-full">
        <div>
          <label for="city" class="pl-1">
            Search for your city
          </label>
          <input
            id="city"
            class="border border-slate-900 w-full min-h-[3rem] rounded-md shadow"
            type="text"
            onInput={e => {
              const value = e.currentTarget.value.toLowerCase()
            }}
          />
        </div>
      </div>
    </main>
  )
}

export default App
