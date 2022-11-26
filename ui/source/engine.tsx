
export interface GetBasePathResponse {
	basepath: string,
}

export interface SetBasePathRequest {
	basepath: string,
}

export async function getBasePath(): Promise<GetBasePathResponse> {
	const r = await fetch("http://localhost:9001/basepath")
	return await r.json()
}

export async function setBasePath(request: SetBasePathRequest): Promise<GetBasePathResponse> {
	const r = await fetch("http://localhost:9001/basepath", { method: "PUT", body: JSON.stringify(request) })
	return await r.json()
}

export interface FileItem {
	path: string,
}

export interface GetFileTreeResponse {
	files: FileItem[],
}

export async function getFileTree(): Promise<GetFileTreeResponse> {
	const r = await fetch("http://localhost:9001/files")
	return await r.json()
}

export function getFileUri(fileItem: FileItem): string {
	return "http://localhost:9001/files/" + fileItem.path
}