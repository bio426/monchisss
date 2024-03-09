import base from "./base"

import * as tProduct from "@/type/product"

const path = "product"

type SubProduct = { name: string, price: number }
type CreatePayload = {
    name: string
    type: string
    price: number
    category: number
    variants?: SubProduct[]
    components?: SubProduct[]
}

export default {
    async list() {
        const res = await base.get(path)

        return res.json<{ rows: tProduct.Product[] }>()
    },
    async create(body: CreatePayload) {
        const res = await base.post(path, { json: body })

        return res.status
    },
    async listCategory() {
        const res = await base.get(path + "/category")

        return res.json<{ rows: tProduct.Category[] }>()
    },
    async createCategory(body: { name: string }) {
        const res = await base.post(path + "/category", { json: body })

        return res.status
    },
}
