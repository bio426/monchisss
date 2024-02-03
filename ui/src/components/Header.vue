<script setup lang="ts">
import { } from "vue"
import { useRouter } from "vue-router"
import { Bars3Icon, HomeIcon, UsersIcon, BuildingStorefrontIcon } from "@heroicons/vue/24/solid"

import authService from "@/service/auth"
import useAuthStore from "@/store/auth"

const props = defineProps<{ title: string }>()

const router = useRouter()
const authStore = useAuthStore()

async function logout() {
    await authService.logout()
    authStore.clearUser()
    router.push({ name: "login" })
}
</script>

<template>
    <div class="navbar bg-base-100">
        <div class="navbar-start">
            <div class="dropdown">
                <label tabindex="0" class="btn btn-ghost btn-circle">
                    <Bars3Icon class="w-6" />
                </label>
                <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
                    <li>
                        <router-link active-class="active" :to="{ name: 'dashboard' }">
                            <HomeIcon class="w-4" />
                            Dashboard
                        </router-link>
                    </li>
                    <!-- Super user modules -->
                    <template v-if="authStore.user?.role == 'super'">
                        <li>
                            <router-link active-class="active" :to="{ name: 'super-user' }">
                                <UsersIcon class="w-4" />
                                User
                            </router-link>
                        </li>
                        <li>
                            <router-link active-class="active" :to="{ name: 'store' }">
                                <BuildingStorefrontIcon class="w-4" />
                                Store
                            </router-link>
                        </li>
                    </template>
                </ul>
            </div>
        </div>
        <div class="navbar-center">
            <h1 class="py-4 text-2xl font-bold text-center">
                {{ props.title }}
            </h1>
        </div>
        <div class="navbar-end">
            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar placeholder">
                    <div class="bg-neutral text-neutral-content rounded-full w-12">
                        <span>{{ authStore.user?.username.at(0)?.toUpperCase() }}</span>
                    </div>
                </div>
                <ul tabindex="0" class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
                    <li>
                        <button @click="logout">Logout</button>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>
