<script setup lang="ts">
import { ref, watch } from "vue"

import * as tOrder from "@/type/order"
import Modal from "../Modal.vue"

const props = defineProps<{ show: boolean, item: tOrder.Item }>()
const emit = defineEmits<{ closing: [], selected: [tOrder.CartItem] }>()

watch(() => props.show, v => {
    if (v) {
        if (props.item.variants != undefined) {
            variantOpts.value = props.item.variants
        }
    }
})

const selected = ref(0)

function handleClose() {
    selected.value = 0
    variantOpts.value = []
    emit("closing")
}

const variantOpts = ref<tOrder.ItemVariant[]>([])

function save() {
    const variant = variantOpts.value.find(opt => opt.id == selected.value)
    if (variant == undefined) return
    const payload = {
        id: props.item.id,
        name: props.item.name,
        variant
    }
    emit("selected", payload)
    handleClose()
}
</script>

<template>
    <Modal :title="props.item.name" :show="props.show" @closing="handleClose">
        <img class="block w-full h-32 mb-4" src="http://placehold.co/100">
        <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Totam odit dolorem necessitatibus! Veniam aperiam,
            ullam
            aspernatur molestiae, est eos vel sequi neque impedit architecto nam! Odit odio provident ea sed?
        </p>
        <hr class="my-2 border-primary">
        <form @submit.prevent="save" autocomplete="off">
            <div class="p-4 border border-primary rounded">
                <span class="font-bold">Choose your variant</span>
                <hr class="my-2 border-primary">
                <div class="form-control">
                    <label class="label cursor-pointer" v-for="variant in variantOpts">
                        <span class="label-text">{{ variant.name }} - S./ {{ variant.price.toFixed(2) }}</span>
                        <input class="radio" type="radio" name="radio" required :value="variant.id" v-model="selected" />
                    </label>
                </div>
            </div>
            <button class="btn btn-block btn-primary mt-8" type="submit">
                Add
            </button>
        </form>
    </Modal>
</template>
