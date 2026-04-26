<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useContainerLogs } from '../composables/useContainerLogs'
import { Card } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Checkbox } from '@/components/ui/checkbox'
import { Button } from '@/components/ui/button'
import { Download, Loader2 } from '@lucide/vue'

const props = defineProps<{
  containerId: string
}>()

const tail = ref(100)
const stderr = ref(false)
const autoRefresh = ref(false)
const refreshInterval = ref(5000)
const filter = ref('')

watch(refreshInterval, (val) => {
  if (val < 1000) {
    refreshInterval.value = 1000
  }
})

const {
  data: logs,
  isLoading,
  isError,
  error,
} = useContainerLogs(() => props.containerId, tail, stderr, autoRefresh, refreshInterval)

const filteredLogs = computed(() => {
  if (!logs.value) return ''
  if (!filter.value) return logs.value
  const filterLower = filter.value.toLowerCase()
  return logs.value
    .split('\n')
    .filter((line) => line.toLowerCase().includes(filterLower))
    .join('\n')
})

async function downloadLogs() {
  if (!logs.value) return
  try {
    const text = logs.value
    const blob = new Blob([text], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `container-logs-${props.containerId.slice(0, 12)}.txt`
    a.click()
    URL.revokeObjectURL(url)
  } catch {
    // download failed silently
  }
}

watch(
  () => props.containerId,
  () => {
    filter.value = ''
  },
)
</script>

<template>
  <Card class="flex flex-col overflow-hidden" style="max-height: calc(100vh - 8rem)">
    <div class="p-4 border-b space-y-3 shrink-0">
      <div class="flex items-center gap-4 flex-wrap">
        <div class="flex items-center gap-2">
          <label class="text-sm font-medium whitespace-nowrap">Tail</label>
          <Input v-model.number="tail" type="number" min="1" max="10000" class="w-24" />
        </div>
        <Checkbox v-model="stderr" label="stderr" />
        <Checkbox v-model="autoRefresh" label="Auto-refresh" />
        <div v-if="autoRefresh" class="flex items-center gap-2">
          <label class="text-sm font-medium whitespace-nowrap">Interval (ms)</label>
          <Input
            v-model.number="refreshInterval"
            type="number"
            min="1000"
            step="500"
            class="w-28"
          />
        </div>
      </div>
      <div class="flex items-center gap-3">
        <Input v-model="filter" placeholder="Filter logs..." class="flex-1" />
        <Button variant="outline" size="sm" @click="downloadLogs">
          <Download class="h-4 w-4" />
          Download
        </Button>
      </div>
    </div>

    <div class="flex-1 overflow-auto p-4">
      <div v-if="isLoading" class="flex items-center justify-center py-8 text-muted-foreground">
        <Loader2 class="h-5 w-5 animate-spin mr-2" />
        Loading logs...
      </div>
      <div v-else-if="isError" class="text-destructive text-sm">
        Failed to load logs: {{ (error as Error)?.message }}
      </div>
      <pre v-else class="text-sm font-mono whitespace-pre-wrap break-all">{{
        filteredLogs || '(no output)'
      }}</pre>
    </div>
  </Card>
</template>
