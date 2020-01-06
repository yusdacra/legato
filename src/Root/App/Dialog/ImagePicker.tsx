import React, { useRef } from 'react';
import { ButtonBase, Avatar } from '@material-ui/core';

interface Props {
	setImageFile: React.Dispatch<React.SetStateAction<File | null>>;
	image: string | undefined;
	setImage: React.Dispatch<React.SetStateAction<string | undefined>>;
}

export const ImagePicker = (props: Props) => {
	const imageUploadRef = useRef<HTMLInputElement | null>(null);

	const onImageSelected = (event: React.ChangeEvent<HTMLInputElement>) => {
		if (event.currentTarget.files && event.currentTarget.files.length > 0) {
			const file = event.currentTarget.files[0];
			props.setImageFile(file);
			if (file.type.startsWith('image/') && file.size < 33554432) {
				const fileReader = new FileReader();
				fileReader.readAsDataURL(file);
				fileReader.addEventListener('load', () => {
					if (typeof fileReader.result === 'string') {
						props.setImage(fileReader.result);
					}
				});
			}
		}
	};

	return (
		<>
			<input
				type="file"
				id="file"
				multiple
				ref={imageUploadRef}
				style={{ display: 'none' }}
				onChange={onImageSelected}
			/>
			<ButtonBase
				style={{ borderRadius: '50%' }}
				onClick={() => {
					if (imageUploadRef.current) {
						imageUploadRef.current.click();
					}
				}}
			>
				<Avatar style={{ width: '100px', height: '100px' }} src={props.image}></Avatar>
			</ButtonBase>
		</>
	);
};
