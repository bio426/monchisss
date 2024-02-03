import { createRouter, createWebHistory } from 'vue-router'
import type { RouteLocationNormalized } from 'vue-router'

import useAuthStore from "@/store/auth"
import useToast from "@/composables/useToast"
import superRouter from "./super"
import store from "./store"

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'dashboard',
            component: () => import("@/view/dashboard.vue")
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/view/login.vue')
        },
        {
            path: "/:pathMatch(.*)*",
            name: "not-found",
            component: () => import("@/view/not-found.vue"),
        },
        superRouter,
        store,
    ]
})

router.beforeEach((to, _) => {
    const authStore = useAuthStore()
    const routeName = to.name as string
    const publicRoutes = ["login", "not-found"]

    if (publicRoutes.includes(routeName)) {
        if (authStore.authenticated && routeName == "login") {
            return { name: "dashboard" }
        }
        return true
    } else {
        if (!authStore.authenticated) {
            unauthorizedAlert()
            return { name: "login" }
        }
        const userRole = authStore.user?.role as string
        const hasRolePermission = checkRolePermission(to, userRole)
        if (!hasRolePermission) {
            forbiddenAlert()
            return { name: "dashboard" }
        }
        return true
    }
})

function checkRolePermission(to: RouteLocationNormalized, role: string): boolean {
    const allowedRoles = to.meta.permittedRoles as (string[] | undefined)
    if (allowedRoles == undefined || allowedRoles.includes(role)) return true
    return false
}

function unauthorizedAlert() {
    const toast = useToast()
    toast.display({ message: "Unauthorized", variant: "error" })
}

function forbiddenAlert() {
    const toast = useToast()
    toast.display({ message: "Forbidden", variant: "error" })
}

export default router
