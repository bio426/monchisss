<script setup lang="ts">
import { reactive, ref, watch } from "vue"
import { TrashIcon } from "@heroicons/vue/24/solid";

import * as tProduct from "@/type/product"
import Modal from "../Modal.vue";

const props = defineProps<{ show: boolean }>()
const emit = defineEmits<{ closing: [] }>();
const modelValue = defineModel<tProduct.ProductVariantUnsaved[]>({ required: true })

watch(() => props.show, v => {
    if (v) {
        variants.value = [...modelValue.value]
    }
})
function handleClose() {
    form.name = ""
    form.price = 0
    variants.value = []
    emit("closing")
}


const form = reactive({ name: "", price: 0 })
function addVariant() {
    variants.value.push({ name: form.name, price: form.price })
    form.name = ""
    form.price = 0
}
function removeVariant(idx: number) {
    variants.value.splice(idx, 1)
}

const variants = ref<tProduct.ProductVariantUnsaved[]>([])
async function save() {
    modelValue.value = [...variants.value]
    handleClose()
}
</script>

<template>
    <Modal title="Manage Product Variants" :show="props.show" @closing="handleClose">
        <form @submit.prevent="addVariant" autocomplete="off">
            <div class="grid grid-cols-2 gap-4">
                <label class="form-control w-full mb-4">
                    <div class="label">
                        <span class="label-text">Name</span>
                    </div>
                    <input class="input input-bordered w-full" type="text" required v-model="form.name" />
                </label>
                <label class="form-control w-full mb-4">
                    <div class="label">
                        <span class="label-text">Price</span>
                    </div>
                    <input class="input input-bordered w-full" type="number" step=".01" min="0.01" required
                        v-model="form.price" />
                </label>
            </div>
            <button class="btn btn-block btn-primary" type="submit">
                Add
            </button>
        </form>
        <hr class="my-2">
        <div class="overflow-x-auto">
            <table class="table whitespace-nowrap">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Price</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(variant, idx) in variants">
                        <td>
                            {{ variant.name }}
                        </td>
                        <td>
                            {{ variant.price.toFixed(2) }}
                        </td>
                        <td>
                            <button class="btn btn-sm btn-square" title="Delete" @click="removeVariant(idx)">
                                <TrashIcon class="w-4 fill-error" />
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <button class="btn btn-block btn-primary mt-8" @click="save">
            Save
        </button>
    </Modal>
</template>
