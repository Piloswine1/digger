import { useQuery } from '@tanstack/vue-query'
import { getActiveContainers } from '../api'

export function useContainers() {
  return useQuery({
    queryKey: ['active-containers'],
    queryFn: getActiveContainers,
  })
}
