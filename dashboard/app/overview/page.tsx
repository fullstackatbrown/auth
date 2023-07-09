'use client'

import { useEffect } from 'react'
import styles from '../page.module.css'
import { useUser, logout } from "@/utils/auth"
import { useRouter } from 'next/navigation'

export default function Overview() {
  const { user, error, loading } = useUser()

  const router = useRouter()

  useEffect(() => {
    if (error) {
      router.replace("/")
    }
  }, [error, router])

  function signOut() {
    logout()
    router.replace("/")
  }

  return (
    <main className={styles.main}>
      <div className={styles.description}>
        {user && (
          <div>
            <div>Name: {user.displayName}</div>
            <div>Email: {user.email}</div>
          </div>
        )}
        {loading && (
          <p>loading</p>
        )}
        <p onClick={signOut} style={{ cursor: "pointer" }}>
          Sign Out
        </p>
      </div>
    </main>
  )
}
