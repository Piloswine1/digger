import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { restartContainer, stopContainer, startContainer } from '../api'

export function useContainerActions() {
  const queryClient = useQueryClient()

  const refetch = () => {
    queryClient.invalidateQueries({queryKey: ['active-containers']})
  }

  const restart = useMutation({
    mutationFn: restartContainer,
    onSuccess: refetch,
  })

  const stop = useMutation({
    mutationFn: stopContainer,
    onSuccess: refetch,
  })

  const start = useMutation({
    mutationFn: startContainer,
    onSuccess: refetch,
  })

  return {restart, stop, start}
}
