import type { RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw = {
    path: "/store",
    name: "store",
    // component: () => import("@/view/store/list.vue"),
    redirect: { name: "store-list" },
    meta: { permittedRoles: ["super"] },
    children: [
        {
            path: "list",
            name: "store-list",
            component: () => import("@/view/store/list.vue")
        },
        {
            path: "create",
            name: "store-create",
            component: () => import("@/view/store/create.vue")
        },
    ]
}

export default routes
