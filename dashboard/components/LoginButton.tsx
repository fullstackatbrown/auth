'use client'

interface LoginButtonProps {
    login: string,
    home: string,
}

export default function LoginButton({login, home}: LoginButtonProps) {
    function redirectToLogin() {
        location.href = `${login}?from=${home}`
      }    

    return (
        <p onClick={redirectToLogin} style={{cursor: "pointer"}}>
          Login
        </p>
    )
}