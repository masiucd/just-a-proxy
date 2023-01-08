import {Component} from "solid-js"
import {WeatherType} from "../lib/types"
import Display from "./Display"

type Props = {
  weather: WeatherType
}

const DisplayWeather: Component<Props> = ({weather}) => {
  return <Display weather={weather} />
}

export default DisplayWeather
