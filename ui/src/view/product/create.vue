<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router";
import { PlusIcon, PencilSquareIcon } from "@heroicons/vue/24/solid";

import * as tProduct from "@/type/product"
import productService from "@/service/product"
import useToast from "@/composables/useToast";
import Overlay from "@/components/Overlay.vue";
import CreateCategoryModal from "@/components/product/CreateCategoryModal.vue";
import ManageVariantsModal from "@/components/product/ManageVariantsModal.vue";

const router = useRouter()
const toast = useToast()

const loading = ref(false)

const form = reactive<{
    name: string
    type: string
    price: number
    category: number
    variants: tProduct.ProductVariantUnsaved[]
}>({ name: "", type: "", price: 0, category: 0, variants: [] })

const createCategory = reactive({ show: false })

const typeOpts = ["simple", "variant", "compound"]

const categoryOpts = ref<tProduct.Category[]>([])
async function getOpts() {
    loading.value = true
    const res = await productService.listCategory()
    categoryOpts.value = res.rows
    loading.value = false
}

const manageVariants = reactive<{
    show: boolean,
    info: tProduct.ProductVariantUnsaved[]
}>({ show: false, info: [] })

async function create() {
    if (form.type == "variant" && form.variants.length == 0) {
        toast.display({ message: "Insuficcient variants", variant: "error" })
        return
    }
    let params: Parameters<typeof productService.create>[0] = {
        type: form.type,
        name: form.name,
        price: form.price,
        category: form.category
    }
    if (form.type == "variant") {
        params.variants = form.variants
    }
    await productService.create(params)
    toast.display({ message: "Product created", variant: "success" })
    router.push({ name: 'product-list' })
}
getOpts()
</script>

<template>
    <div class="relative min-h-screen">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Create Product</h1>
            <Overlay :show="loading">
                <form @submit.prevent="create" autocomplete="off">
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Type</span>
                        </div>
                        <select class="select select-bordered w-full" required v-model="form.type">
                            <option disabled selected>---</option>
                            <option :value="type" v-for="type in typeOpts">{{ type }}</option>
                        </select>
                    </label>
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
                        <input class="input input-bordered w-full" type="number" step=".01" required v-model="form.price" />
                    </label>
                    <div class="form-control w-full mb-4">
                        <div class="label">
                            <label class="label-text" for="s1">Category</label>
                            <span class="label-text-alt">
                                <button class="btn btn-sm btn-square"
                                    @click.prevent="createCategory.show = !createCategory.show">
                                    <PlusIcon class="w-4 h-4" />
                                </button>
                            </span>
                        </div>
                        <select class="select select-bordered w-full" id="s1" v-model="form.category" required>
                            <option disabled selected>---</option>
                            <option :value="category.id" v-for="category in categoryOpts">{{ category.name }}</option>
                        </select>
                    </div>
                    <div class="form-control w-full mb-4" v-if="form.type == 'variant'">
                        <div class="label">
                            <label class="label-text" for="s1">Variants</label>
                        </div>
                        <div class="join">
                            <input class="input input-bordered join-item w-full" type="number" readonly
                                :value="form.variants.length" />
                            <button class="btn btn-square join-item" role="button"
                                @click.prevent="manageVariants.show = !manageVariants.show">
                                <PencilSquareIcon class="w-6 h-6" />
                            </button>
                        </div>
                    </div>

                    <button class="btn btn-block btn-primary mt-8" type="submit">
                        Create
                    </button>
                </form>
            </Overlay>
        </div>
        <CreateCategoryModal :show="createCategory.show" @closing="createCategory.show = false" @refresh="getOpts" />
        <ManageVariantsModal :show="manageVariants.show" v-model="form.variants" @closing="manageVariants.show = false" />
    </div>
</template>
