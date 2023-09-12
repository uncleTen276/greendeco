'use client'
import * as React from 'react'
import { Input as BaseInput, InputProps } from '@mui/base/Input'
import clsx from 'clsx'
import { EnvelopeIcon, LockClosedIcon } from '@heroicons/react/24/solid'

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
			startAdornment={type && renderStartAndormentIcon(type)}
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

//NOTE: Input renderStartAdornment

type StartIconType = {
	type: 'email' | 'password'
	icon: React.ReactNode
}

const StartIconArray: StartIconType[] = [
	{
		type: 'email',
		icon: <EnvelopeIcon className='aspect-square w-[1.6rem] text-primary-418-60' />,
	},
	{
		type: 'password',
		icon: <LockClosedIcon className='aspect-square w-[1.6rem] text-primary-418-60' />,
	},
]

const renderStartAndormentIcon = (type: React.HTMLInputTypeAttribute) => {
	const startIcon = StartIconArray.find((startIcon) => startIcon['type'] == type)
	if (startIcon) return startIcon.icon
}

export default CustomInput
