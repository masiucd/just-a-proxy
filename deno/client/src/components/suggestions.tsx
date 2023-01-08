import {Accessor, Component, For} from "solid-js"

import {cn} from "../lib/utils"

const suggested = Object.freeze(["Warszawa", "London", "Paris"])

type Props = {
  search: Accessor<string>
  setSearch: (value: string) => void
}

const Suggestions: Component<Props> = props => {
  return (
    <ul class="py-2 flex gap-2 animate-in slide-in-from-top-96">
      <For each={suggested}>
        {city => (
          <li>
            <button
              onClick={() => {
                props.setSearch(city)
              }}
              class={cn(
                "text-sm hover:text-blue-600 p-1 border border-slate-900 rounded-sm shadow",
                props.search() === city ? " border-blue-600 bg-blue-400 " : ""
              )}
            >
              {city}
            </button>
          </li>
        )}
      </For>
    </ul>
  )
}

export default Suggestions
