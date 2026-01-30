import Link from 'next/link'

import styles from './Header.module.scss'

export interface HeaderProps {
  login?: boolean
}

export default function Header({ login }: HeaderProps) {
  return (
    <header className={styles.header}>
      <img src="@/../public/logo.png" alt="Logo" className={styles.logo} />
      <nav>
        <ul className={styles.navlist}>
          <li>
            <Link href="docs">docs</Link>
          </li>
          <li>
            <Link href="price">price</Link>
          </li>
          <li>
            <Link href="notion">notion</Link>
          </li>
        </ul>
      </nav>
      <div className={styles.authButton}>{login ? 'Logout' : 'Login'}</div>
    </header>
  )
}
