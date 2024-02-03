type UserRole = "super" | "owner" | "employee"

export type User = {
    username: string
    role: UserRole
}

export type ExpirableUser = User & { expiryDate: string }
