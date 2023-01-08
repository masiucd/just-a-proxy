import {Component, createEffect, createSignal, For, Show} from "solid-js"
import DisplayWeather from "./components/display_weather"
import FakeDisplayWeather from "./components/fake_display_weather"
import Form from "./components/form"
import Suggestions from "./components/suggestions"
import "./index.css"
import {Weather, WeatherType} from "./lib/types"
import {cn} from "./lib/utils"

const App: Component = () => {
  const [search, setSearch] = createSignal("")
  const [weather, setWeather] = createSignal<WeatherType | null>(null)

  http: return (
    <main class="min-h-screen flex bg-gradient-to-r from-slate-100 to-blue-200">
      <div class="max-w-2xl m-auto w-full">
        <Form search={search} setSearch={setSearch} setWeather={setWeather} />
        <div class="p-2">
          <h2 class="text-4xl">Suggested cities</h2>
          <Suggestions search={search} setSearch={setSearch} />
        </div>
        <div class="p-2">
          <Show when={weather() !== null} fallback={<FakeDisplayWeather />}>
            <DisplayWeather weather={weather() as WeatherType} />
          </Show>
        </div>
      </div>
    </main>
  )
}

export default App
