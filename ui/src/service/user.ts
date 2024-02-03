import base from "./base"

import * as tUser from "@/type/user"

const path = "user"

export default {
    async list() {
        const res = await base.get(path)

        return res.json<{ rows: tUser.User[] }>()
    },
    async create(body: { username: string, password: string, role: string }) {
        const res = await base.post(path, { json: body })

        return res.status
    },
}
