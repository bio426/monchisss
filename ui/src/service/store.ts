import base from "./base"

// import * as tUser from "@/type/user"

const path = "store"

export default {
    async list() {
        const res = await base.get(path)

        return res.json<{ rows: any[] }>()
    },
    async create(body: { name: string, token: string, admin: number }) {
        const res = await base.post(path, { json: body })

        return res.status
    },
}
