<script setup lang="ts">
import { computed } from "vue"
import { TrashIcon } from "@heroicons/vue/24/solid";

import * as tOrder from "@/type/order"

const props = defineProps<{ item: tOrder.CartItem }>()

const isSimple = computed(() => props.item.price != undefined)
const isVariant = computed(() => props.item.variant != undefined)
</script>

<template>
    <div class="p-4 shadow-lg border border-neutral">
        <div class="flex justify-between">
            <div>
                <span class="uppercase">{{ props.item.name }}</span>
            </div>
            <div v-if="isVariant">
                <span class="text-sm">{{ props.item.variant?.name }}</span>
            </div>
        </div>
        <hr class="border-primary my-2">
        <div class="flex justify-between">
            <div>
                <button class="btn btn-ghost btn-square btn-sm">
                    <TrashIcon class="w-5 h-5" />
                </button>
            </div>
            <div>
                <span class="font-bold" v-if="isSimple">S/. {{ props.item.price?.toFixed(2) }}</span>
                <span class="font-bold" v-else-if="isVariant">S/. {{ props.item.variant?.price.toFixed(2) }}</span>
            </div>
        </div>
    </div>
</template>
