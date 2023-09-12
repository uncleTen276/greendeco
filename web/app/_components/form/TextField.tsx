import * as React from 'react'
import { FormControl, FormControlProps } from '@mui/base/FormControl'
import Input from './Input'
import clsx from 'clsx'
export default function TextField(
	props: FormControlPropsWithLabelAndTextHelpder<FormControlProps>,
) {
	return (
		<>
			<FormControl
				{...props}
				required
				className={clsx('flex flex-col gap-[4px]', props.className)}
			>
				{props.label && (
					<label className='font-bold'>
						{props.label} {props.required ? '*' : ''}
					</label>
				)}
				<Input
					type={props.type}
					className='w-full'
					value={props.value}
					error={props.error}
					disabled={props.disabled}
					onChange={props.onChange}
					defaultValue={props.defaultValue}
				/>
				{props.helperText && (
					<p className={clsx({ 'textHelper--error': props.error })}>{props.helperText}</p>
				)}
			</FormControl>
		</>
	)
}

type FormControlPropsWithLabelAndTextHelpder<T> = Partial<T> & {
	label?: string
	helperText?: string
	type?: React.HTMLInputTypeAttribute
}
