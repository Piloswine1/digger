<script setup lang="ts">
import { computed } from 'vue'
import { Card } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'
import { Container, RotateCw, Square, Play } from '@lucide/vue'

import type { ContainerInfo } from '../api'
import { useContainerActions } from '../composables/useContainerActions'

const props = defineProps<{
  container: ContainerInfo
}>()

const containerStatus = computed(() => {
  if (props.container.Status.startsWith("Exited")) {
    return "Exited"
  }
  return props.container.Status;
})

const {restart, stop, start} = useContainerActions()

const shortId = computed(() => props.container.Id.slice(0, 12))
const containerVariant = computed(() => {
  const stat = props.container.Status;
  if (stat.startsWith('Up')) return 'success';
  if (stat.startsWith('Exited')) return 'error';
  return 'secondary';
})
const isRunning = computed(() => props.container.Status.startsWith('Up'))
const selected = defineModel<string>();
</script>

<template>
  <Card
    :class="
      cn(
        'cursor-pointer transition-all hover:ring-2 hover:ring-primary/50 ml-1 mr-3',
        selected === container.Id && 'ring-2 ring-primary',
      )
    "
    @click="selected = container.Id"
  >
    <div class="p-4 flex items-center gap-3">
      <Container class="h-5 w-5 text-muted-foreground shrink-0" />
      <div class="flex-1 min-w-0" :title="container.Name">
        <div class="font-medium break-all">{{ container.Name }}</div>
      </div>
      <div class="flex flex-col items-end gap-1">
        <Badge variant="secondary" class="font-mono text-xs shrink-0">
          {{ shortId }}
        </Badge>
        <Badge
          :variant="containerVariant"
          class="font-mono text-xs shrink-0"
          :title="container.Status"
        >
          {{ containerStatus }}
        </Badge>
      </div>
      <div class="flex items-center gap-0.5 shrink-0" @click.stop>
        <template v-if="isRunning">
          <Button
            variant="ghost"
            size="icon-sm"
            title="Restart"
            :disabled="restart.isPending.value"
            @click="restart.mutate(container.Id)"
          >
            <RotateCw />
          </Button>
          <Button
            variant="ghost"
            size="icon-sm"
            title="Stop"
            :disabled="stop.isPending.value"
            @click="stop.mutate(container.Id)"
          >
            <Square />
          </Button>
        </template>
        <Button
          v-else
          variant="ghost"
          size="icon-sm"
          title="Start"
          :disabled="start.isPending.value"
          @click="start.mutate(container.Id)"
        >
          <Play />
        </Button>
      </div>
    </div>
  </Card>
</template>
