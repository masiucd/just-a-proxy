import {Accessor, Component} from "solid-js"

import {Weather, WeatherType} from "../lib/types"

type Props = {
  setWeather: (value: WeatherType | null) => void
  setErrorMessage: (message: string | null) => void
  setSearch: (value: string) => void
  search: Accessor<string>
}
// eslint-disable-next-line solid/no-destructure
const Form: Component<Props> = ({
  setWeather,
  setErrorMessage,
  setSearch,
  search,
}) => (
  <form
    class="flex w-full items-end p-2"
    onSubmit={async e => {
      e.preventDefault()
      setErrorMessage(null)
      setWeather(null)
      const response = await fetch(`http://localhost:8000/api/${search()}`)
      if (response.ok) {
        const {data} = await response.json()
        setWeather(Weather.parse(data))
      } else {
        // handle error
        setErrorMessage("City not found")
      }
    }}
  >
    <div class="flex-1">
      <label for="city" class="pl-1 flex ">
        <p class="text-2xl mb-1">Search for your city</p>
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
)

export default Form
