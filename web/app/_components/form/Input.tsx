'use client'
import * as React from 'react'
import { Input as BaseInput, InputProps } from '@mui/base/Input'
import clsx from 'clsx'

const CustomInput = React.forwardRef(function CustomInput(
	props: InputProps,
	ref: React.ForwardedRef<HTMLInputElement>,
) {
	const { className } = props
	return (
		<BaseInput
			{...props}
			ref={ref}
			slotProps={{
				root: {
					className: clsx(['base-input', className], {
						'input-disabled': props.disabled,
						'input-error': props.error,
					}),
				},
				input: {
					className: 'w-full outline-none bg-inherit',
				},
			}}
		/>
	)
})

const Input = React.forwardRef(function Input(
	props: InputProps,
	ref: React.ForwardedRef<HTMLInputElement>,
) {
	const renderInputType = (type: React.HTMLInputTypeAttribute) => {
		return <span>{type}</span>
	}
	return (
		<CustomInput
			{...props}
			ref={ref}
			startAdornment={props.type && renderInputType(props.type)}
		/>
	)
})

export default CustomInput
export { Input }
