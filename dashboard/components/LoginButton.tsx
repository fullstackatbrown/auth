'use client'

import { login } from "@/utils/auth"

interface LoginButtonProps {
  authHost: string,
  appHost: string,
}

export default function LoginButton({ authHost, appHost }: LoginButtonProps) {
  function redirectToLogin() {
    login(authHost, appHost);
  }

  return (
    <p onClick={redirectToLogin} style={{ cursor: "pointer" }}>
      Login
    </p>
  )
}