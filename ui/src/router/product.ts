import type { RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw = {
    path: "/product",
    name: "product",
    redirect: { name: "product-list" },
    meta: { permittedRoles: ["owner"] },
    children: [
        {
            path: "list",
            name: "product-list",
            component: () => import("@/view/product/list.vue")
        },
        {
            path: "create",
            name: "product-create",
            component: () => import("@/view/product/create.vue")
        },
    ]
}

export default routes
