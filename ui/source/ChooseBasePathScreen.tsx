import { createEffect, createSignal } from "solid-js"

export default function ChooseBasePathScreen() {
	const [basePath, setBasePath] = createSignal("")

	createEffect(() => {
		fetch("http://localhost:9001/basepath")
			.then(resp => resp.json())
			.then(resp => setBasePath(resp.basepath))
	})

	const handleClick = function () {
		const req = { basePath: basePath() }
		fetch("http://localhost:9001/basepath", { method: "PUT", body: JSON.stringify(req) })
			.then(resp => resp.json())
			.then(resp => setBasePath(resp.basepath))
	}

	return <div class="BasePathEditor">
		<input
			type="text"
			placeholder="base path"
			value={basePath()}
			onInput={(e: InputEvent) => setBasePath(e.target instanceof HTMLInputElement ? e.target.value : "")}
		/>
		<input
			type="button"
			value="Update"
			onClick={handleClick}
		/>
	</div>
}