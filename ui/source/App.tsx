import { createEffect, createSignal, Match, Switch } from "solid-js"
import ChooseBasePathScreen from "./ChooseBasePathScreen"
import { getBasePath } from "./engine"
import GalleryScreen from "./GalleryScreen"

type Route = "loading" | "choose_base_path" | "gallery"


export default function App() {

	const [route, setRoute] = createSignal<Route>("loading")

	getBasePath().then(r => {
		if (r.basepath.length == 0) {
			setRoute("choose_base_path")
		} else {
			setRoute("gallery")
		}
	})

	return <div class="App">
		<Switch>
			<Match when={route() === "loading"}>
				<div class="LoadingEmoji">
					🌙
				</div>
			</Match>
			<Match when={route() === "choose_base_path"}>
				<ChooseBasePathScreen />
			</Match>
			<Match when={route() === "gallery"}>
				<GalleryScreen />
			</Match>
		</Switch>
	</div>
}
