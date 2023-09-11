'use client'
import * as React from 'react'
import { Input as BaseInput, Input, InputProps } from '@mui/base/Input'

const CustomInput = React.forwardRef(function CustomInput(
	props: InputProps,
	ref: React.ForwardedRef<HTMLInputElement>,
) {
	return (
		<BaseInput
			{...props}
			ref={ref}
			slotProps={{
				input: {
					className: 'base-input',
				},
			}}
		/>
	)
})

export default CustomInput
