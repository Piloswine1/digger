<script setup lang="ts">
import type { PrimitiveProps } from 'reka-ui'
import type { BadgeVariants } from '.'
import { Primitive } from 'reka-ui'
import { cn } from '@/lib/utils'
import { badgeVariants } from '.'

const props = defineProps<
  PrimitiveProps & {
    variant?: BadgeVariants['variant']
  }
>()
</script>

<template>
  <Primitive
    data-slot="badge"
    :class="cn(badgeVariants({ variant }), $attrs.class)"
    v-bind="props"
  >
    <slot />
  </Primitive>
</template>
