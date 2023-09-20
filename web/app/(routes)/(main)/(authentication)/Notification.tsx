import React from 'react'
import { ToastContentProps, toast } from 'react-toastify'
import Link from 'next/link'

export const notifyRegisterSuccess = () => {
	toast.success(<RegisterSuccessMessage />, {
		position: 'top-center',
	})
}
export const notifyRegisterFail = (errorMessage?: string) => {
	toast.error(<RegisterErrorMessage message={errorMessage ? errorMessage : undefined} />, {
		position: 'top-center',
	})
}

const RegisterSuccessMessage = () => (
	<div className='flex flex-col gap-[4px] pl-compact pr-common text-body-sm'>
		<h3 className='capitalize text-primary-625'>Registered Successfully</h3>
		<p>
			Moving to the{' '}
			<Link
				href={'../login'}
				className='hover:font-bold'
			>
				Login Page
			</Link>
		</p>
	</div>
)
const RegisterErrorMessage = ({ message }: { message?: string }) => (
	<div className='flex flex-col gap-[4px] pl-compact pr-common text-body-sm'>
		<h3 className='capitalize text-status-error'>{message ? message : 'Registered Failed'}</h3>
	</div>
)
