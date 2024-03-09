<script setup lang="ts">
import { reactive, ref } from "vue"

import * as tProduct from "@/type/product"
import productService from "@/service/product"
import useToast from "@/composables/useToast";
import Modal from "../Modal.vue";
import Overlay from "../Overlay.vue";

const toast = useToast()

const props = defineProps<{ show: boolean }>()
const emit = defineEmits<{ closing: [], refresh: [] }>();

const loading = ref(false)
function handleClose() {
    loading.value = false
    form.name = ""
    emit("closing")
}

const categoryOpts = ref<tProduct.Category[]>([])
async function getOpts() {
    loading.value = true
    const res = await productService.listCategory()
    categoryOpts.value = res.rows
    loading.value = false
}
getOpts()

const form = reactive({ name: "" })
async function create() {
    loading.value = true
    await productService.createCategory({ name: form.name })
    loading.value = false
    toast.display({ message: "Category created", variant: "success" })
    handleClose()
    emit("refresh")
}
</script>

<template>
    <Modal title="Create Category" :show="props.show" @closing="handleClose">
        <Overlay :show="loading">
            <form @submit.prevent="create" autocomplete="off">
                <label class="form-control w-full mb-8">
                    <div class="label">
                        <span class="label-text">Name</span>
                    </div>
                    <input class="input input-bordered w-full" type="text" required v-model="form.name" />
                </label>

                <button class="btn btn-block btn-primary" type="submit">
                    Create
                </button>
            </form>
        </Overlay>
    </Modal>
</template>
