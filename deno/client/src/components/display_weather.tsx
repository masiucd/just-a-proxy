import {Component} from "solid-js"

import {WeatherType} from "../lib/types"
import Display from "./display"

type Props = {
  weather: WeatherType
}

const DisplayWeather: Component<Props> = props => {
  // eslint-disable-next-line solid/reactivity
  const {weather} = props
  return <Display weather={weather} />
}

export default DisplayWeather
