import type { RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw = {
    path: "/store",
    name: "store",
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
        {
            path: "detail/:id",
            name: "store-detail",
            component: () => import("@/view/store/detail.vue")
        },
    ]
}

export default routes
