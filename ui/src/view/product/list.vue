<script setup lang="ts">
import { ref } from "vue"
import { PencilIcon, TrashIcon } from "@heroicons/vue/24/solid";

import * as tProduct from "@/type/product"
import productService from "@/service/product"
import Header from "@/components/Header.vue";
import Overlay from "@/components/Overlay.vue";

const loading = ref(false)

const rows = ref<tProduct.Product[]>([])

async function getRows() {
    loading.value = true
    const res = await productService.list()
    rows.value = res.rows
    loading.value = false
}
getRows()
</script>

<template>
    <div class="relative min-h-screen w-full">
        <div class="w-11/12 mx-auto">
            <Header title="Product" />
            <div class="mb-4">
                <router-link class="btn btn-primary" :to="{ name: 'product-create' }">Create</router-link>
            </div>
            <Overlay :show="loading">
                <div class="overflow-x-auto">
                    <table class="table">
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>Price</th>
                                <th>Category</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="row in rows" :key="row.name">
                                <td>{{ row.name }}</td>
                                <td>{{ row.price }}</td>
                                <td>{{ row.category }}</td>
                                <td>
                                    <div class="flex gap-2">
                                        <button class="btn btn-xs btn-square" title="Edit">
                                            <PencilIcon class="w-4 y-4 fill-warning" />
                                        </button>
                                        <button class="btn btn-xs btn-square" title="Edit">
                                            <TrashIcon class="w-4 y-4 fill-error" />
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <div class="text-center" v-if="rows.length == 0">No data available</div>
                </div>
            </Overlay>
        </div>
    </div>
</template>
