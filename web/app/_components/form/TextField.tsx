import * as React from 'react'
import { FormControl, FormControlProps } from '@mui/base/FormControl'
import Input from './Input'
import clsx from 'clsx'
import { UseFormRegisterReturn } from 'react-hook-form'

type CustomFormControlProps<T> = Partial<T> & {
	label?: string
	helperText?: string
	type?: React.HTMLInputTypeAttribute
	register?: UseFormRegisterReturn
}

export default function TextField(props: CustomFormControlProps<FormControlProps>) {
	const {
		className,
		label,
		helperText,
		type,
		required,
		value,
		error,
		disabled,
		defaultValue,
		register,
		...otherFormControlProps
	} = props

	return (
		<>
			<FormControl
				{...otherFormControlProps}
				className={clsx('flex flex-col gap-compact', className)}
			>
				{label && (
					<label className='font-bold'>
						{label} {required ? '*' : ''}
					</label>
				)}
				<Input
					type={type}
					className='w-full'
					value={value}
					error={error}
					disabled={disabled}
					defaultValue={defaultValue}
					{...register}
				/>
				{helperText && <p className={clsx({ 'text-status-error': error })}>{helperText}</p>}
			</FormControl>
		</>
	)
}
