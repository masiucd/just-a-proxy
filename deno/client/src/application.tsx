import "./index.css"

import {Component, createSignal, Show} from "solid-js"

import DisplayWeather from "./components/display_weather"
import Form from "./components/form"
import Suggestions from "./components/suggestions"
import {WeatherType} from "./lib/types"

const App: Component = () => {
  const [search, setSearch] = createSignal("")
  const [errorMessage, setErrorMessage] = createSignal<null | string>(null)
  const [weather, setWeather] = createSignal<WeatherType | null>(null)

  return (
    <main class="min-h-screen flex bg-gradient-to-r from-slate-100 to-blue-200">
      <div class="max-w-2xl m-auto w-full">
        <Form
          search={search}
          setSearch={setSearch}
          setWeather={setWeather}
          setErrorMessage={setErrorMessage}
        />
        <div class="p-2">
          <h2 class="text-4xl">Suggested cities</h2>
          <Suggestions search={search} setSearch={setSearch} />
        </div>
        <div class="p-2 h-[20em] max-w-[24em]">
          <Show
            when={weather() !== null && errorMessage() === null}
            fallback={null}
          >
            <DisplayWeather weather={weather() as WeatherType} />
          </Show>
          <Show when={errorMessage() !== null} fallback={null}>
            <p class="text-red-500">{errorMessage()}</p>
          </Show>
        </div>
      </div>
    </main>
  )
}

export default App
