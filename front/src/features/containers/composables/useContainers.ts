import { toValue, type MaybeRefOrGetter } from 'vue'
import { useQuery } from '@tanstack/vue-query'

import { getActiveContainers, type ContainersMode } from '../api'

export function useContainers(mode: MaybeRefOrGetter<ContainersMode>) {
  return useQuery({
    queryKey: ['active-containers', mode],
    queryFn: ({signal}) =>
      getActiveContainers(toValue(mode), signal),
  })
}
