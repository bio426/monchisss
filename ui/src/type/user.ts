import * as tAuth from "./auth"

export type User = tAuth.User & {
    id: number
    active: boolean
    createdAt: string
}
