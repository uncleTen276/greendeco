import * as React from 'react'
import { FormControl, FormControlProps } from '@mui/base/FormControl'
import Input from './Input'
import clsx from 'clsx'

type CustomFormControlProps<T> = Partial<T> & {
	label?: string
	helperText?: string
	type?: React.HTMLInputTypeAttribute
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
		onChange,
		defaultValue,
		...otherFormControlProps
	} = props

	return (
		<>
			<FormControl
				{...otherFormControlProps}
				className={clsx('flex flex-col gap-[4px]', className)}
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
					onChange={onChange}
					defaultValue={defaultValue}
				/>
				{helperText && <p className={clsx({ 'text-status-error': error })}>{helperText}</p>}
			</FormControl>
		</>
	)
}
