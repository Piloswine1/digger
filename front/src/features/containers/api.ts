export interface ContainerInfo {
  Id: string
  Name: string
  Status: string
}

export type ContainersMode = 'all' | 'active'
export async function getActiveContainers(
  mode: ContainersMode,
  signal?: AbortSignal,
): Promise<ContainerInfo[]> {
  const params = new URLSearchParams()
  if (mode === 'all') {
    params.set('all', 'true')
  }
  const res = await fetch(
    `/api/v1/containers?${params.toString()}`,
    {signal},
  )
  if (!res.ok) {
    throw new Error(`Failed to fetch containers: ${res.statusText}`)
  }
  return res.json()
}

export async function getContainerLogs(
  id: string,
  limit: number,
  stderr: boolean,
  signal?: AbortSignal,
): Promise<string> {
  const params = new URLSearchParams()
  if (limit) {
    params.set('limit', String(limit))
  }
  if (stderr) {
    params.set('stderr', 'true')
  }
  const res = await fetch(
    `/api/v1/containers/${encodeURIComponent(id)}/logs?${params.toString()}`,
    {signal},
  )
  if (!res.ok) {
    throw new Error(`Failed to fetch logs: ${res.statusText}`)
  }
  return res.text()
}

export async function restartContainer(id: string): Promise<string> {
  const res = await fetch(
    `/api/v1/containers/${encodeURIComponent(id)}/restart`,
    {method: 'POST'},
  )
  if (!res.ok) {
    throw new Error(`Failed to restart container: ${res.statusText}`)
  }
  return res.text()
}

export async function stopContainer(id: string): Promise<string> {
  const res = await fetch(
    `/api/v1/containers/${encodeURIComponent(id)}/stop`,
    {method: 'POST'},
  )
  if (!res.ok) {
    throw new Error(`Failed to stop container: ${res.statusText}`)
  }
  return res.text()
}

export async function startContainer(id: string): Promise<string> {
  const res = await fetch(
    `/api/v1/containers/${encodeURIComponent(id)}/start`,
    {method: 'POST'},
  )
  if (!res.ok) {
    throw new Error(`Failed to start container: ${res.statusText}`)
  }
  return res.text()
}
