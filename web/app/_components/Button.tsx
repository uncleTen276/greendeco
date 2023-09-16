import clsx from 'clsx'
import { Button as BaseButton, ButtonProps } from '@mui/base/Button'
import React from 'react'

const Button = React.forwardRef(function Button(
	props: ButtonProps,
	ref: React.ForwardedRef<HTMLButtonElement>,
) {
	const { className, ...otherButtonProps } = props
	return (
		<BaseButton
			{...otherButtonProps}
			ref={ref}
			className={clsx('btn', className)}
		/>
	)
})

export default Button
