import RegisterForm from './RegisterForm'

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
			</div>
		</>
	)
}
