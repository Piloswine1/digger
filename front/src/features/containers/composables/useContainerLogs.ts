import { useQuery } from '@tanstack/vue-query'
import { toValue, type MaybeRefOrGetter } from 'vue'
import { getContainerLogs } from '../api'

export function useContainerLogs(
  id: MaybeRefOrGetter<string>,
  limit: MaybeRefOrGetter<number>,
  stderr: MaybeRefOrGetter<boolean>,
  autoRefresh: MaybeRefOrGetter<boolean>,
  refreshInterval: MaybeRefOrGetter<number>,
) {
  return useQuery({
    queryKey: ['container-logs', id, limit, stderr] as const,
    queryFn: () => getContainerLogs(toValue(id), toValue(limit), toValue(stderr)),
    refetchInterval: () => (toValue(autoRefresh) ? toValue(refreshInterval) : false),
    enabled: () => !!toValue(id),
  })
}
