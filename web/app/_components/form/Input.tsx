'use client'
import * as React from 'react'
import { Input as BaseInput, InputProps } from '@mui/base/Input'
import clsx from 'clsx'
import { log } from 'console'

//NOTE: Input
const CustomInput = React.forwardRef(function CustomInput(
	props: InputProps,
	ref: React.ForwardedRef<HTMLInputElement>,
) {
	const { className, type } = props
	return (
		<BaseInput
			{...props}
			ref={ref}
			startAdornment={type && renderInputType(type)}
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

//NOTE: Input startAdornment
const renderInputType = (type: React.HTMLInputTypeAttribute) => {
	const startIcon = StartIconArray.find((startIcon) => startIcon['type'] == type)
	if (startIcon) return startIcon.icon
}

type StartIconType = {
	type: 'email' | 'password'
	icon: React.ReactNode
}

const StartIconArray: StartIconType[] = [
	{
		type: 'email',
		icon: <span>Email</span>,
	},
	{
		type: 'password',
		icon: <span>Password</span>,
	},
]

export default CustomInput
