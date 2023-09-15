import RegisterForm from './RegisterForm'
import Link from 'next/link'

export default function RegisterPage() {
	return (
		<>
			<div className='flex-col-start gap-common'>
				<div>
					<span className='mb-compact block text-body-lg'>
						Welcome to <span className=' text-heading-2 font-bold'>GreenDeco</span> 👋
					</span>
					<div className='flex-col-start gap-[4px]'>
						<h1>Create An Account</h1>
						<p className='text-body-md'>Become one of the plant lovers now!</p>
					</div>
				</div>
				<RegisterForm />
				<span className='text-center text-body-md'>
					Don&apos;t you have an account? <Link href={'/login'}>Sign up</Link>
				</span>
			</div>
		</>
	)
}
