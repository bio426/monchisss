<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router";

import * as tUser from "@/type/user"
import storeService from "@/service/store"
import userService from "@/service/user"
import Overlay from "@/components/Overlay.vue";

const router = useRouter()

const loading = ref(false)

const form = reactive({
    name: "",
    token: "",
    admin: 0
})

const adminOpts = ref<tUser.UserOpt[]>([])
async function getOwnerOpts() {
    loading.value = true
    const data = await userService.listInactiveOwners()
    adminOpts.value = data.rows
    loading.value = false
}
getOwnerOpts()

async function create() {
    loading.value = true
    await storeService.create({ name: form.name, token: form.token, admin: form.admin })
    loading.value = false
    router.push({ name: "store-list" })
}
</script>

<template>
    <div class="relative w-screen min-h-screen">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Create Store</h1>
            <Overlay :show="loading">
                <form @submit.prevent="create" autocomplete="off">
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Name</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.name" />
                    </label>
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Token</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.token" />
                    </label>
                    <label class="form-control w-full mb-8">
                        <div class="label">
                            <span class="label-text">Admin</span>
                        </div>
                        <select class="select select-bordered w-full" required v-model="form.admin">
                            <option disabled selected>---</option>
                            <option :value="admin.id" v-for="admin in adminOpts">{{ admin.username }}</option>
                        </select>
                    </label>

                    <button class="btn btn-block btn-primary" type="submit">
                        Create
                    </button>
                </form>
            </Overlay>
        </div>
    </div>
</template>
