import {Component} from "solid-js"
import {WeatherType} from "../lib/types"
import {formatDate} from "../lib/utils"

type Props = {
  weather: WeatherType
}

const Display: Component<Props> = ({weather}) => {
  const {location, current} = weather
  return (
    <div class="max-w-[20em] bg-white p-1 shadow rounded">
      <div class="flex justify-between mb-2 border-b-2 border-b-blue-600">
        <div>
          <p class="text-2xl">{location.country}</p>
          <p class="text-xl">{location.name}</p>
        </div>
        <img src={current.condition.icon} alt={current.condition.text} />
      </div>
      <div class="flex flex-wrap gap-3">
        <p>{current.temp_c}C</p>
        <p>{current.temp_f}F</p>
        <p>{current.condition.text} </p>
      </div>
      <div>
        <p>
          <span class="font-bold pr-2">Local time:</span>
          {formatDate(location.localtime)}
        </p>
        <p>
          <span class="font-bold pr-2">Last updated:</span>
          {formatDate(current.last_updated)}
        </p>
      </div>
    </div>
  )
}

export default Display
