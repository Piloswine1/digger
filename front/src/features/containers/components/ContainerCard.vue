<script setup lang="ts">
import { computed } from 'vue'
import type { ContainerInfo } from '../api'
import { Card } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { cn } from '@/lib/utils'
import { Container } from '@lucide/vue'

const props = defineProps<{
  container: ContainerInfo
}>()

const shortId = computed(() => props.container.Id.slice(0, 12))
const containerVariant = computed(() => {
  const stat = props.container.Status;
  if (stat.startsWith('Up')) return 'success';
  if (stat.startsWith('Exited')) return 'error';
  return 'secondary';
})
const selected = defineModel<string>();
</script>

<template>
  <Card
    :class="
      cn(
        'cursor-pointer transition-all hover:ring-2 hover:ring-primary/50',
        selected === container.Id && 'ring-2 ring-primary',
      )
    "
    @click="selected = container.Id"
  >
    <div class="p-4 flex items-center gap-3">
      <Container class="h-5 w-5 text-muted-foreground shrink-0" />
      <div class="flex-1 min-w-0">
        <div class="font-medium truncate">{{ container.Name }}</div>
      </div>
      <div class="flex flex-col items-end gap-1">
        <Badge variant="secondary" class="font-mono text-xs shrink-0">
          {{ shortId }}
        </Badge>
        <Badge
          :variant="containerVariant"
          class="font-mono text-xs shrink-0"
        >
          {{ container.Status }}
        </Badge>
      </div>
    </div>
  </Card>
</template>
