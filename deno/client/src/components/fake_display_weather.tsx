import Display from "./Display"

let location = {
  name: "Warszawa",
  region: "",
  country: "Poland",
  lat: 52.25,
  lon: 21,
  tz_id: "Europe/Warsaw",
  localtime_epoch: 1673178230,
  localtime: "2023-01-08 12:43",
}

let current = {
  last_updated_epoch: 1673177400,
  last_updated: "2023-01-08 12:30",
  temp_c: 3,
  temp_f: 37.4,
  is_day: 1,
  condition: {
    text: "Partly cloudy",
    icon: "//cdn.weatherapi.com/weather/64x64/day/116.png",
    code: 1003,
  },
}
function FakeDisplayWeather() {
  return <Display weather={{location, current}} />
}

export default FakeDisplayWeather
