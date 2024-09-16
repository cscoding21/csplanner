import * as tus from 'tus-js-client';
import { authService } from '$lib/services/auth';
import { appConfig } from '$lib/appConfig';
import { isFunction } from '$lib/utils/check';

export const handleFileUpload = (
	file: any,
	context: string,
	artifactType: string,
	completed: any
) => {
	const as = authService();

	// Create a new tus upload
	const upload = new tus.Upload(file, {
		// Endpoint is the upload creation URL from your tus server
		endpoint: appConfig.fileEndpoint,
		//---let's give headers a try
		headers: as.getAuthHeaders(),
		// Retry delays will enable tus-js-client to automatically retry on errors
		retryDelays: [0, 3000, 5000, 10000, 20000],
		// Attach additional meta data about the file for the server
		metadata: {
			filename: file.name,
			filetype: file.type,
			context: context,
			artifacttype: artifactType
		},
		// Callback for errors which cannot be fixed using retries
		onError: function (err) {
			console.error('Failed because: ' + err);
		},
		// Callback for reporting upload progress
		onProgress: function (bytesUploaded, bytesTotal) {
			const percentage = ((bytesUploaded / bytesTotal) * 100).toFixed(2);
			console.log(bytesUploaded, bytesTotal, percentage + '%');
		},
		// Callback for once the upload is completed
		onSuccess: function () {
			if (isFunction(completed)) {
				completed(upload);
			}
		}
	});

	// Check if there are any previous uploads to continue.
	upload.findPreviousUploads().then(function (previousUploads) {
		// Found previous uploads so we select the first one.
		if (previousUploads.length) {
			upload.resumeFromPreviousUpload(previousUploads[0]);
		}

		// Start the upload
		upload.start();
	});
};
