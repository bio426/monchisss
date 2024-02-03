<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router";

import userService from "@/service/user"
import Overlay from "@/components/Overlay.vue";

const router = useRouter()

const loading = ref(false)

const form = reactive({
    username: "",
    password: "",
    role: "",
})

const roleOpts = ["owner", "employee"]

async function create() {
    loading.value = true
    await userService.create({ username: form.username, password: form.password, role: form.role })
    loading.value = false
    router.push({ name: "super-user" })
}
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Create User</h1>
            <Overlay :show="loading">
                <form @submit.prevent="create" autocomplete="off">
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Username</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                    </label>
                    <label class="form-control w-full mb-4">
                        <div class="label">
                            <span class="label-text">Password</span>
                        </div>
                        <input class="input input-bordered w-full" type="text" required v-model="form.password" />
                    </label>
                    <label class="form-control w-full mb-8">
                        <div class="label">
                            <span class="label-text">Role</span>
                        </div>
                        <select class="select select-bordered w-full" required v-model="form.role">
                            <option disabled selected>---</option>
                            <option :value="role" v-for="role in roleOpts">{{ role }}</option>
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
