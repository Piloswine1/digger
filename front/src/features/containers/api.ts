export interface ActiveContainer {
  Id: string
  Name: string
}

export async function getActiveContainers(): Promise<ActiveContainer[]> {
  const res = await fetch('/api/v1/containers')
  if (!res.ok) {
    throw new Error(`Failed to fetch containers: ${res.statusText}`)
  }
  return res.json()
}

export async function getContainerLogs(
  id: string,
  limit: number,
  stderr: boolean,
): Promise<string> {
  const params = new URLSearchParams()
  if (limit) {
    params.set('limit', String(limit))
  }
  if (stderr) {
    params.set('stderr', 'true')
  }
  const res = await fetch(`/api/v1/containers/${encodeURIComponent(id)}/logs?${params.toString()}`)
  if (!res.ok) {
    throw new Error(`Failed to fetch logs: ${res.statusText}`)
  }
  return res.text()
}
