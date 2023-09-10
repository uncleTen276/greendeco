import RegisterForm from './RegisterForm'
import Link from 'next/link'

export default function RegisterPage() {
	return (
		<>
			<div className='flex-col-start gap-common'>
				<div>
					<span className='mb-cozy block text-body-xl'>
						Welcome to <span className='text-heading font-bold'>GreenDeco</span> 👋
					</span>
					<div className='flex-col-start gap-compact'>
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
