import { createSignal, For } from "solid-js"
import { FileItem, getFileTree, getFileUri } from "./engine"

export default function GalleryScreen() {
	const [filePathList, setFilePathList] = createSignal<FileItem[]>([])

	getFileTree().then(resp => {
		setFilePathList(resp.files)
	})

	return <div class="Gallery">
		<div class="GalleryFileList">
			<For each={filePathList()}>{(fileItem) => (
				<div class="GalleryItem">
					<img class="GalleryItem" src={getFileUri(fileItem)} />
					<span>{fileItem.path}</span>
				</div>
			)}</For>
		</div>
	</div>
}