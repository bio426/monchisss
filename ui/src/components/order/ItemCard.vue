<script setup lang="ts">
import { } from "vue"
import { PlusIcon } from "@heroicons/vue/24/solid";

import * as tOrder from "@/type/order"
import util from "@/util/common"

const props = defineProps<{ item: tOrder.Item }>()
const emit = defineEmits<{ selected: [tOrder.Item], showVariantSelector: [tOrder.Item] }>()

function selectItem() {
    if (props.item.type == "variant") {
        emit('showVariantSelector', props.item)
        return
    }
    emit('selected', props.item)
}
</script>

<template>
    <div class="text-center shadow shadow-black">
        <img class=" w-full mb-4" src="https://placehold.co/100">
        <span class="block mb-8 text-lg font-semibold ">{{ util.capitalizePhrase(props.item.name) }}</span>
        <span class="block mb-4" v-if="props.item.type == 'simple'">S/. {{ props.item.price?.toFixed(2) }}</span>
        <span class="block mb-4" v-else-if="props.item.type == 'variant' && props.item.variants != undefined">
            S/. {{ props.item.variants[0].price.toFixed(2) }} - {{ props.item.variants[props.item.variants.length -
                1].price.toFixed(2) }}
        </span>
        <button class="btn btn-primary btn-block" @click="selectItem">
            <PlusIcon class="w-6 h-6" />
        </button>
    </div>
</template>
