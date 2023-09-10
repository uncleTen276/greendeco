import { MIN_PASSWORD, MAX_PASSWORD } from '../constants/variables'
import * as z from 'zod'

export const RegisterSchema = z
	.object({
		name: z
			.string()
			.min(1, 'Name is required')
			.max(32, 'Name must be less than 100 characters'),
		email: z.string().min(1, 'Email is required').email('Email is invalid'),
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
