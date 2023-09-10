import { ReactNode } from 'react'
import { Metadata } from 'next'

export const metadata: Metadata = {
	title: 'Create An Account',
	description: 'Be a member by creating an account',
}
export default function AuthenticationLayout({ children }: { children: ReactNode }) {
	return (
		<div className='flex-center h-screen w-screen bg-primary-580-20'>
			<div className='container my-auto grid h-[95%] max-h-full grid-cols-2 overflow-hidden rounded-xl bg-white shadow-38'>
				<div className='flex-center max-h-full p-[3.2rem]'>{children}</div>
				<div className=''>
					<img
						className='object-fill'
						src='https://static.vecteezy.com/system/resources/previews/018/815/357/original/green-tropical-forest-background-monstera-leaves-palm-leaves-branches-exotic-plants-background-for-banner-template-decor-postcard-abstract-foliage-and-botanical-wallpaper-vector.jpg'
						alt=''
					/>
				</div>
			</div>
		</div>
	)
}
