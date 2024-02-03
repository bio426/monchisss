import type { RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw = {
    path: "/super",
    name: "super",
    redirect: { name: "super-user" },
    meta: { permittedRoles: ["super"] },
    children: [
        {
            path: "user",
            name: "super-user",
            component: () => import("@/view/super/user.vue")
        },
        {
            path: "user/create",
            name: "super-user-create",
            component: () => import("@/view/super/user-create.vue")
        },
    ]
}

export default routes
