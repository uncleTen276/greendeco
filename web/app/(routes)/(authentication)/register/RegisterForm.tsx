'use client'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { RegisterSchema, RegisterInputType } from '@/app/_configs/schemas/authentication'

export default function RegisterForm() {
	const defaultInputValues: RegisterInputType = {
		firstName: '',
		lastName: '',
		email: '',
		phoneNumber: '',
		password: '',
		passwordConfirm: '',
	}
	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<RegisterInputType>({
		mode: 'onBlur',
		reValidateMode: 'onSubmit',
		resolver: zodResolver(RegisterSchema),
		defaultValues: defaultInputValues,
	})
	return (
		<>
			<form
				autoComplete='off'
				onSubmit={handleSubmit((data) => console.log(data))}
				className='flex w-full flex-col gap-cozy text-body-sm'
			>
				<div className='flex-row-between gap-cozy'>
					<div className='flex-1'>
						<input
							type='text'
							placeholder='First Name'
							{...register('firstName')}
						/>
						{errors?.firstName?.message && <p>{errors.firstName.message}</p>}
					</div>
					<div className='flex-1'>
						<input
							type='text'
							placeholder='Last Name'
							{...register('lastName')}
						/>
						{errors?.lastName?.message && <p>{errors.lastName.message}</p>}
					</div>
				</div>
				<div>
					<input
						type='email'
						placeholder='Email'
						{...register('email')}
					/>
					{errors?.email?.message && <p>{errors.email.message}</p>}
				</div>
				<div>
					<input
						type='tel'
						placeholder='Phone Number'
						{...register('phoneNumber')}
					/>
					{errors?.phoneNumber?.message && <p>{errors.phoneNumber.message}</p>}
				</div>
				<div>
					<input
						type='password'
						placeholder='Password'
						{...register('password')}
					/>
					{errors?.password?.message && <p>{errors.password.message}</p>}
				</div>
				<div>
					<input
						type='password'
						placeholder='Confirm Password'
						{...register('passwordConfirm')}
					/>
					{errors?.passwordConfirm?.message && <p>{errors.passwordConfirm.message}</p>}
				</div>
				<button type='submit'>Sign Up</button>
			</form>
		</>
	)
}
