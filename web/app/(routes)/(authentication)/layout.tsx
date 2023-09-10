import { ReactNode } from 'react'

export default function AuthenticationLayout({ children }: { children: ReactNode }) {
	return <div className='h-screen w-screen'>{children}</div>
}
