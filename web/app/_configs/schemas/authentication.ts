import { MIN_PASSWORD, MAX_PASSWORD } from '../constants/variables'
import * as z from 'zod'

export const RegisterSchema = z
	.object({
		firstName: z
			.string()
			.min(1, 'First name is required')
			.max(32, 'Name must be less than 32 characters'),
		lastName: z
			.string()
			.min(1, 'Last name is required')
			.max(32, 'Name must be less than 32 characters'),
		email: z.string().min(1, 'Email is required').email('Email is invalid'),
		phoneNumber: z.string().min(1, 'Phone number is required').max(9, 'Invalid phone number'),

		password: z
			.string()
			.min(MIN_PASSWORD, `Password must be more than ${MIN_PASSWORD} characters`)
			.max(MAX_PASSWORD, `Password must be less than ${MAX_PASSWORD} characters`),
		passwordConfirm: z.string().min(1, 'Please confirm your password'),
	})
	.refine((data) => data.password === data.passwordConfirm, {
		path: ['passwordConfirm'],
		message: 'Passwords do not match',
	})

export type RegisterInputType = z.infer<typeof RegisterSchema>
