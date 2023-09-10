import * as React from 'react'
import { FormControl, FormControlProps } from '@mui/base/FormControl'
import { Input } from '@mui/base/Input'

export default function TextField(props: FormControlProps) {
	return (
		<>
			<FormControl
				{...props}
				defaultValue=''
				required
			>
				<Input
					placeholder='Write your name here'
					slotProps={{
						input: {
							className:
								'p-[1.2rem] max-w-full border-primary-5555-20 border-[0.1rem] rounded-[0.3rem]',
						},
					}}
				/>
			</FormControl>
		</>
	)
}
