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
				onSubmit={handleSubmit((data) => console.log(data))}
				className='flex flex-col gap-cozy'
			>
				<input
					type='text'
					placeholder='first name'
					{...register('firstName')}
				/>
				{errors?.firstName?.message && <p>{errors.firstName.message}</p>}
				<input
					type='text'
					placeholder='last name'
				/>
				<input type='text' />
			</form>
		</>
	)
}
