export default function RegisterForm() {
	return (
		<>
			<form className='flex flex-col gap-cozy'>
				<input
					type='text'
					placeholder='first name'
				/>
				<input
					type='text'
					placeholder='last name'
				/>
				<input type='text' />
			</form>
		</>
	)
}
