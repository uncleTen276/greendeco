'use client'
import { useForm, SubmitHandler } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { RegisterSchema, RegisterFormInputType } from '@/app/_configs/schemas/authentication'
import { useMutation } from '@tanstack/react-query'
import { registerAccount } from '@/app/_api/axios/authentication'
import { AxiosError } from 'axios'

export default function RegisterForm() {
	const defaultInputValues: RegisterFormInputType = {
		firstName: '',
		lastName: '',
		email: '',
		phoneNumber: '',
		password: '',
		passwordConfirm: '',
	}
	const {
		reset,
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<RegisterFormInputType>({
		mode: 'onBlur',
		reValidateMode: 'onSubmit',
		resolver: zodResolver(RegisterSchema),
		defaultValues: defaultInputValues,
	})

	const registerMutation = useMutation({
		mutationFn: registerAccount,
		onSuccess: (data) => {
			reset()
			console.log('success', data)
		},
		onError: (error: AxiosError) => console.log(error.response?.data),
	})

	const onSubmitHandler: SubmitHandler<RegisterFormInputType> = (values) => {
		// ? Execute the Mutation
		registerMutation.mutate({
			identifier: values.email,
			firstName: values.firstName,
			lastName: values.lastName,
			email: values.email,
			password: values.password,
			phoneNumber: values.phoneNumber,
		})
	}

	return (
		<>
			<form
				autoComplete='off'
				onSubmit={handleSubmit(onSubmitHandler)}
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
