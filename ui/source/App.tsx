import styles from "./App.module.css"

export default function App() {
	return <div class={styles.App}>
		<header class={styles.header}>
			<img src="https://images.unsplash.com/photo-1600716051809-e997e11a5d52?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1450&q=80" class={styles.logo} alt="logo" />
			<p>
				Edit <code>src/App.tsx</code> and save to reload.
			</p>
			<a
				class={styles.link}
				href="https://github.com/solidjs/solid"
				target="_blank"
				rel="noopener noreferrer"
			>
				Learn Solid
			</a>
		</header>
	</div>

}
