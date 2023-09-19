'use client'
import { useForm, SubmitHandler } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { RegisterSchema, RegisterFormInputType } from '@/app/_configs/schemas/authentication'
import { useMutation } from '@tanstack/react-query'
import { registerAccount } from '@/app/_api/axios/authentication'
import { AxiosError } from 'axios'
import { notifyRegisterFail, notifyRegisterSuccess } from '../Notification'
import TextField from '@/app/_components/form/TextField'
import Button from '@/app/_components/Button'

export default function RegisterForm() {
	const defaultInputValues: RegisterFormInputType = {
		firstName: '',
		lastName: '',
		email: '',
		phoneNumber: '',
		password: '',
		passwordConfirm: '',
	}

	//NOTE: Validation with useForm
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
		//NOTE: The callback used for the mutation
		mutationFn: registerAccount,
		//NOTE: Execuse after receiving suscess responses
		onSuccess: () => {
			reset()
			notifyRegisterSuccess()
		},
		//NOTE: Execuse after receving failure responses
		onError: (e) => {
			if (e instanceof AxiosError) {
				notifyRegisterFail(e.response?.data.msg)
			}
		},
	})

	const onSubmitHandler: SubmitHandler<RegisterFormInputType> = (values) => {
		//NOTE: Execute the Mutation
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
				onSubmit={handleSubmit(onSubmitHandler)}
				className='flex w-full flex-col gap-cozy text-body-sm'
			>
				<div className='flex-row-between gap-cozy'>
					<div className='flex-1'>
						<TextField
							type='text'
							label='First Name'
							placeholder='First Name'
							register={register('firstName')}
							error={Boolean(errors?.firstName)}
							helperText={errors?.firstName?.message}
						/>
					</div>
					<div className='flex-1'>
						<TextField
							type='text'
							label='Last Name'
							placeholder='Last Name'
							register={register('lastName')}
							error={Boolean(errors?.lastName)}
							helperText={errors?.lastName?.message}
						/>
					</div>
				</div>
				<div>
					<TextField
						type='email'
						label='Email'
						placeholder='Your Email'
						register={register('email')}
						error={Boolean(errors?.email)}
						helperText={errors?.email?.message}
					/>
				</div>
				<div>
					<TextField
						type='tel'
						label='Phone Number'
						placeholder='Phone Number'
						register={register('phoneNumber')}
						error={Boolean(errors?.phoneNumber)}
						helperText={errors?.phoneNumber?.message}
					/>
				</div>
				<div>
					<TextField
						type='password'
						label='Password'
						placeholder='Password'
						register={register('password')}
						error={Boolean(errors?.password)}
						helperText={errors?.password?.message}
					/>
				</div>
				<div>
					<TextField
						type='password'
						label='Confirm Password'
						placeholder='Confirm Password'
						register={register('passwordConfirm')}
						error={Boolean(errors?.passwordConfirm)}
						helperText={errors?.passwordConfirm?.message}
					/>
				</div>
				<Button
					type='submit'
					disabled={registerMutation.isLoading}
				>
					{registerMutation.isLoading ? 'Sending...' : 'Sign Up'}
				</Button>
			</form>
		</>
	)
}
