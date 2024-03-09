<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter, useRoute } from "vue-router";

import * as tStore from "@/type/store"
import storeService from "@/service/store"
import Overlay from "@/components/Overlay.vue";

const router = useRouter()
const route = useRoute()

let storeId = 0
function checkParams() {
    const checkErr = new Error("Bad Route Param")
    const param = route.params.id
    if (typeof param != "string") throw checkErr
    storeId = parseInt(param)
    if (isNaN(storeId)) throw checkErr
}
checkParams()

const loading = ref(false)

const userRows = ref<tStore.User[]>([])

async function getUsers() {
    loading.value = true
    const res = await storeService.listUsers(storeId)
    userRows.value = res.rows
    loading.value = false
}
getUsers()

</script>

<template>
    <div class="relative min-h-screen">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Store Detail</h1>
            <Overlay :show="loading">
                <div class="p-4 border border-primary rounded">
                    <h2 class="text-xl">Users</h2>
                    <hr class="border-primary my-2">
                    <div class="overflow-x-auto">
                        <table class="table whitespace-nowrap">
                            <thead>
                                <tr>
                                    <th>Username</th>
                                    <th>Role</th>
                                    <th>Created</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="row in userRows">
                                    <td>{{ row.username }}</td>
                                    <td>{{ row.role }}</td>
                                    <td>{{ new Date(row.createdAt).toLocaleDateString() }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </Overlay>
        </div>
    </div>
</template>
