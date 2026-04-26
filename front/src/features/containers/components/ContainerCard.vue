<script setup lang="ts">
import { computed } from 'vue'
import type { ActiveContainer } from '../api'
import { Card } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { cn } from '@/lib/utils'
import { Container } from '@lucide/vue'

const props = defineProps<{
  container: ActiveContainer
  selected?: boolean
}>()

const emit = defineEmits<{
  click: []
}>()

const shortId = computed(() => props.container.Id.slice(0, 12))
</script>

<template>
  <Card
    :class="
      cn(
        'cursor-pointer transition-all hover:ring-2 hover:ring-primary/50',
        selected && 'ring-2 ring-primary',
      )
    "
    @click="emit('click')"
  >
    <div class="p-4 flex items-center gap-3">
      <Container class="h-5 w-5 text-muted-foreground shrink-0" />
      <div class="flex-1 min-w-0">
        <div class="font-medium truncate">{{ container.Name }}</div>
      </div>
      <Badge variant="secondary" class="font-mono text-xs shrink-0">
        {{ shortId }}
      </Badge>
    </div>
  </Card>
</template>
