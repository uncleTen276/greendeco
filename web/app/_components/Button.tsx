import clsx from 'clsx'
import { Button as BaseButton, ButtonProps } from '@mui/base/Button'
import React from 'react'

const Button = React.forwardRef(function Button(
	props: ButtonProps,
	ref: React.ForwardedRef<HTMLButtonElement>,
) {
	const { className, type } = props
	return (
		<BaseButton
			{...props}
			ref={ref}
			className={clsx(
				'cursor-pointer rounded-lg border-none bg-violet-500 px-4 py-2 font-sans text-sm font-semibold text-white disabled:cursor-not-allowed disabled:opacity-50',
				className,
			)}
		/>
	)
})

export default Button
