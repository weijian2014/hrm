interface TokenInfo {
   user_name: string
   access_token: string
   refresh_token: string
   expired_at: string
}

interface ChangePasswordRequest {
   username: string
   old_password: string
   new_password: string
   new_password_repeat: string
}
