import styles from './page.module.css'
import LoginButton from '@/components/LoginButton'

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <LoginButton authHost={process.env.AUTH_ROOT_URL || ""} appHost={process.env.DASHBOARD_HOME || ""} />
      </div>
    </main>
  )
}
