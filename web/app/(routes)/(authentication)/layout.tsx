import { ReactNode } from 'react'

export default function AuthenticationLayout({ children }: { children: ReactNode }) {
	return (
		<div className='flex-center h-screen w-screen bg-primary-580-20'>
			<div className='container m-auto grid h-[90%] max-h-full grid-cols-2 overflow-hidden rounded-xl bg-white shadow-38'>
				<div className='flex-center max-h-full p-[3.2rem]'>{children}</div>
				<div className=''></div>
			</div>
		</div>
	)
}
