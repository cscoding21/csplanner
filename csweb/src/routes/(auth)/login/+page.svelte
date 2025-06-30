<script lang="ts">
	import { parseErrors, mergeErrors } from '$lib/forms/helpers';
	import { loginForm, loginSchema } from '$lib/forms/login.validation';
	import { goto } from '$app/navigation';
	import { EmailInput, PasswordInput, AlertError } from '$lib/components';
	import { authService } from '$lib/services/auth';

	let errors: any = $state({});
	let loginFailed: boolean = $state(false);
	let as = authService();

	const submitForm = async () => {
		errors = {};

		loginSchema
			.validate(loginForm, { abortEarly: false })
			.then(() => {
				console.log('client-size validation passed');

				as.login({ email: loginForm.email, password: loginForm.password })
					.then((r) => {
						console.log(r);
						if (r) {
							loginFailed = false;
							goto('/home');
						} else {
							loginFailed = true;
						}
					})
					.catch((err) => {
						loginFailed = true;
						console.log(err);
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.error(errors);
			});
	};
</script>

<svelte:head>
	<title>csPlanner: Login</title>
</svelte:head>

<div class="mt-8">
	<div
		class="w-full rounded-lg border border-gray-200 bg-white p-4 shadow dark:border-gray-700 dark:bg-gray-800 sm:p-6 md:p-8"
	>
		<form class="space-y-6" action="#">
			<h5 class="text-xl font-medium text-gray-900 dark:text-white">
				Sign in to <span class="text-md text-amber-500">cs</span>Planner
			</h5>
			<AlertError
				showAlert={loginFailed}
				title="Login Attempt Failed"
				message="We were unable to validate your login credentials"
			/>

			<EmailInput bind:value={loginForm.email} fieldName="Email" error={errors.email} />

			<PasswordInput bind:value={loginForm.password} fieldName="Password" error={errors.password} />

			<button
				type="button"
				onclick={submitForm}
				class="w-full rounded-lg bg-blue-700 px-5 py-2.5 text-center text-sm font-medium text-white hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
				>Login</button
			>
		</form>
	</div>
</div>
