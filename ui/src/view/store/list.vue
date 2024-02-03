<script setup lang="ts">
import { ref } from "vue"
import { CheckIcon, XMarkIcon, PencilIcon, TrashIcon } from "@heroicons/vue/24/solid"

import * as tUser from "@/type/user"
import userService from "@/service/user"
import Header from "@/components/Header.vue"
import Overlay from "@/components/Overlay.vue"

const loading = ref(false)

const rows = ref<tUser.User[]>([])

async function getRows() {
    loading.value = true
    const res = await userService.list()
    rows.value = res.rows
    loading.value = false
}
getRows()
</script>

<template>
    <div class="relative min-h-screen w-full">
        <div class="w-11/12 mx-auto">
            <Header title="Store" />
            <div class="mb-4">
                <router-link class="btn btn-primary" :to="{ name: 'store-create' }">Create</router-link>
            </div>
            <Overlay :show="loading">
                <div class="overflow-x-auto">
                    <table class="table">
                        <thead>
                            <tr>
                                <th>Name</th>
                                <th>Admin</th>
                                <th>Employees</th>
                                <th>Created</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="row in rows" :key="row.id">
                                <td>{{ row.username }}</td>
                                <td>{{ row.role }}</td>
                                <td>
                                    <span v-if="row.active">yes</span>
                                    <span v-else>no</span>
                                    <!-- <CheckIcon class="w-4 y-4 fill-success" v-if="row.active" /> -->
                                    <!-- <XMarkIcon class="w-4 y-4 fill-error" v-else /> -->
                                </td>
                                <td>{{ new Date(row.createdAt).toLocaleDateString() }}</td>
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
